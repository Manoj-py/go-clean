package routes

import (
	handlers "github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/handlers"
	"net/http"

	httprouters "github.com/Manoj-py/backend-architecture/common/httprouter"
)

func Routes(router httprouters.RouterI, handler handlers.HandlerI) http.Handler {
	mux := router.NewChiDefaultRouter()
	mux.Post("/", handler.Createuser)
	return mux

}
