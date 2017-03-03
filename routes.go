package main

import (
	"net/http"

	"github.com/deep/GoMVC/controller/mafiaAPI"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		mafiaAPI.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		mafiaAPI.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		mafiaAPI.TodoShow,
	},
}
