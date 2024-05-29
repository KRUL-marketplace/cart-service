package config

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcHostEnvName)
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", grpcPortEnvName)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
