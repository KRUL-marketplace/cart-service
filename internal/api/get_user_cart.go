package api

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/converter"
	desc "github.com/KRUL-marketplace/cart-service/pkg/cart-service"
)

func (i *Implementation) GetUserCart(ctx context.Context, req *desc.GetUserCartRequest) (*desc.GetUserCartResponse, error) {
	cart, err := i.cartService.GetUserCart(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetUserCartResponse{
		Cart: converter.ToCartDescFromService(cart),
	}, nil
}
