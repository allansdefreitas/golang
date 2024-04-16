package routes

import (
	"log"
	"net/http"

	"github.com/allansdefreitas/go-rest-api/controllers"
	"github.com/allansdefreitas/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()

	r.Use(middleware.SetContentTypeMiddleware)

	// r.http.HandleFunc("/", controllers.Home)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetOnePersonality).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonality).Methods("Put")
	r.HandleFunc("/api/personalities/{id}", controllers.RemovePersonality).Methods("Delete")
	// log.Fatal(http.ListenAndServe(":8000", r))
	// to allow CORS to all (*), but it could be specifically 'localhost:8000', for instance
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
