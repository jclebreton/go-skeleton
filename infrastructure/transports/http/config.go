// nolint:unused
// All definition in this package expect for the config object are used to
// ensure config validation. Validation process is made through reflexion so
// unused linter does not see it.
package http

import (
	"fmt"
)

type Config struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

func (config Config) GetAddress() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

func (config Config) GetURL() string {
	return fmt.Sprintf("%s://%s", config.Scheme, config.GetAddress())
}
