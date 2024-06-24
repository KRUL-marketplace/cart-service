package api

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/converter"
	desc "github.com/KRUL-marketplace/cart-service/pkg/cart-service"
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
