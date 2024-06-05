package api

import (
	"cart-service/internal/converter"
	desc "cart-service/pkg/cart-service"
	"context"
)

func (i *Implementation) Add(ctx context.Context, req *desc.AddProductRequest) (*desc.AddProductResponse, error) {
	id, err := i.cartService.Add(ctx, converter.ToAddProductRequestFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.AddProductResponse{
		Id: id,
	}, nil
}
