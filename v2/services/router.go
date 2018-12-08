package services

import "github.com/gorilla/mux"

// NewRouter Returns a Mux Router after registering all the
// routes.
func NewRouter() *mux.Router {
	router := mux.
		NewRouter().
		StrictSlash(true).
		PathPrefix("/" + VERSION).
		Subrouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
