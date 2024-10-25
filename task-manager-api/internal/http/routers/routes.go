package routers

import (
	handlers "github.com/Manoj-py/backend-architecture/task-manager-api/internal/http/handlers"
	"net/http"

	httprouters "github.com/Manoj-py/backend-architecture/common/httprouter"
)

func Routes(mux *httprouters.ChiRouter, handler *handlers.Handler) http.Handler {
	r := mux.NewDefaultRouter()
	r.Get("/", handler.Createuser)
	return r

}
