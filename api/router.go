package api

import "github.com/gorilla/mux"

func New() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	return router
}
