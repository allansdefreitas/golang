package database

import (
	"log"

	"github.com/allansdefreitas/api-go/gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {

	host := "localhost"
	database := "root"
	user := "root"
	password := "root"
	port := "5432"

	connectionString := "host=" + host + " user=" + user + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error during connect to database")
	}
	DB.AutoMigrate(&models.Student{}) //create a table based on an instance of an struct

}
