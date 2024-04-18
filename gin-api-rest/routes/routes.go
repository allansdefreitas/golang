package routes

import (
	"github.com/allansdefreitas/api-go/gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()              // Default conf
	r.LoadHTMLGlob("templates/*")   // Specify that we have HTML pages
	r.Static("/assets", "./assets") // Specify that we have static files

	r.GET("/students", controllers.ShowAllStudents) //Gin calls the function 'showAllStudents' internally,
	r.GET("/:name", controllers.Greeting)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	// r.PUT("students", controllers.UpdateStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)

	//HTML pages
	r.GET("/greetings", controllers.ShowHTMLGreetings)
	r.GET("/index", controllers.ShowPageIndex)
	r.NoRoute(controllers.RouteNotFound) //handler not found requests

	r.Run(":8000") //Run a server. Default is :8080

}
