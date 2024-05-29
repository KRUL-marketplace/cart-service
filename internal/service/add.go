package service

import (
	"cart-service/internal/repository/model"
	"context"
)

func (s *cartService) Add(ctx context.Context, userId string, cartProductInfo *model.CartProductInfo) (string, error) {
	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.cartRepository.Add(ctx, userId, cartProductInfo)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}
