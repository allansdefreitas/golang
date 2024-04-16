package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/allansdefreitas/go-rest-api/database"
	"github.com/allansdefreitas/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-type", "application/json")
	var personalities []models.Personality
	database.DB.Find(&personalities)

	json.NewEncoder(w).Encode(personalities)
	log.Println("personalitiesFound: " + personalities[0].Name)
	// json.NewEncoder(w).Encode(models.Personalities)

}

func GetOnePersonality(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	log.Println("GetOnePersonality")
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)

	// for _, personality := range models.Personalities {

	// 	log.Println("personality: " + personality.Name)

	// 	if strconv.Itoa(personality.Id) == id {
	// 		json.NewEncoder(w).Encode(personality)
	// 	}
	// }

}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {

	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	//Persists on database
	database.DB.Create(&personality)

	//return persisted object in JSON format
	json.NewEncoder(w).Encode(personality)

}

func RemovePersonality(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.First(&personality, id)

	//We only specify the entity personality (even it be nil) and its id. That's enough
	database.DB.Delete(&personality, id)

	json.NewEncoder(w).Encode(personality)

}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.First(&personality, id)

	json.NewDecoder(r.Body).Decode(&personality)

	//Persists on database
	database.DB.Save(&personality)

	//return persisted object in JSON format
	json.NewEncoder(w).Encode(personality)
}

// With PK field (id) sent in request body
func UpdatePersonality2(w http.ResponseWriter, r *http.Request) {

	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	//Persists on database
	database.DB.Save(&personality)

	//return persisted object in JSON format
	json.NewEncoder(w).Encode(personality)
}
