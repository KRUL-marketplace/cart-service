package service

import (
	"cart-service/internal/repository/model"
	"context"
)

func (s *cartService) Delete(ctx context.Context, req *model.DeleteProductRequest) (string, error) {
	msg, err := s.cartRepository.Delete(ctx, req)
	if err != nil {
		return msg, err
	}

	return msg, nil
}
