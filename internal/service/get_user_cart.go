package service

import (
	"cart-service/internal/repository/model"
	"context"
)

func (s *cartService) GetUserCart(ctx context.Context, userId string) (*model.Cart, error) {
	cart, err := s.cartRepository.GetUserCart(ctx, userId)
	if err != nil {
		return nil, err
	}

	return cart, nil
}
