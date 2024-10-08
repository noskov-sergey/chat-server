package config

import (
	"errors"
	"net"
	"os"
)

const (
	swaggerHostEnvName = "SWAGGER_HOST"
	swaggerPortEnvName = "SWAGGER_PORT"
)

type SwaggerConfig interface {
	Address() string
}

type swaggerConfig struct {
	host string
	port string
}

func NewSWAGGERConfig() (*swaggerConfig, error) {
	host := os.Getenv(swaggerHostEnvName)

	port := os.Getenv(swaggerPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("swagger port not found")
	}

	return &swaggerConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *swaggerConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
