package app

import (
	"cart-service/client/db"
	"cart-service/client/db/pg"
	"cart-service/client/db/transaction"
	"cart-service/internal/api"
	"cart-service/internal/config"
	"cart-service/internal/repository"
	"cart-service/internal/service"
	"context"
	"log"
)

type serviceProvider struct {
	cartRepository repository.Repository
	cartService    service.CartService

	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	pgConfig      config.PGConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager

	cartImpl *api.Implementation
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

func (s *serviceProvider) CartRepository(ctx context.Context) repository.Repository {
	if s.cartRepository == nil {
		s.cartRepository = repository.NewRepository(s.DBClient(ctx))
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

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}
