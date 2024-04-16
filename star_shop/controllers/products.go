package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/allansdefreitas/star_shop/models"
)

var templ = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {

	allProducts := models.FetchAllProducts()
	templ.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	log.Println("TEMPLATE:NEW")
	templ.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER:INSERT")
	if r.Method == "POST" {

		name := r.FormValue("name") //Get data from page
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			// panic(err.Error())
			log.Println("Error during casting of price:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during casting of quantity:", err)
		}

		models.CreateNewProduct(name, description, priceConverted, quantityConverted)

	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	log.Println("TEMPLATE:EDIT")

	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)
	templ.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	log.Println("CONTROLLER:INSERT")
	if r.Method == "POST" {

		//Get data from page
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConverted, err := strconv.ParseFloat(price, 64)
		if err != nil {
			// panic(err.Error())
			log.Println("Error during casting of price:", err)
		}

		quantityConverted, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error during casting of quantity:", err)
		}

		idConverted, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error during casting of id:", err)
		}

		models.UpdateProduct(idConverted, name, description, priceConverted, quantityConverted)

	}

	http.Redirect(w, r, "/", 301)
}

// main...routes..controller...model
