package http

import (
	"github.com/go-chi/chi"
)

func (transport *HTTP) initRoutes(r chi.Router) {
	r.Get("/ping", transport.endpoints.Ping)
	r.Get("/uptime", transport.endpoints.Uptime)

	r.NotFound(transport.endpoints.NotFound)
}
