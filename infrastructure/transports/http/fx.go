package http

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"skeleton-go/infrastructure/fxapp"
)

// Transport contains all the dependencies used to run HTTP transport layer.
var Transport = fx.Provide(
	func() *Config {
		return &Config{Scheme: "http", Host: "localhost", Port: 8080}
	},

	fx.Annotated{Name: "http", Target: NewHTTP},
)

// FxParams is the parameter used by uber-go/fx for the dependency injection.
type FxParams struct {
	fx.In
	Lifecycle  fx.Lifecycle
	Shutdowner fx.Shutdowner
	Logger     *zap.Logger
	Transport  fxapp.TCP `name:"http"`
}

// Run registers the HTTP transport.
func Run(p FxParams) {
	fxServer := fxapp.NewServer(p.Lifecycle, p.Shutdowner, p.Logger)
	fxServer.Run("http", p.Transport)
}
