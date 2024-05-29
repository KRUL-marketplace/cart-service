package api

import (
	"cart-service/internal/service"
	desc "cart-service/pkg/cart-service"
)

type Implementation struct {
	desc.UnimplementedCartServiceServer
	cartService service.CartService
}

func NewImplementation(cartService service.CartService) *Implementation {
	return &Implementation{
		cartService: cartService,
	}
}
