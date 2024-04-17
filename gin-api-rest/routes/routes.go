package routes

import (
	"github.com/allansdefreitas/api-go/gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()            // Default conf
	r.LoadHTMLGlob("templates/*") // Specify that we have HTML pages

	r.GET("/students", controllers.ShowAllStudents) //Gin calls the function 'showAllStudents' internally,
	r.GET("/:name", controllers.Greeting)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	// r.PUT("students", controllers.UpdateStudent)
	r.PATCH("/students/:id", controllers.UpdateStudent)
	r.GET("/index", controllers.ShowPageIndex)
	// so does not need params

	r.Run(":8000") //Run a server. Default is :8080

}
