package router

import (
	"net/http"

	"github.com/golang_tutorial/internal/private_pkg/http/server"
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

func CreateRouter(routes []Route) {
	// create mux method router
	r := mux.NewRouter()
	for _, route := range routes {
		// add middleware
		for _, middleware := range route.Middlewares {
			r.Use(middleware)
		}
		r.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
	server.ListenAndServe(r)
}
