package api

import (
	"github.com/KRUL-marketplace/cart-service/internal/service"
	desc "github.com/KRUL-marketplace/cart-service/pkg/cart-service"
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
