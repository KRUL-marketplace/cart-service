package api

import (
	"cart-service/internal/converter"
	desc "cart-service/pkg/cart-service"
	"context"
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
