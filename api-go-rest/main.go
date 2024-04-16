package main

import (
	"fmt"

	"github.com/allansdefreitas/go-rest-api/database"
	"github.com/allansdefreitas/go-rest-api/models"
	"github.com/allansdefreitas/go-rest-api/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{Id: 1, Name: "Filipe Camarão", History: "History 1"},
		{Id: 2, Name: "André Vidal de Negreiros", History: "History 2"},
		{Id: 3, Name: "Henrique Dias", History: "History 3"},
		{Id: 4, Name: "Francisco Barreto de Meneses", History: "History 4"},
	}

	database.ConnectToDatabase()
	fmt.Println("Starting Go REST server")
	routes.HandleRequest()
}
