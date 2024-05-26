package server

import (
	"github.com/dacbd/wgas/routes"
	"net/http"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()

	routes.AddRoutes(mux)
	var handler *http.ServeMux = mux

	return handler
}
