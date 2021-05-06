package main

import (
	"time"

	"go.uber.org/fx"
	"skeleton-go/domain/usecases"
	"skeleton-go/infrastructure/fxapp"
	"skeleton-go/infrastructure/logger"
	http2 "skeleton-go/infrastructure/transports/http"
)

const (
	startTimeout = 30 * time.Second
	stopTimeout  = startTimeout
)

func main() {
	app := fx.New(
		fx.NopLogger, // remove for debug

		// Here, all the implementations you want to use for your app
		logger.Development,
		usecases.Admin,
		http2.Transport,

		// Here, the functions that are executed eagerly on application start
		fx.Invoke(logger.Flush),
		fx.Invoke(http2.Run),
	)

	fxapp.Start(app, startTimeout)
	<-app.Done()
	fxapp.Shutdown(app, stopTimeout)
}
