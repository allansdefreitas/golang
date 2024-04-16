package models

import (
 "allansdefreitas/star_shop/db"

)
type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func FetchAllProducts() []Product {

	db := ConnectToDataBase()

	var query string = "select * from products"
	selectAllProducts, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}

	tempProduct := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
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

}
