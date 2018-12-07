package api

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/v1/people", GetPeople},
	Route{"Item", "GET", "/v1/people/{id}", GetPerson},
	Route{"Create", "POST", "/v1/people/{id}", CreatePerson},
	Route{"Update", "PUT", "/v1/people/{id}", UpdatePerson},
	Route{"Delete", "DELETE", "/v1/people/{id}", DeletePerson},
}
