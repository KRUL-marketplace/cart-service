package service

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/model"
)

func (s *cartService) GetUserCart(ctx context.Context, userId string) (*model.Cart, error) {
	cart, err := s.cartRepository.GetUserCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
