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
	Add(ctx context.Context, req *model.AddProductRequest) (string, error)
	GetUserCart(ctx context.Context, userId string) (*model.Cart, error)
	Delete(ctx context.Context, req *model.DeleteProductRequest) (string, error)
}

func NewService(cartRepository repository.Repository, txManager db.TxManager) CartService {
	return &cartService{
		cartRepository: cartRepository,
		txManager:      txManager,
	}
}
