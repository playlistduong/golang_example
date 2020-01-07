package server

import (
	"net/http"

	"github.com/golang_tutorial/internal/private_pkg/log"
)

func ListenAndServe(h http.Handler) {
	log.Info("Server is running port 8000")
	srv := &http.Server{
		Handler: h,
		Addr:    "127.0.0.1:8000",
	}

	srv.ListenAndServe()
}
