package service

import (
	"cart-service/internal/repository/model"
	"context"
)

func (s *cartService) Delete(ctx context.Context, userId string, cartProductInfo *model.DeleteCartProductInfo) (string, error) {
	msg, err := s.cartRepository.Delete(ctx, userId, cartProductInfo)
	if err != nil {
		return msg, err
	}

	return msg, nil
}
