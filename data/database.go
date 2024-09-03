package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var insertNewCategory *sql.Stmt
var changeProductCategory *sql.Stmt

func listDrivers() {
	for _, driver := range sql.Drivers() {
		Printfln("Driver: %v", driver)
	}
}
func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "products.db")
	if err == nil {
		Printfln("Opened database")
		insertNewCategory, _ = db.Prepare("INSERT INTO Categories (Name) VALUES (?)")
		changeProductCategory, _ = db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")
	}
	return
}
