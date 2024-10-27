package handlers

import (
	"github.com/Manoj-py/backend-architecture/common/httpUtils"
	"github.com/Manoj-py/backend-architecture/task-manager-api/internal/models"
	"net/http"
)

func (h *Handler) Createuser(w http.ResponseWriter, r *http.Request) {
	var user models.UserRequestModel
	if err := httpUtils.ReadJsonAndValidate(w, r, &user); err != nil {
		httpUtils.ErrorJson(w, err)
		return
	}
	svcResponse, err := h.svc.CreateUser(r.Context(), user)
	if err != nil {
		httpUtils.ErrorJson(w, err)
		return
	}
	httpUtils.WriteJson(w, http.StatusOK, svcResponse)
}
