package user

import (
	"net/http"

	"github.com/golang_tutorial/internal/private_pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/users",
			Method:  http.MethodGet,
			Handler: h.GetAllUser,
		},
		{
			Path:    "/api/v1/users",
			Method:  http.MethodPut,
			Handler: h.UpdateUser,
		},
		{
			Path:    "/api/v1/users",
			Method:  http.MethodPost,
			Handler: h.CreateUser,
		},
	}
}
