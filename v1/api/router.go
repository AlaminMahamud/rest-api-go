package api

import (
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(Logger(route.HandlerFunc, route.Name))
	}

	return router
}
