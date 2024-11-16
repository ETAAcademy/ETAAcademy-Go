# ETAAcademy-Go: 02. Web App

<table>
  <tr>
    <th>title</th>
    <th>tags</th>
  </tr>
  <tr>
    <td>02. Web App</td>
    <td>
      <table>
        <tr>
          <th>go</th>
          <th>basic</th>
          <td>Web App</td>
        </tr>
      </table>
    </td>
  </tr>
</table>

[Github](https:github.com/ETAAcademy)｜[Twitter](https:twitter.com/ETAAcademy)｜[ETA-Go](https://github.com/ETAAcademy/ETAAcademy-Go)

Authors: [Eta](https:twitter.com/pwhattie), looking forward to your joining

# Building Web Applications in Go

The Go programming language makes it simple to build powerful web applications without relying on external frameworks. Using Go's `net/http` package, you can create a functional web server with just two lines of code. Below, we’ll explore how to handle web requests, set up servers, and create custom handlers.

**Quick Start: A Simple Web Server**

The following code demonstrates how to create a web server in Go that responds with "hello goooooo" to any request:

<details><summary><b> Code</b></summary>

```go
package main

import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello goooooo"))
    })

    http.ListenAndServe("localhost:8080", nil)
}
```

</details>

Here’s what happens:

1. `http.HandleFunc` maps the root path (`/`) to a handler function that writes "hello goooooo" as a response.
2. `http.ListenAndServe` starts the server on `localhost:8080` and uses the default multiplexer (`http.DefaultServeMux`) to route incoming requests.

---

## **1. Web Requests**

1. **Handling Requests:**
   Go provides two key functions to handle web requests:

   - `http.Handle` for associating specific handlers with routes.
   - `http.HandleFunc` for directly associating a function with a route.

2. **Creating Web Servers:**
   The `http.ListenAndServe` function starts a web server with two parameters:

   - The network address (e.g., `localhost:8080`).
   - A handler. If set to `nil`, Go uses the default multiplexer (`http.DefaultServeMux`).

   Example using a custom server structure:

   ```go
   server := http.Server{
       Addr:    "localhost:8080",
       Handler: nil, // Uses DefaultServeMux if nil
   }
   server.ListenAndServe()
   ```

3. **HTTPS Support:**
   To serve HTTPS requests, you can use:
   - `http.ListenAndServeTLS`.
   - `Server.ListenAndServeTLS`.

### **1.1 Creating a Web Server**

The `http.Handle` and `http.HandleFunc` functions in Go handle (Handle) web requests.

The handler receives requests, and Go creates a goroutine for each request. Instead of manually creating such handlers, it's more common to use Go's built-in `http.DefaultServeMux`.

Go offers several methods to create a web server, and `http.ListenAndServe()` is one of them. The first parameter is the network address (if it's an empty string `""`, it listens on all network interfaces at port 80). The second parameter is a handler (if it’s `nil`, it uses the `DefaultServeMux`). The `DefaultServeMux` is a multiplexer (similar to a router), and it acts as a handler itself, distributing requests to different handlers based on the request paths.

The source code for `http.Server` is a struct with two fields and a function: `Addr` (network address) and `Handler`. If `Handler` is `nil`, it defaults to `DefaultServeMux`. The `http.ListenAndServe()` function internally calls the `Server` struct’s `ListenAndServe()` method without parameters.

You can define a custom server using `http.Server{}`, equivalent to `http.ListenAndServe("localhost:8080", nil)`.

To serve HTTPS, using SSL/TLS, you need to call `http.ListenAndServeTLS()` or `server.ListenAndServeTLS()` on the `http.Server` struct.

**Understanding Handlers**

The two methods for creating a web server both mention handlers. A handler is an interface in Go, defined by a method `ServeHTTP(ResponseWriter, *Request)`. This method takes two parameters, both passed by reference. Anything implementing this method can be a handler.

The first parameter is `w http.ResponseWriter`, an interface representing the response writer, which is a pointer to a struct implementing various methods. The second parameter is a pointer to a `Request` struct, `r *http.Request`, representing the web request.

<details><summary><b> Code</b></summary>

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

```go
package main

import "net/http"

func main() {
	mh := myHandler{} // Using the handler pointer

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: &mh,
	}
	server.ListenAndServe()
}

// Custom handler implementing the ServeHTTP method
type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("helllllllo my handler"))
}
```

</details>

**Handling Multiple Paths with Different Handlers**

**Method 1:** Without specifying the `Handler` field in the `Server` struct (let it be `nil` to use `DefaultServeMux`), attach handlers to `DefaultServeMux` using `http.Handle`.

In the `http` package, there's a `Handle` function, and `ServerMux` struct also has a `Handle` method. When calling `http.Handle`, it actually calls the `Handle` method on `DefaultServeMux`.

<details><summary><b> Code</b></summary>

```go
package main

import "net/http"

func main() {
	mh := myHandler{}
	about := aboutHandler{}
	hello := helloHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}

	// Different paths with corresponding handlers
	http.Handle("/hello", &hello)
	http.Handle("/about", &about)
	http.Handle("/home", &mh)
	server.ListenAndServe()
}

// Custom handlers implementing ServeHTTP
type myHandler struct{}
type helloHandler struct{}
type aboutHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home handler"))
}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello handler"))
}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about handler"))
}
```

</details>

**Method 2:** Using `http.HandleFunc` instead of a handler. A handler function behaves like a handler because it has the same signature as the `ServeHTTP` method. Go defines a function type `HandlerFunc` that allows converting a function with the right signature into a handler.

<details><summary><b> Code</b></summary>

```go
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

Example Using `http.HandleFunc`:

```go
package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	// Routing using HandleFunc
	http.HandleFunc("/welcome", welcomeExample)
	server.ListenAndServe()
}

func welcomeExample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}
```

</details>

**5 Built-in Handlers:**

1. **`NotFoundHandler`**: Responds with "404 page not found" for each request.
2. **`RedirectHandler`**: For redirection.
3. **`StripPrefix`**: A middleware or decorator-like handler that calls another handler.
4. **`TimeoutHandler`**: Similar to a decorator for adding a timeout.
5. **`FileServer`**: Serves requests based on a root filesystem.

This approach allows you to create powerful web applications in Go using just its standard library, without any external frameworks.

**Tips: Two Main Functions**

- `http.Handle`: First argument is the route path (e.g., `"/home"`), second is a `Handler`.
- `http.HandleFunc`: First argument is the route path, second is a handler **function** (converts handler functions to `Handler`).

---

### **1.2 Form Requests**

Form data is a common type of input in web applications, typically sent via POST requests. In Go, handling such data involves working with various fields and methods, such as `Form`, `PostForm`, and `MultipartForm`, as well as handling uploaded files and processing JSON data. This article explores these aspects in detail.

### **Form Requests Overview**

A POST request from a form refers to data from an HTML form submitted as name-value pairs in the request's body. The format of these name-value pairs depends on the form's `enctype` attribute:

1. **`application/x-www-form-urlencoded`**  
   The browser encodes the form data as query strings. This format is ideal for simple text data.  
   Example:

   ```
   first_name=sau%20sheong&last_name=chang
   ```

2. **`multipart/form-data`**  
   Each name-value pair is converted into a MIME message part, with its own `Content-Type` and `Content-Disposition`. This format is typically used for large data, such as file uploads.

The HTTP method used (`GET` or `POST`) determines how data is sent:

- **GET**: All data is appended to the URL as query parameters.
- **POST**: Data is sent in the request body.

**Accessing Form Data**

The `http.Request` object in Go provides methods to extract data from both the URL and the request body:

- **`Form`**: Contains data from both the URL query and the body.
- **`PostForm`**: Contains data from the body only (ignoring URL parameters).
- **`MultipartForm`**: Contains data for forms with `multipart/form-data` encoding.

To access these fields, you must call `ParseForm` or `ParseMultipartForm` to parse the incoming request.

<details><summary><b> Code</b></summary>

HTML form example:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Form Test</title>
  </head>
  <body>
    <form
      action="http://localhost:8080/process?first_name=Eittah"
      method="post"
      enctype="application/x-www-form-urlencoded"
    >
      First Name:
      <input type="text" name="first_name" />
      Last Name:
      <input type="text" name="last_name" />
      <input type="submit" />
    </form>
  </body>
</html>
```

