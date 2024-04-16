package routes

import (
	"github.com/allansdefreitas/api-go/gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()                              // Default conf
	r.GET("/students", controllers.ShowAllStudents) //Gin calls the function 'showAllStudents' internally,
	r.GET("/:name", controllers.Greeting)
	r.GET("/students/:id", controllers.FindStudentById)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	// r.PUT("students", controllers.UpdateStudent)
	r.PATCH("students/:id", controllers.UpdateStudent)

	// so does not need params

	r.Run(":8000") //Run a server. Default is :8080

}
