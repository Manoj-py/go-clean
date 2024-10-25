package handlers

import (
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
)

type Handler struct {
	store db.Store
}

func NewHandler(store db.Store) *Handler {
	return &Handler{
		store: store,
	}
}