Go server code:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// Access the form data
		fmt.Fprintln(w, r.PostForm) // Outputs form data as key-value pairs
	})
	server.ListenAndServe()
}
```

</details>

When the form is submitted, `r.PostForm` will contain the key-value pairs from the form.

**Handling File Uploads**

The `multipart/form-data` encoding is commonly used for uploading files. To handle uploaded files, follow these steps:

1. Call `ParseMultipartForm` to parse the request.
2. Retrieve the uploaded file's header using the `File` field.
3. Open the file using the file header’s `Open` method.
4. Read the file contents using `ioutil.ReadAll` or other methods.

Alternatively, you can use `FormFile` for single-file uploads. This approach does not require calling `ParseMultipartForm` and is faster when handling one file.

<details><summary><b> Code</b></summary>

HTML form for file upload:

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>File Upload</title>
  </head>
  <body>
    <form
      action="http://localhost:8080/process"
      method="post"
      enctype="multipart/form-data"
    >
      Upload File:
      <input type="file" name="uploaded" />
      <input type="submit" />
    </form>
  </body>
</html>
```

</details>

**Method 1: Using ParseMultipartForm**

<details><summary><b> Code</b></summary>

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	fileHeader := r.MultipartForm.File["uploaded"][0] // Get the first uploaded file
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file) // Read the file into a byte slice
		if err == nil {
			fmt.Fprintln(w, string(data)) // Print the file content
		}
	}
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
```

</details>

**Method 2: Using FormFile**

An alternative and simpler method is to use `FormFile`, which retrieves the first file uploaded with a specific key without calling `ParseMultipartForm`.

<details><summary><b> Code</b></summary>

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// Use FormFile to get the uploaded file; no need to call ParseMultipartForm
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
```

