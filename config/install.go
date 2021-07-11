package config

import (
	"github.com/palantir/witchcraft-go-server/v2/config"
)

type Install struct {
	config.Install `yaml:",inline"`
}
