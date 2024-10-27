package httprouters

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (r *Router) NewChiDefaultRouter() *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Timeout(60 * time.Second))
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	return mux
}
