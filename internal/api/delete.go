package api

import (
	"cart-service/internal/converter"
	desc "cart-service/pkg/cart-service"
	"context"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteProductRequest) (*desc.DeleteProductResponse, error) {
	_, err := i.cartService.Delete(ctx, converter.ToDeleteProductRequestFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.DeleteProductResponse{
		Message: "Success",
	}, nil
}
