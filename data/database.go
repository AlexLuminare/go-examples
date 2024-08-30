package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		Printfln("Driver: %v", driver)
	}
}
func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "products.db")
	if err == nil {
		Printfln("Opened database")
	}
	return
}
