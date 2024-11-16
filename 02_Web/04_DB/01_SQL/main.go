package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// Database driver, _ registers automatically during init. Using "_" means the package is imported but not directly used in the code.
	// No variables, does not affect the logic inside the code.
	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB // Points to the database

func main() {
	// Connection string
	connStr := fmt.Sprintf("server=%s;user id = %s; "+
		""+
		"password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	fmt.Println(connStr)

	// Open the database connection
	db, err := sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx := context.Background()

	// Ping the database to check the connection
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Test if the connection is successful
	fmt.Println("connected!")
}

const (
	server   = "xxxx.xxx"
	port     = 1433
	user     = "xxx"
	password = "123"
	database = "go-db"
)
