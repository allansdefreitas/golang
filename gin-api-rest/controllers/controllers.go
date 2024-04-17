package controllers

import (
	"net/http"

	"github.com/allansdefreitas/api-go/gin/database"
	"github.com/allansdefreitas/api-go/gin/models"
	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	var studentsFromDatabase []models.Student
	database.DB.Find(&studentsFromDatabase)
	c.JSON(200, studentsFromDatabase)
}

// func ShowAllStudents(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"id":   "1",
// 		"name": "Abreu e Lima",
// 	})
// }

func FindStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student) // Return found student

}

func CreateNewStudent(c *gin.Context) {

	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	} // in case of error occuring

	// Validate student before insert
	err := models.ValidateStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student) //return JSON success message

}

func DeleteStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student

	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfuly",
	})
	// c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")

	database.DB.First(&student, id)
	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "Student not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error during updating of student",
		})
		return
	}

	// Validate student before update
	err := models.ValidateStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)

	// err := c.ShouldBindJSON(&student);

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error()})
	// 	return
	// } // in case of error occuring

}

func FindStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")

	//Find by where, get the first and put it in student variable
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API says: ": "Whats Up, " + name + ". How are you?",
	})
}

func ShowPageIndex(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "Welcome!",
	})
}
