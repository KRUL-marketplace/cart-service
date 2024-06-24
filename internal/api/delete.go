package api

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/converter"
	desc "github.com/KRUL-marketplace/cart-service/pkg/cart-service"
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
