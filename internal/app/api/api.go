package api

import (
	"github.com/golang_tutorial/internal/app/user"
	"github.com/golang_tutorial/internal/private_pkg/http/router"
)

func CreateRouter() {
	userHandler := user.NewHandler()
	usersRoutes := userHandler.Routes()
	var routes []router.Route
	routes = append(routes, usersRoutes...)

	router.CreateRouter(routes)
}
