package converter

import (
	"cart-service/internal/repository/model"
	desc "cart-service/pkg/cart-service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCartProductInfoFromDesc(info *desc.CartProductInfo) *model.CartProductInfo {
	cartProductInfo := &model.CartProductInfo{
		ProductId: info.ProductID,
		Name:      info.Name,
		Slug:      info.Slug,
		Image:     info.Image,
		Price:     info.Price,
		Quantity:  info.Quantity,
	}

	return cartProductInfo
}

func ToDeleteCartProductInfofromDesc(info *desc.DeleteProductInfo) *model.DeleteCartProductInfo {
	return &model.DeleteCartProductInfo{
		ProductId: info.ProductID,
		Quantity:  info.Quantity,
	}
}

func ToCartFromService(info *model.Cart) *desc.Cart {
	cart := &desc.Cart{
		CartId:     info.CartID,
		UserId:     info.UserID,
		CreatedAt:  timestamppb.New(info.CreatedAt),
		TotalPrice: info.TotalPrice,
	}

	for _, cartProduct := range info.Products {
		cart.CartProduct = append(cart.CartProduct, &desc.CartProduct{
			Id:        cartProduct.ID,
			CreatedAt: timestamppb.New(cartProduct.CreatedAt),
			Info: &desc.CartProductInfo{
				ProductID: cartProduct.Info.ProductId,
				Name:      cartProduct.Info.Name,
				Slug:      cartProduct.Info.Slug,
				Image:     cartProduct.Info.Image,
				Price:     cartProduct.Info.Price,
				Quantity:  cartProduct.Info.Quantity,
			},
		})
	}

	return cart
}
