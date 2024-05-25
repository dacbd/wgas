package server

import (
	"github.com/dacbd/gaa/routes"
	"net/http"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	routes.AddRoutes(mux)
	var handler *http.ServeMux = mux

	return handler
}
