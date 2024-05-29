package repository

import (
	"cart-service/client/db"
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
	Add(ctx context.Context, userId string, cartProductInfo *model.CartProductInfo) (string, error)
	GetUserCart(ctx context.Context, userId string) (*model.Cart, error)
	Delete(ctx context.Context, userId string, cartProductInfo *model.DeleteCartProductInfo) (string, error)
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repository {
	return &repo{
		db: db,
	}
}
