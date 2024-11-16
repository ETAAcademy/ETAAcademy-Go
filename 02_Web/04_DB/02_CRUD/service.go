package main

import (
	"database/sql"
	"log"
)

func getOne(id int) (a app, err error) {
	a = app{}

	// [order] is a SQL keyword, so it must be enclosed in brackets
	err = db.QueryRow("select id, name, status,"+
		"level, [order] from dbo.App").Scan(
		&a.ID, &a.name, &a.status, &a.level, &a.order)
	return
}

// Returns a slice of results
func getMany(id int) (apps []app, err error) {

	// [order] is a SQL keyword, so it must be enclosed in brackets
	rows, err := db.Query("select id, name, status,"+
		"level, [order] from dbo.App where id > @id", sql.Named("Id", id))

	// Iterate over the rows
	for rows.Next() {
		a := app{}
		err = rows.Scan(
			&a.ID, &a.name, &a.status, &a.level, &a.order)
		if err != nil {
			log.Fatalln(err.Error())
		}
		apps = append(apps, a)
	}
	return apps, err
}

func (a *app) Update() (err error) {
	// Update the record with new values
	_, err = db.Exec("UPDATE dbo.App SET Name=@name, [Order]=@Order WHERE Id=@Id", sql.Named("Name", a.name), sql.Named("Order", a.order), sql.Named("Id", a.ID))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func (a *app) Delete() (err error) {
	// Delete the record by ID
	_, err = db.Exec("DELETE FROM dbo.App WHERE Id=@Id", sql.Named("Id", a.ID))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}

func (a *app) Insert() (err error) {
	// Insert a new record into the database
	statement := `INSERT INTO dbo.App(
	Name,
	NickName,
	Status,
	Level,
	[Order],
	Pinyin)
	VALUES
	(@Name,
	@Status,
	@Level, 
	@Order,
	'...');
	SELECT isNull(SCOPE_IDENTITY(), -1);` // Get the ID of the newly inserted record

	stmt, err := db.Prepare(statement) // Prepare the SQL statement, stmt must be closed after use
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer stmt.Close()

	// Execute the query and scan the resulting ID
	err = stmt.QueryRow(
		sql.Named("Name", a.name),
		sql.Named("Status", a.status),
		sql.Named("Level", a.level),
		sql.Named("Order", a.order)).Scan(
		&a.ID,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}
	return
}
