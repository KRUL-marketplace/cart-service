package service

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/repository"
	"github.com/KRUL-marketplace/cart-service/model"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
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
