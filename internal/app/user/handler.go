package user

import (
	"net/http"

	"github.com/golang_tutorial/internal/private_pkg/http/respond"
)

type (
	service interface {
		GetAllUser()
		UpdateUser()
		CreateUser()
	}
	Handler struct {
	}
)

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"success": true,
		"page":    "@@@ users",
	}
	respond.JSON(w, 200, res)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"success": true,
		"page":    "@@@ users",
	}
	respond.JSON(w, 200, res)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	res := map[string]interface{}{
		"success": true,
		"page":    "@@@ users",
	}
	respond.JSON(w, 200, res)
}