</details>

**MultipartReader**

For handling `multipart/form-data` requests as a stream, you can use the `MultipartReader` method instead of `ParseMultipartForm`. This approach reads the form data as a stream instead of parsing it all at once, which can be useful for large file uploads.

**Tips**

- **ParseForm** and **ParseMultipartForm**: Used to parse URL and form data in POST requests.
- **Form**, **PostForm**, and **MultipartForm**: Fields in the `Request` struct for accessing form data.
- **FormFile**: Simplifies file retrieval, especially for single file uploads.
- **MultipartReader**: For handling large `multipart/form-data` as a stream.

This allows you to handle various form submissions and file uploads effectively in Go web applications using the built-in libraries.

---

## **2. HTTP Responses**

In Go, sending responses from the server to the client is a fundamental part of web development. The `ResponseWriter` interface plays a central role in this process. This part explains how to use `ResponseWriter` to send different types of responses, manage headers, and handle built-in response utilities effectively.

**ResponseWriter Overview**

The `ResponseWriter` interface is used by HTTP handlers to construct responses to clients. Behind the scenes, it is supported by the non-exported `http.response` struct, making it accessible only through the `ResponseWriter` interface.

**Writing to ResponseWriter**

The `Write` method takes a `byte` slice as its parameter and writes it to the HTTP response body. If the content type is not explicitly set in the header before calling `Write`, the first 512 bytes of the written data will be used to detect and set the content type.

<details><summary><b> Code</b></summary>

```go
package main

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Go Web</title>
</head>
<body>
    Hello, World!
</body>
</html>`
	w.Write([]byte(str)) // Convert string to byte slice and write to body
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/write", writeExample)
	server.ListenAndServe()
}
```

</details>

In this example, a simple HTML page is sent as the response body.

### **Setting Status Codes with WriteHeader**

The `WriteHeader` method allows you to set an HTTP status code explicitly. If `WriteHeader` is not called, the status defaults to `http.StatusOK (200)` when the `Write` method is first invoked.

- **Usage:**  
  `WriteHeader` is typically used to send error codes or redirect statuses.
- **Important:**  
  Once `WriteHeader` is called, headers can no longer be modified.

<details><summary><b> Code</b></summary>

```go
package main

import "net/http"

