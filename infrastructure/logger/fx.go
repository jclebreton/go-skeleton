package logger

import (
	"context"
	"runtime"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

// /!\ Failures to sync stdout and stderr on MacOSX:
// Users can't do anything about failures to sync stderr and stdout.
// See https://github.com/uber-go/zap/issues/328 and https://github.com/uber-go/zap/issues/370
const syncBug = "sync /dev/stderr: inappropriate ioctl for device"

var (
	// Development provides the logger to use for dev environment
	Development = fx.Provide(func() (*zap.Logger, error) {
		logger, err := zap.NewDevelopment()
		if err != nil {
			return nil, err
		}

		zap.ReplaceGlobals(logger)

		return logger, nil
	})

	// Prod provides the logger to use for prod environment
	Prod = fx.Provide(func() (*zap.Logger, error) {
		logger, err := zap.NewProduction()
		if err != nil {
			return nil, err
		}

		zap.ReplaceGlobals(logger)

		return logger, nil
	})
)

// Flush allows to sync logger before shutdown
// /!\ This graceful function must be the last invoked to ensure logging
func Flush(lifecycle fx.Lifecycle, logger *zap.Logger) {
	logger = logger.Named("Lifecycle")

	lifecycle.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				if err := logger.Sync(); err != nil && !(runtime.GOOS == "darwin" && err.Error() == syncBug) {
					logger.Error("🧨💥 Unable to sync logger", zap.Error(err))
					return err
				}

				logger.Debug("📢 Logger flushed")

				return nil
			},
		},
	)
}
