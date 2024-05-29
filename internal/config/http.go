package config

import (
	"fmt"
	"net"
	"os"
)

const (
	httpHostEnvName = "HTTP_HOST"
	httpPortEnvName = "HTTP_PORT"
)

type HTTPConfig interface {
	Address() string
}

type httpConfig struct {
	host string
	port string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", httpHostEnvName)
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", httpPortEnvName)
	}

	return &httpConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *httpConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
