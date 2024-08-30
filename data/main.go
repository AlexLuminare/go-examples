package main

import (
	"database/sql"
)

type Product struct {
	Id    int
	Name  string
	Price float64
	Category
}

type Category struct {
	Id   int
	Name string
}

func queryDatabase(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`
	        SELECT Products.Id, Products.Name, Products.Price,
	
	
	                Categories.Id as Cat_Id, Categories.Name as CatName 
	                FROM Products, Categories
	        WHERE Products.Category = Categories.Id
	            AND Products.Id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price,
			&p.Category.Id, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}
	return
}

func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`
	        INSERT INTO Products (Name, Category, Price)
	        VALUES (?, ?, ?)`, p.Name, p.Category.Id, p.Price)
	if err == nil {
		id, err = res.LastInsertId()
		if err != nil {
			Printfln("Result error: %v", err.Error())
		}
	} else {
		Printfln("Exec error: %v", err.Error())
	}
	return
}

func main() {
	db, err := openDatabase()
	if err == nil {
		newProduct := Product{Name: "Stadium", Category: Category{Id: 2}, Price: 79500}
		newID := insertRow(db, &newProduct)
		p := queryDatabase(db, int(newID))
		Printfln("New Product: %v", p)
		db.Close()
	} else {
		panic(err)
	}
}
