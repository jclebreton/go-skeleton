package usecases

import (
	"time"
)

var startedAt = time.Now()

func (uc Usecases) GetAppUptime() time.Duration {
	return time.Now().Sub(startedAt)
}
