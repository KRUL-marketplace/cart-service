package service

import (
	"cart-service/client/db"
	"cart-service/internal/repository"
	"cart-service/internal/repository/model"
	"context"
)

type cartService struct {
	cartRepository repository.Repository
	txManager      db.TxManager
}

type CartService interface {
	Add(ctx context.Context, userId string, cartProductInfo *model.CartProductInfo) (string, error)
	GetUserCart(ctx context.Context, userId string) (*model.Cart, error)
	Delete(ctx context.Context, userId string, cartProductInfo *model.DeleteCartProductInfo) (string, error)
}

func NewService(cartRepository repository.Repository, txManager db.TxManager) CartService {
	return &cartService{
		cartRepository: cartRepository,
		txManager:      txManager,
	}
}
