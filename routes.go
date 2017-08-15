package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		GetPeopleEndpoint,
	},
	Route{
		"People",
		"GET",
		"/people",
		GetPeopleEndpoint,
	},
	Route{
		"ShowPerson",
		"GET",
		"/people/{id}",
		GetPersonEndpoint,
	},
	Route{
		"AddPerson",
		"POST",
		"/people/{id}",
		CreatePersonEndpoint,
	},
	Route{
		"DeletePerson",
		"DELETE",
		"/people/{id}",
		DeletePersonEndpoint,
	},
}
