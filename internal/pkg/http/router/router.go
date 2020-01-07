package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

type (
	Middleware = func(http.Handler) http.Handler
	Route      struct {
		Path        string
		Method      string
		Handler     http.HandlerFunc
		Middlewares []Middleware
	}
)

func CreateRouter(routes []Route) http.Handler {
	// create mux method router
	r := mux.NewRouter()
	for _, route := range routes {
		// add middleware
		for _, middleware := range route.Middlewares {
			r.Use(middleware)
		}
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}

	return r
}
