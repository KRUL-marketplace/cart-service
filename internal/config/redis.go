package config

import (
	"fmt"
	"net"
	"os"
)

const (
	redisHostEnvName = "REDIS_HOST"
	redisPortEnvName = "REDIS_PORT"
)

type RedisConfig interface {
	Address() string
}

type redisConfig struct {
	host string
	port string
}

func NewRedisConfig() (RedisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", redisHostEnvName)
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", redisPortEnvName)
	}

	return &redisConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
