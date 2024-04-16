package main

import (
	"log"
	"net/http"

	"github.com/allansdefreitas/star_shop/routes"
	_ "github.com/lib/pq"
)

// the _ char acter says that this package will only be used on running time

/*
go mod init projectName
go get github.com/lib/pq 
*/

func main() {
	log.Println("MAIN")
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
