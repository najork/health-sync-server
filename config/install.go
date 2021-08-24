package config

import (
	"github.com/palantir/witchcraft-go-server/v2/config"
)

type Install struct {
	config.Install `yaml:",inline"`
	Credentials    GarminCredentials `yaml:"credentials,omitempty"`
}

type GarminCredentials struct {
	Email    string `yaml:"email,omitempty"`
	Password string `yaml:"password,omitempty"`
}
