package config

import (
	"github.com/palantir/witchcraft-go-server/v2/config"
)

type Runtime struct {
	config.Runtime `yaml:",inline"`
}
