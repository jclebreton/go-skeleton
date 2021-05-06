package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"skeleton-go/domain/usecases"
	"skeleton-go/infrastructure/fxapp"
	"skeleton-go/infrastructure/transports/http/endpoints"
)

type HTTP struct {
	httpConfig *Config
	endpoints  *endpoints.Endpoints
	httpServer *http.Server
}

func NewHTTP(httpConfig *Config, logger *zap.Logger, uc *usecases.Usecases) fxapp.TCP {
	return &HTTP{
		httpConfig: httpConfig,
		endpoints:  endpoints.NewEndpoints(logger, uc),
	}
}

func (transport *HTTP) ListenAndServe() error {
	r := chi.NewRouter()
	transport.initMiddlewares(r)
	transport.initRoutes(r)

	transport.httpServer = &http.Server{
		Addr:    transport.httpConfig.GetAddress(),
		Handler: r,
	}

	return transport.httpServer.ListenAndServe()
}

func (transport *HTTP) Shutdown() error {
	return transport.httpServer.Shutdown(context.Background())
}

func (transport *HTTP) GetAddress() string {
	return transport.httpConfig.GetAddress()
}

func (transport *HTTP) initMiddlewares(router chi.Router) {

}
