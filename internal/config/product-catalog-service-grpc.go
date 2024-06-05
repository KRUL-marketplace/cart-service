package config

import (
	"fmt"
	"net"
	"os"
)

const (
	productCatalogServiceHost = "PRODUCT_CATALOG_SERVICE_GRPC_HOST"
	productCatalogServicePort = "PRODUCT_CATALOG_SERVICE_GRPC_PORT"
)

type ProductCatalogServiceGRPCConfig interface {
	Address() string
}

type productCatalogServiceGRPCConfig struct {
	host string
	port string
}

func NewProductCatalogServiceGRPCConfig() (ProductCatalogServiceGRPCConfig, error) {
	host := os.Getenv(productCatalogServiceHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", productCatalogServiceHost)
	}

	port := os.Getenv(productCatalogServicePort)
	if len(port) == 0 {
		return nil, fmt.Errorf("environment variable %s must be set", productCatalogServicePort)
	}

	return &productCatalogServiceGRPCConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *productCatalogServiceGRPCConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
