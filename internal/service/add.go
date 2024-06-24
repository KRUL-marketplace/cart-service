package service

import (
	"context"
	"github.com/KRUL-marketplace/cart-service/internal/repository/model"
)

func (s *cartService) Add(ctx context.Context, req *model.AddProductRequest) (string, error) {
	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.cartRepository.Add(ctx, req)
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
