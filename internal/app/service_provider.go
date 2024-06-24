package app

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/api"
	"github.com/KRUL-marketplace/cart-service/internal/config"
	"github.com/KRUL-marketplace/cart-service/internal/connector/product_catalog_service"
	"github.com/KRUL-marketplace/cart-service/internal/repository"
	"github.com/KRUL-marketplace/cart-service/internal/service"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db/pg"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db/transaction"
	"github.com/KRUL-marketplace/common-libs/pkg/closer"
	product_service "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type serviceProvider struct {
	cartRepository repository.Repository
	cartService    service.CartService

	grpcConfig                      config.GRPCConfig
	httpConfig                      config.HTTPConfig
	pgConfig                        config.PGConfig
	swaggerConfig                   config.SwaggerConfig
	productCatalogServiceGRPCConfig config.ProductCatalogServiceGRPCConfig
	redisConfig                     config.RedisConfig

	dbClient  db.Client
	txManager db.TxManager

	cartImpl *api.Implementation

	productCatalogServiceClient product_catalog_service.ProductCatalogServiceClient

	redisClient *redis.Client
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) ProductCatalogServiceGRPCConfig() config.ProductCatalogServiceGRPCConfig {
	if s.productCatalogServiceGRPCConfig == nil {
		cfg, err := config.NewProductCatalogServiceGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get product catalog service grpc config: %s", err.Error())
		}

		s.productCatalogServiceGRPCConfig = cfg
	}

	return s.productCatalogServiceGRPCConfig
}

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) CartRepository(ctx context.Context) repository.Repository {
	if s.cartRepository == nil {
		s.cartRepository = repository.NewRepository(
			s.DBClient(ctx),
			s.RedisClient(ctx),
			s.ProductCatalogServiceClient(ctx),
		)
	}

	return s.cartRepository
}

func (s *serviceProvider) CartService(ctx context.Context) service.CartService {
	if s.cartService == nil {
		s.cartService = service.NewService(
			s.CartRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.cartService
}

func (s *serviceProvider) CartImpl(ctx context.Context) *api.Implementation {
	if s.cartImpl == nil {
		s.cartImpl = api.NewImplementation(s.CartService(ctx))
	}

	return s.cartImpl
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) RedisClient(ctx context.Context) redis.Client {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(&redis.Options{
			Addr:     s.RedisConfig().Address(),
			Password: "",
			DB:       0,
		})
	}

	return *s.redisClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ProductCatalogServiceClient(ctx context.Context) product_catalog_service.ProductCatalogServiceClient {
	if s.productCatalogServiceClient == nil {
		conn, err := grpc.DialContext(ctx,
			s.ProductCatalogServiceGRPCConfig().Address(), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("product catalog service client init error")
		}

		s.productCatalogServiceClient = product_catalog_service.NewProductCatalogServiceClient(
			product_service.NewProductCatalogServiceClient(conn),
		)
	}

	return s.productCatalogServiceClient
}
