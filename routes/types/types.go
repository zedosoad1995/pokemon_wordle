package route_types

import "net/http"

type RouteHandler func(w http.ResponseWriter, r *http.Request) error

type ErrorRes struct {
	Message string
}
