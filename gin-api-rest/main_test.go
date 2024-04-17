package main

/**
Author: Allan Freitas
Comments:

#Run a specific test function:
go test -run <testname>

*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/allansdefreitas/api-go/gin/controllers"
	"github.com/allansdefreitas/api-go/gin/database"
	"github.com/allansdefreitas/api-go/gin/errors"
	"github.com/allansdefreitas/api-go/gin/models"
	constants "github.com/allansdefreitas/api-go/gin/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) //To show test results at terminal in a better way
	routes := gin.Default()
	return routes
}

// func TestFailure(t *testing.T) {
// 	t.Fatalf("Test failed successfully")
// }

/*
This same test can be done in Postman, with:

	pm.test("Status code is 200", function () {
	    pm.response.to.have.status(200);
	});

	pm.test("Verifyng the response content", function () {
	    pm.response.to.have.body('{"API says: ":"Whats Up, Allan. How are you?"}');
	});
*/
func TestGreetingWithParamVerifyStatusCode200(t *testing.T) {

	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)
	// Params: HTTP Method, Endpoint, Request Body
	req, _ := http.NewRequest(constants.GET, "/Allan", nil)

	// this variable will store the all content of response
	response := httptest.NewRecorder()

	// Do request
	r.ServeHTTP(response, req)

	// if response.Code != http.StatusOK {
	// 	t.Fatalf("Status error: Received status code was %d and the expected was %d",
	// 		response.Code, http.StatusOK)
	// }

	// Test the STATUS CODE
	//test, expected, actual
	assert.Equal(t, http.StatusOK, response.Code, "They should be equal")

	//Test the BODY Using testfy
	responseMock := `{"API says: ":"Whats Up, Allan. How are you?"}`

	responseBody, _ := io.ReadAll(response.Body)
	responseBodyString := string(responseBody)

	assert.Equal(t, responseMock, responseBodyString)
}

func TestHandlerShowAllStudents(t *testing.T) {

	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent() // Run this line at the end of function
	r := SetupTestRoutes()
	r.GET("/students", controllers.ShowAllStudents)
	req, _ := http.NewRequest(constants.GET, "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	// fmt.Println(response.Body)

}

var IDMockStudent int //What abou clean code here? Dirty code here!

func CreateMockStudent() {
	var student models.Student
	student.Name = "John Doe"
	student.CPF = "10022255588"
	student.RG = "4561237"

	database.DB.Create(&student)
	IDMockStudent = int(student.ID)

}
func CreateMockStudentID() int {
	var student models.Student
	student.Name = "John Doe"
	student.CPF = "10022255588"
	student.RG = "4561237"

	database.DB.Create(&student)
	return int(student.ID)
}

func CreateMockStudentIDStruct() (int, models.Student) {
	var student models.Student
	student.Name = "John Doe"
	student.CPF = "10022255588"
	student.RG = "4561237"

	database.DB.Create(&student)
	var intId = int(student.ID)
	return intId, student
}

func DeleteMockStudent() {

	var student models.Student
	database.DB.Delete(&student, IDMockStudent)
}

func TestHandlerFindStudentByCPF(t *testing.T) {

	database.ConnectToDatabase()
	CreateMockStudent()
	defer DeleteMockStudent()
	r := SetupTestRoutes()

	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)
	givenCPF := "99988855563"
	req, _ := http.NewRequest(constants.GET, "/students/cpf/"+givenCPF, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	// fmt.Println(response.Body)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestHandlerFindStudentById(t *testing.T) {
	database.ConnectToDatabase()
	var id int = CreateMockStudentID()
	defer DeleteMockStudent()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.FindStudentById)

	path := "/students/" + strconv.Itoa(id)
	req, err := http.NewRequest(constants.GET, path, nil)
	errors.HandleError(err)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMock models.Student

	json.Unmarshal(response.Body.Bytes(), &studentMock)
	fmt.Println(studentMock.Name, id)

	var expectedName = "John Doe"
	var expectedRG = "4561237"
	var expectedCPF = "10022255588"
	assert.Equal(t, expectedName, studentMock.Name, "Names must be equal")
	assert.Equal(t, expectedRG, studentMock.RG, "RG numbers must be equal")
	assert.Equal(t, expectedCPF, studentMock.CPF, "CPF numbers must be equal")
	assert.Equal(t, http.StatusOK, response.Code, "Status codes must be equal")
	// if err != nil {
	// 	panic(err.Error())
	// }
}

/* Cleaner code */
func TestHandlerFindStudentByIdAllan(t *testing.T) {
	database.ConnectToDatabase()
	var id, studentMockExpected = CreateMockStudentIDStruct()
	defer DeleteMockStudent()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.FindStudentById)

	path := "/students/" + strconv.Itoa(id)
	req, err := http.NewRequest(constants.GET, path, nil)
	errors.HandleError(err)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMockObtained models.Student

	json.Unmarshal(response.Body.Bytes(), &studentMockObtained)
	fmt.Println(studentMockObtained.Name, id)

	assert.Equal(t, studentMockExpected.Name, studentMockObtained.Name, "Names must be equal")
	assert.Equal(t, studentMockExpected.RG, studentMockObtained.RG, "RG numbers must be equal")
	assert.Equal(t, studentMockExpected.CPF, studentMockObtained.CPF, "CPF numbers must be equal")
	assert.Equal(t, http.StatusOK, response.Code, "Status codes must be equal")
}

func TestHandlerDeleteStudent(t *testing.T) {
	database.ConnectToDatabase()
	r := SetupTestRoutes()
	var id, _ = CreateMockStudentIDStruct()
	r.DELETE("/students/:id", controllers.DeleteStudent)
	path := "/students/" + strconv.Itoa(id)
	req, err := http.NewRequest(constants.DELETE, path, nil)
	errors.HandleError(err)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Status codes must be equal")
}

func TestHandlerUpdateStudent(t *testing.T) {
	database.ConnectToDatabase()
	r := SetupTestRoutes()
	r.PATCH("students/:id", controllers.UpdateStudent)
	var _, studentMockCreated = CreateMockStudentIDStruct()
	id := int(studentMockCreated.ID)
	defer DeleteMockStudent()

	studentMockExpected := models.Student{Name: "John Huss", CPF: "99922255588", RG: "9991237"}

	jsonValue, err := json.Marshal(studentMockExpected)
	errors.HandleError(err)

	// idString := strconv.FormatUint(uint64(studentMockCreated.ID), 10)
	path := "/students/" + strconv.Itoa(id)

	//Put json in the body of request and put it in bytes format
	req, err := http.NewRequest(constants.PATCH, path, bytes.NewBuffer(jsonValue))
	errors.HandleError(err)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var studentMockObtained models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMockObtained)

	assert.Equal(t, http.StatusOK, response.Code, "Status codes must be equal")
	assert.Equal(t, studentMockExpected.Name, studentMockObtained.Name, "Names must be equal")
	assert.Equal(t, studentMockExpected.RG, studentMockObtained.RG, "RG numbers must be equal")
	assert.Equal(t, studentMockExpected.CPF, studentMockObtained.CPF, "CPF numbers must be equal")

	fmt.Println(studentMockCreated.Name)
	fmt.Println(studentMockObtained.Name)
}
