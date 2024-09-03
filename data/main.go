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

func queryDatabaseByID(db *sql.DB, id int) (p Product) {
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

func queryDatabaseByGroup(db *sql.DB, categoryName string) []Product {
	products := []Product{}
	rows, err := db.Query(`
	        SELECT Products.Id, Products.Name, Products.Price, 
	                Categories.Id as Cat_Id, Categories.Name as CatName
	                FROM Products, Categories
	        WHERE Products.Category = Categories.Id
	            AND Categories.Name = ?`, categoryName)
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price,
				&p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
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

func insertAndUseCategory(name string, productIDs ...int) {
	result, err := insertNewCategory.Exec(name)
	if err == nil {
		newID, _ := result.LastInsertId()
		for _, id := range productIDs {
			changeProductCategory.Exec(int(newID), id)
		}
	} else {
		Printfln("Prepared statement error: %v", err)
	}
}

func insertAndUseCategoryTransactioned(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if err == nil {
		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if err != nil || updatedFailed {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func main() {
	db, err := openDatabase()
	if err == nil {
		// newProduct := Product{Name: "Stadium", Category: Category{Id: 2}, Price: 79500}
		// newID := insertRow(db, &newProduct)
		// insertAndUseCategory("Misc Products", 2)
		insertAndUseCategoryTransactioned(db, "Misc Products", 2)
		p := queryDatabaseByID(db, 2)
		Printfln("Product: %v", p)
		insertAndUseCategoryTransactioned(db, "Category_2", 100)
		// p := queryDatabaseByID(db, int(newID))
		// Printfln("New Product: %v", p)
		db.Close()
	} else {
		panic(err)
	}
}
