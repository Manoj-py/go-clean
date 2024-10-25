package handlers

import (
	"encoding/json"
	db "github.com/Manoj-py/backend-architecture/task-manager-api/internal/db/sqlc"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/models"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func (h *Handler) Createuser(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequestModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		slog.Error("1 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"success": "fail", "response": err.Error()})
		return
	}
	userId, _ := uuid.NewUUID()

	res, err := h.store.CreateUser(r.Context(), db.CreateUserParams{
		Userid:   userId,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     db.NullUserRole{},
	})
	if err != nil {
		slog.Error("1 ", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]string{"success": "fail", "response": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	slog.Info("User created successfully", res)
	_ = json.NewEncoder(w).Encode(map[string]any{"success": "true", "response": res})
}
