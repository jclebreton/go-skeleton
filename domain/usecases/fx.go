package usecases

import (
	"go.uber.org/fx"
)

// Admin contains all the "admin" usecases
var Admin = fx.Provide(NewUsecases)
