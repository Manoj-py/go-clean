package httprouters

import (
	"github.com/go-chi/chi/v5"
)

type RouterI interface {
	NewChiDefaultRouter() *chi.Mux
}

type Router struct {
}

func NewRouter() RouterI {
	return &Router{}
}
