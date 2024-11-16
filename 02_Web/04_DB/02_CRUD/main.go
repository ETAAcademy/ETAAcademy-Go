package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// Database driver, the "_" ensures the init function is called automatically to register the driver.
	// No variables are used, so it does not affect the logic in the code.
	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB // Pointer to the database

func main() {
	// Connection string
	connStr := fmt.Sprintf("server=%s;user id = %s; "+
		""+
		"password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	fmt.Println(connStr)

	var err error
	// The db needs to be global, so we cannot use := here
	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ctx := context.Background()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// Test if the connection to the database is successful
	fmt.Println("connected!")

	// Query
	a, _ := getOne(103)
	fmt.Println(a)
	a.name += "1234"
	a.order++

	err = a.Update()
	if err != nil {
		log.Fatalln(err.Error())
	}
	a1, _ := getOne(103)
	fmt.Println(a1)

	a2 := app{
		name:   "Test",
		order:  123,
		level:  10,
		status: 1,
	}

	err = a2.Insert()
	if err != nil {
		log.Fatalln(err.Error())
	}

	one, _ := getOne(a2.ID)
	fmt.Println(one)
}

const (
	server   = "xxxx.xxx"
	port     = 1433
	user     = "xxx"
	password = "123"
	database = "go-db"
)
