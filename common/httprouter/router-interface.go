package httprouters

import "net/http"

type Router interface {
	NewDefaultRouter() http.Handler
}
