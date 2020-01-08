package api

import (
	"github.com/golang_example/internal/app/user"
	"github.com/golang_example/internal/pkg/http/router"
	"github.com/golang_example/internal/pkg/http/server"
)

func CreateRouter() {
	userHandler := user.NewHandler()
	usersRoutes := userHandler.Routes()
	var routes []router.Route
	routes = append(routes, usersRoutes...)

	httpHandler := router.CreateRouter(routes)

	server.ListenAndServe(httpHandler)

}
