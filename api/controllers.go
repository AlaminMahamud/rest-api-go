package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	people := PopulateData()
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people := UpdateData(person)
	json.NewEncoder(w).Encode(people)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
