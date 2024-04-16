package main

import (
	"github.com/allansdefreitas/api-go/gin/database"
	"github.com/allansdefreitas/api-go/gin/routes"
)

func main() {
	database.ConnectToDatabase()

	routes.HandleRequests()
}
