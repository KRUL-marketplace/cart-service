package repository

import (
	"cart-service/client/db"
	"cart-service/internal/connector/product_catalog_service"
	"cart-service/internal/repository/model"
	"context"
)

const (
	tableName       = "carts"
	cartIdColumn    = "cart_id"
	userIdColumn    = "user_id"
	itemIdColumn    = "item_id"
	cartIdFkColumn  = "cart_id"
	productIdColumn = "product_id"
	quantityColumn  = "quantity"
	nameColumn      = "name"
	slugColumn      = "slug"
	imageColumn     = "image"
	priceColumn     = "price"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
	statusColumn    = "status"
)

type Repository interface {
	Add(ctx context.Context, req *model.AddProductRequest) (string, error)
	GetUserCart(ctx context.Context, userId string) (*model.Cart, error)
	Delete(ctx context.Context, req *model.DeleteProductRequest) (string, error)
}

type repo struct {
	db                          db.Client
	productCatalogServiceClient product_catalog_service.ProductCatalogServiceClient
}

func NewRepository(db db.Client, productCatalogServiceClient product_catalog_service.ProductCatalogServiceClient) Repository {
	return &repo{
		db:                          db,
		productCatalogServiceClient: productCatalogServiceClient,
	}
}