func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://apple.com")
	w.WriteHeader(302) // Redirect with HTTP 302
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/redirect", headerExample)
	server.ListenAndServe()
}
```

</details>

In this example, the client is redirected to "http://apple.com" with a 302 status code.

**Managing Headers**

The `Header` method of `ResponseWriter` returns a map of headers, allowing modifications. Any changes made before `WriteHeader` or `Write` are reflected in the HTTP response.

<details><summary><b> Code</b></summary>

```go
package main

import (
	"encoding/json"
	"net/http"
)

type POST struct {
	User   string
	Thread []string
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set content type to JSON
	post := &POST{
		User:   "lord",
		Thread: []string{"first", "second", "third"},
	}
	jsonData, _ := json.Marshal(post)
	w.Write(jsonData) // Write JSON data to response body
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
```

</details>

This example creates a JSON response containing a simple data structure.

**Built-in Response Utilities**

Go provides several built-in functions to simplify common response scenarios:

1. **`NotFound`**  
   Sends a 404 status code with optional additional information.  
   **Usage:**

   ```go
   http.NotFound(w, r)
   ```

2. **`ServeFile`**  
   Serves a file from the filesystem.  
   **Usage:**

   ```go
   http.ServeFile(w, r, "path/to/file")
   ```

3. **`ServeContent`**  
   Serves content from any object that implements `io.ReadSeeker`. This function supports range requests, allowing partial content responses.  
   **Usage:**

   ```go
   http.ServeContent(w, r, "example.txt", modTime, file)
   ```

4. **`Redirect`**  
   Sends a redirect response to the client with a specific status code.  
   **Usage:**
   ```go
   http.Redirect(w, r, "http://new-url.com", http.StatusMovedPermanently)
   ```

---

## **3. Templates**

Templates are a powerful feature in Go for generating dynamic HTML content. They allow you to combine static HTML with dynamic data, producing customized responses for web applications. Go's standard library provides two template engines: **text/template** (for text-based templates) and **html/template** (designed specifically for HTML). Most Go web frameworks use these libraries as their default template engines.

**What Are Templates?**

A web template is a pre-designed HTML structure that can be repeatedly used to generate dynamic HTML pages. Template engines merge a predefined template with contextual data, producing the final output. In Go, this involves parsing the template and rendering it with data.

**Basic Usage of Templates**

To use a template in Go, follow two steps:

1. **Parse the Template Source**  
   Templates can be parsed from files, strings, or patterns to create a `template` struct.
2. **Execute the Parsed Template**  
   The parsed template is executed by combining it with data. The resulting output is written to the `ResponseWriter`.

The following example shows how to use Go's template engine to generate HTML output dynamically:

- The `template.ParseFiles` method parses the `tmpl.html` file and creates a template struct.
- The `t.Execute` method renders the template, replacing `{{.}}` with the provided data ("Hello Kitty" in this case).

<details><summary><b> Code</b></summary>

```go
package main

import (
	"html/template"
	"net/http"
)

// Custom handler for processing templates
func process(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	t, _ := template.ParseFiles("tmpl.html")
	// Execute the template and pass data to replace `{{.}}`
	t.Execute(w, "Hello Kitty")
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/test", process)
	server.ListenAndServe()
}
```

</details>

**Understanding Template Actions**

Actions are enclosed within double curly braces (`{{}}`) and instruct the template engine to perform specific operations. For instance, `{{.}}` represents the current data context and will be replaced by the provided value during execution.

**Parsing Templates**

Go offers various ways to parse templates:

1. **`ParseFiles`**  
   Parses one or more files and returns the first template.
   ```go
   t, _ := template.ParseFiles("file1.html", "file2.html")
   ```
2. **`ParseGlob`**  
   Uses pattern matching to parse multiple files.
   ```go
   t, _ := template.ParseGlob("templates/*.html")
   ```
3. **`Parse`**  
   Parses a template string directly.
   ```go
   t, _ := template.New("example").Parse("Hello, {{.}}!")
   ```

<details><summary><b> Code</b></summary>

Parsing Multiple Templates

```go
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// Parse multiple templates
	templates, _ := template.ParseFiles("template1.html", "template2.html")
	// Execute a specific template from the set
	templates.ExecuteTemplate(w, "template2.html", "Hello, Go!")
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
```

</details>

In this example:

- `template.ParseFiles` parses multiple files (`template1.html` and `template2.html`).
- The `ExecuteTemplate` method renders a specific template (`template2.html`).

**Executing Templates**

Go provides two primary methods for executing templates:

1. **`Execute`**  
   Executes a single template or the first template in a set.

   ```go
   t.Execute(w, data)
   ```

2. **`ExecuteTemplate`**  
   Executes a specific template in a set by name.
   ```go
   t.ExecuteTemplate(w, "template_name", data)
   ```

---

## **4. Connecting to a Database**

Connecting to a database in Go involves loading a database driver and using Go’s `sql` package for interaction. The `sql` package provides robust tools for establishing and managing connections, querying data, and performing CRUD operations. This guide explains these concepts with examples.

**Loading the Driver**

To interact with a specific SQL database, a driver is required. The driver contains the logic necessary to communicate with the database. Drivers are registered automatically when their packages are imported.

For example, to use Microsoft SQL Server, install its driver using:

```bash
go get github.com/denisenkom/go-mssqldb
```

Then import the package in your Go program:

```go
import _ "github.com/denisenkom/go-mssqldb"
```

The underscore (`_`) ensures the package is imported only for its side effects (self-registration) and is not used directly in the code.

**Connecting with `sql.Open`**

Use the `sql.Open` function to create a connection pool. It takes two parameters: the driver name and the data source name (DSN). The returned `*sql.DB` struct represents the connection pool and is thread-safe, making it suitable for concurrent use.

```go
db, err := sql.Open("sqlserver", connStr)
if err != nil {
    log.Fatalln(err.Error())
}
```

**Important Notes:**

- `sql.Open` does not establish an actual connection immediately. It initializes the connection lazily.
- Use `db.PingContext` to verify the connection.

<details><summary><b> Code</b></summary>

Full Example: Connecting to SQL Server

```go
package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB // Global database connection pool

const (
	server   = "example.server"
	port     = 1433
	user     = "username"
	password = "password"
	database = "go-db"
)

func main() {
	// Connection string
	connStr := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error
	db, err = sql.Open("sqlserver", connStr) // Initialize connection pool
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil { // Verify connection
		log.Fatalln(err.Error())
	}

	fmt.Println("Connected to the database!")
}
```

</details>

**Querying the Database**

The `sql.DB` struct provides methods to execute queries:

- **`Query`**: Executes a query that returns multiple rows.
- **`QueryRow`**: Executes a query that returns a single row.
- **`QueryContext`** and **`QueryRowContext`**: Context-aware variants that support timeouts, cancellations, or additional context values.

<details><summary><b> Code</b></summary>

Querying a Single Row

```go
func getOne(id int) (app, error) {
	a := app{}
	err := db.QueryRow("SELECT id, name, status, level, [order] FROM dbo.App WHERE id = @id",
		sql.Named("id", id)).Scan(&a.ID, &a.Name, &a.Status, &a.Level, &a.Order)
	return a, err
}
```

Querying Multiple Rows

```go
func getMany(id int) ([]app, error) {
	rows, err := db.Query("SELECT id, name, status, level, [order] FROM dbo.App WHERE id > @id",
		sql.Named("id", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []app
	for rows.Next() {
		a := app{}
		if err := rows.Scan(&a.ID, &a.Name, &a.Status, &a.Level, &a.Order); err != nil {
			return nil, err
		}
		apps = append(apps, a)
	}
	return apps, rows.Err()
}
```

Model for Mapping Database Rows

```go
type app struct {
	ID     int
	Name   string
	Status int
	Level  int
	Order  int
}
```

</details>

**Performing CRUD Operations**

- Inserting Data

Use `Exec` or `Prepare` to execute an `INSERT` statement. Use `SCOPE_IDENTITY()` in SQL Server to retrieve the last inserted ID.

<details><summary><b> Code</b></summary>

```go
func (a *app) Insert() error {
	statement := `INSERT INTO dbo.App (name, status, level, [order])
		VALUES (@Name, @Status, @Level, @Order);
		SELECT ISNULL(SCOPE_IDENTITY(), -1);`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(
		sql.Named("Name", a.Name),
		sql.Named("Status", a.Status),
		sql.Named("Level", a.Level),
		sql.Named("Order", a.Order)).Scan(&a.ID)
}
```

</details>

- Updating Data

<details><summary><b> Code</b></summary>

```go
func (a *app) Update() error {
	_, err := db.Exec("UPDATE dbo.App SET name=@Name, [order]=@Order WHERE id=@ID",
		sql.Named("Name", a.Name), sql.Named("Order", a.Order), sql.Named("ID", a.ID))
	return err
}
```

</details>

- Deleting Data

<details><summary><b> Code</b></summary>

```go
func (a *app) Delete() error {
	_, err := db.Exec("DELETE FROM dbo.App WHERE id=@ID", sql.Named("ID", a.ID))
	return err
}
```

</details>

**Best Practices**

1. **Use Contexts:** Prefer `Context` variants like `QueryContext` and `ExecContext` for better control over query execution.
2. **Close Rows:** Always call `rows.Close()` after iterating through query results to release resources.
3. **Handle Errors Gracefully:** Check for errors after every database operation.
4. **Use Connection Pooling:** The `*sql.DB` struct manages a connection pool automatically, so reuse it across your application.
5. **Secure Connections:** Ensure sensitive data like connection strings are encrypted or stored securely.

With these tools, you can effectively connect to, query, and manage SQL databases in Go, enabling robust data-driven applications.

---

# **5. Route**

Routing in Go involves mapping different URL paths to specific handlers, allowing the application to respond dynamically based on the requested path. While early implementations often place routing logic directly in the `main` function, this approach becomes less manageable as a project grows. A better practice is to introduce a **controller layer**, separating concerns and improving maintainability.

**Role of the Controller Layer**

The controller layer handles routing and directs incoming requests to the appropriate handlers or static resources. This structured approach ensures that the application is easier to navigate, debug, and extend.

### **Project Structure Overview**

- **`main` function:** Focuses on application setup tasks, such as starting the server and initializing routes.
- **Controller Layer:** Manages static resources (e.g., CSS, JS, images) and routes incoming requests to specific handlers.

All requests first pass through a primary controller and are then distributed to appropriate handlers for processing.

**Routing Types**

**Static Routes**

A static route maps a fixed path to a specific page or handler. For example:

```go
http.HandleFunc("/about", aboutHandler)
```

This always serves the same content when the `/about` path is requested.

**Dynamic Routes with Parameters**

Dynamic routes allow parameters to be passed via the URL. For example:

```go
http.HandleFunc("/user/{id}", userHandler)
```

This enables creating pages with the same template but different data based on the URL parameter (e.g., user profiles).

**Using Third-Party Routers**

Go's standard library provides basic routing capabilities, but for more advanced routing features, third-party routers are often preferred. Two popular options are:

1. **Gorilla Mux**
   - Highly flexible and feature-rich.
   - Supports complex routing patterns and middleware.
   - Slightly slower compared to simpler alternatives.
2. **HttpRouter**
   - Focused on performance.
   - Lightweight and straightforward.
   - Ideal for applications requiring high-speed routing with minimal overhead.

**Controller Implementation Example**

Here’s how to structure routing with a controller layer.

- Main Function: Setting Up Routes

The `main` function delegates routing responsibilities to the controller layer:

<details><summary><b> Code</b></summary>

```go
package main

import (
	"your_project/controller"
	"net/http"
)

func main() {
	controller.RegisterRoutes() // Register routes from the controller layer

	// Start the server
	http.ListenAndServe(":8080", nil)
}
```

</details>

- Controller: Registering Routes\*\*

The controller organizes and registers routes for different sections of the application:

<details><summary><b> Code</b></summary>

```go
package controller

// Entry point for registering routes
func RegisterRoutes() {
	// Static resources
	registerStaticResources()

	// Application-specific routes
	registerHomeRoutes()
	registerAboutRoutes()
	registerContactRoutes()
}
```

</details>

Each route registration function defines specific routes for its area:

<details><summary><b> Code</b></summary>

```go
func registerHomeRoutes() {
	http.HandleFunc("/", homeHandler)
}

func registerAboutRoutes() {
	http.HandleFunc("/about", aboutHandler)
}

func registerContactRoutes() {
	http.HandleFunc("/contact", contactHandler)
}
```

</details>

- Handlers

Each handler processes a specific type of request:

<details><summary><b> Code</b></summary>

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Home Page!"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Us"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Contact Us"))
}
```

</details>

**Why Use a Controller Layer?**

1. **Separation of Concerns:** Keeps routing logic organized and distinct from application setup.
2. **Scalability:** Makes it easier to add new features and routes without bloating the `main` function.
3. **Reusability:** Encourages modular design, where individual route groups can be reused or modified independently.

---

## **6. Middleware Enhancing**

Middleware in Go acts as a layer between incoming HTTP requests and their designated handlers. It intercepts requests, processes them, and decides whether to pass them to the handler or return a response. Similarly, it can modify the response before sending it back to the client. Middleware is commonly used for tasks such as logging, security, request timeout handling, and response compression.

**Purpose of Middleware**

Middleware serves various roles in web applications, including:

1. **Logging:** Tracks requests to analyze usage patterns or troubleshoot issues.
2. **Security:** Verifies user identity and permissions (e.g., authentication).
3. **Request Timeout:** Ensures that requests do not hang indefinitely.
4. **Response Compression:** Optimizes the response size to improve performance.

**Main Application**

Here’s how to implement middleware in Go, step by step:

In this example, a simple HTTP server is created, which serves information about a company. The server uses an authentication middleware to secure the `/companies` endpoint.

<details><summary><b> Code</b></summary>

```go
package main

import (
	"encoding/json"
	"go-web-zero/p23/middleware"
	"net/http"
)

type Company struct {
	ID      int
	Name    string
	Country string
}

func main() {
	// Define the handler for the /companies endpoint
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}
		enc := json.NewEncoder(w)
		enc.Encode(c)
	})

	// Wrap the server with authentication middleware
	http.ListenAndServe("localhost:8080", new(middleware.AuthMiddleware))
}
```

</details>

**Authentication Middleware**

The authentication middleware ensures that requests contain an `Authorization` header. If the header is absent, the middleware denies access by returning an `HTTP 401 Unauthorized` response. If the header is present, it forwards the request to the next handler.

<details><summary><b> Code</b></summary>

```go
package middleware

import "net/http"

// AuthMiddleware represents the middleware structure with a Next handler
type AuthMiddleware struct {
	Next http.Handler
}

// ServeHTTP processes each request through the middleware
func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// If no next handler is defined, use the default router
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}

	// Check the Authorization header
	auth := r.Header.Get("Authorization")
	if auth != "" {
		// Before routing: Custom logic can go here

		// Forward the request to the next handler
		am.Next.ServeHTTP(w, r)

		// After routing: Additional logic can be executed here
	} else {
		// Return unauthorized if the Authorization header is missing
		w.WriteHeader(http.StatusUnauthorized)
	}
}
```

</details>

**How It Works**

1. **Chaining Requests:** Middleware is designed as a chain. Each middleware has a `Next` field, which points to the next handler or middleware in the chain. If no `Next` handler is specified, the `DefaultServeMux` is used to process the request.

2. **Intercepting Requests:**

   - Before forwarding the request, middleware can perform checks or preprocessing (e.g., validate headers, log requests).
   - After the next handler processes the request, middleware can modify the response or perform additional actions.

3. **Authorization Check:** In the example, the `AuthMiddleware` checks for the presence of the `Authorization` header. Requests without this header receive an `HTTP 401 Unauthorized` status.

**Testing the Middleware**

```go
GET http://localhost:8080/companies HTTP/1.1

# with auth
GET http://localhost:8080/companies HTTP/1.1
Authorization: abcdefgo
```

**Benefits of Middleware**

- **Modularity:** Middleware decouples cross-cutting concerns (e.g., logging, authentication) from the core application logic.
- **Reusability:** Middleware can be reused across multiple projects or endpoints.
- **Extensibility:** New functionality can be easily added by introducing additional middleware layers.

## Conclusion

Go's `net/http` package simplifies building robust web applications with routing, request handling, dynamic templates, and middleware integration, enabling efficient, secure, and modular development.

<div align="center"> 
<img src="images/02_Go_Web.gif" width="50%" />
</div>
