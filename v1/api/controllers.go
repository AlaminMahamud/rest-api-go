package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPeople lists all the people that matches with the list with no limit and offset.
func GetPeople(w http.ResponseWriter, r *http.Request) {
	people := PopulateData()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(people)
}

// GetPerson returns only a single person
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json;charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
		}
	}
}

// CreatePerson creates a list of person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people := UpdateData(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(people)
}

// UpdatePerson udpates a person
func UpdatePerson(w http.ResponseWriter, r *http.Request) {}

// DeletePerson deletes a person
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
