package models

import (
	"log"

	"github.com/allansdefreitas/star_shop/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func FetchAllProducts() []Product {

	db := db.ConnectToDataBase()

	var query string = "select * from products ORDER BY id asc"
	statementSelect, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	tempProduct := Product{}
	products := []Product{}

	for statementSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := statementSelect.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		tempProduct.Id = id
		tempProduct.Name = name
		tempProduct.Description = description
		tempProduct.Price = price
		tempProduct.Quantity = quantity

		products = append(products, tempProduct)
	}

	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	log.Println("MODEL:CREATE")
	db := db.ConnectToDataBase()

	query := `INSERT INTO products
	(prod_name, prod_description, prod_price, prod_quantity) 
	VALUES ($1, $2, $3, $4)`
	statementInsert, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	statementInsert.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(productId string) {
	db := db.ConnectToDataBase()

	query := "DELETE FROM products WHERE id=$1"
	statementDelete, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	statementDelete.Exec(productId)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectToDataBase()

	query := "SELECT * FROM products WHERE id=$1"
	productFromDatabase, err := db.Query(query, id)
	if err != nil {
		panic(err.Error())
	}

	tempProduct := Product{}

	for productFromDatabase.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productFromDatabase.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		tempProduct.Id = id
		tempProduct.Name = name
		tempProduct.Description = description
		tempProduct.Price = price
		tempProduct.Quantity = quantity

	}
	defer db.Close()
	return tempProduct
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	log.Println("MODEL:UPDATE")
	db := db.ConnectToDataBase()

	query := `UPDATE products
	SET prod_name = $2,
	prod_description = $3,
	prod_price = $4,
	prod_quantity = $5
	WHERE id = $1`
	statementUpdate, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	statementUpdate.Exec(id, name, description, price, quantity)
	defer db.Close()
}
