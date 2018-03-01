package main

import "net/http"

//Route object creates to keep track of routes for router
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes is an array of Route objects
type Routes []Route

var routes = Routes{
	//Home page
	///Authenticated
	Route{
		"Index",
		"GET",
		"/mcfly/main",
		indexHandler,
	},
	Route{
		"Upload",
		"POST",
		"/mcfly/upload",
		uploadHandler,
	},
}
