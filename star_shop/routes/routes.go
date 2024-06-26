package routes

import (
	"log"
	"net/http"

	"github.com/allansdefreitas/star_shop/controllers"
)

func LoadRoutes() {
	log.Println("ROUTES")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)

	http.HandleFunc("/delete", controllers.Delete)

	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
