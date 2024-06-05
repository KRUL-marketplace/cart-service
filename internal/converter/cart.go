package converter

import (
	"cart-service/internal/repository/model"
	desc "cart-service/pkg/cart-service"
	"database/sql"
	product_service "github.com/KRUL-marketplace/product-catalog-service/pkg/product-catalog-service"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func ToAddProductRequestFromDesc(req *desc.AddProductRequest) *model.AddProductRequest {
	return &model.AddProductRequest{
		UserID:    req.GetUserId(),
		ProductID: req.GetProductId(),
		Quantity:  req.GetQuantity(),
	}
}

func ToDeleteProductRequestFromDesc(req *desc.DeleteProductRequest) *model.DeleteProductRequest {
	return &model.DeleteProductRequest{
		UserID:    req.GetUserId(),
		ProductID: req.GetProductId(),
		Quantity:  req.GetQuantity(),
	}
}

func ToCartProductInfoModelFromDesc(req *product_service.ProductInfo) *model.CartProductInfo {
	return &model.CartProductInfo{
		Name:  req.GetName(),
		Slug:  req.GetSlug(),
		Image: "test",
		Price: req.GetPrice(),
		Brand: model.Brand{
			ID: req.GetBrand().GetId(),
			Info: model.BrandInfo{
				Name:        req.GetBrand().GetInfo().GetName(),
				Slug:        req.GetBrand().GetInfo().GetSlug(),
				Description: req.GetBrand().GetInfo().GetDescription(),
			},
			CreatedAt: req.GetBrand().GetCreatedAt().AsTime(),
			UpdatedAt: sql.NullTime{
				Time:  req.GetBrand().GetUpdatedAt().AsTime(),
				Valid: req.GetBrand().GetUpdatedAt().IsValid(),
			},
		},
	}
}

func ToCartDescFromService(cart *model.Cart) *desc.Cart {
	var updatedAt time.Time
	if cart.UpdatedAt.Valid {
		updatedAt = cart.UpdatedAt.Time
	}

	result := &desc.Cart{
		CartId:     cart.CartID,
		UserId:     cart.UserID,
		TotalPrice: cart.TotalPrice,
		CreatedAt:  timestamppb.New(cart.CreatedAt),
		UpdatedAt:  timestamppb.New(updatedAt),
	}

	for _, product := range cart.Products {
		result.Products = append(result.Products, ToCartProductDescFromService(&product))
	}

	return result
}

func ToCartProductDescFromService(product *model.CartProduct) *desc.CartProduct {
	var infoUpdatedAt time.Time
	if product.UpdatedAt.Valid {
		infoUpdatedAt = product.UpdatedAt.Time
	}

	return &desc.CartProduct{
		ItemId:    product.ItemID,
		ProductId: product.ProductID,
		Info:      ToCartProductInfoDescFromService(&product.Info),
		Quantity:  product.Quantity,
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: timestamppb.New(infoUpdatedAt),
	}
}

func ToCartProductInfoDescFromService(productInfo *model.CartProductInfo) *desc.CartProductInfo {
	var brandUpdatedAt time.Time
	if productInfo.Brand.UpdatedAt.Valid {
		brandUpdatedAt = productInfo.Brand.UpdatedAt.Time
	}

	return &desc.CartProductInfo{
		Name:  productInfo.Name,
		Slug:  productInfo.Slug,
		Image: productInfo.Image,
		Price: productInfo.Price,
		Brand: &desc.Brand{
			Id: productInfo.Brand.ID,
			Info: &desc.BrandInfo{
				Name:        productInfo.Brand.Info.Name,
				Slug:        productInfo.Brand.Info.Slug,
				Description: productInfo.Brand.Info.Description,
			},
			CreatedAt: timestamppb.New(productInfo.Brand.CreatedAt),
			UpdatedAt: timestamppb.New(brandUpdatedAt),
		},
	}
}
