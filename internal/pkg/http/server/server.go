package server

import (
	"net/http"

	"github.com/golang_example/internal/pkg/log"
)

/**
* Create serve http
 */
func ListenAndServe(h http.Handler) {
	log.Info("Server is running port 8000")
	srv := &http.Server{
		Handler: h,
		Addr:    "127.0.0.1:8000",
	}

	srv.ListenAndServe()
}
