package handlers

import (
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/service"
	"net/http"
)

type HandlerI interface {
	Createuser(w http.ResponseWriter, r *http.Request)
}
type Handler struct {
	store db.Store
	svc   service.ServiceI
}

func NewHandler(store db.Store, svc service.ServiceI) HandlerI {
	return &Handler{
		store: store,
		svc:   svc,
	}
}
