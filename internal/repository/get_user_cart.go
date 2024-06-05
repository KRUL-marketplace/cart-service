package repository

import (
	"cart-service/client/db"
	"cart-service/internal/converter"
	"cart-service/internal/repository/model"
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"time"
)

func (r *repo) GetUserCart(ctx context.Context, userId string) (*model.Cart, error) {
	builder := sq.Select(
		"c.cart_id",
		"c.user_id",
		"c.created_at",
		"c.updated_at",
		"ci.item_id",
		"ci.product_id",
		"ci.quantity",
		"ci.created_at",
		"ci.updated_at",
	).
		From("carts as c").
		LeftJoin("cart_items ci ON c.cart_id = ci.cart_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"c.user_id": userId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "cart_repository.GetByUserId " + userId,
		QueryRaw: query,
	}

	var cart model.Cart
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cart.Products = []model.CartProduct{}
	var productIDS []string

	for rows.Next() {
		var cartProduct model.CartProduct
		var itemID, productID sql.NullString
		var quantity sql.NullInt32
		var createdAt, updatedAt sql.NullTime

		err := rows.Scan(
			&cart.CartID,
			&cart.UserID,
			&cart.CreatedAt,
			&cart.UpdatedAt,
			&itemID,
			&productID,
			&quantity,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		if itemID.Valid {
			cartProduct.ItemID = itemID.String
		} else {
			cartProduct.ItemID = ""
		}

		if productID.Valid {
			cartProduct.ProductID = productID.String
		} else {
			cartProduct.ProductID = ""
		}

		if quantity.Valid {
			cartProduct.Quantity = uint32(quantity.Int32)
		} else {
			cartProduct.Quantity = 0
		}

		if createdAt.Valid {
			cartProduct.CreatedAt = createdAt.Time
		} else {
			cartProduct.CreatedAt = time.Time{}
		}

		if updatedAt.Valid {
			cartProduct.UpdatedAt = updatedAt
		} else {
			cartProduct.UpdatedAt = sql.NullTime{}
		}

		if productID.Valid {
			productIDS = append(productIDS, productID.String)
			cart.Products = append(cart.Products, cartProduct)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	result, err := r.productCatalogServiceClient.GetById(ctx, productIDS)
	if err != nil {
		return nil, err
	}

	for i, product := range result.GetProduct() {
		cart.Products[i].Info = *converter.ToCartProductInfoModelFromDesc(product.GetInfo())
		cart.TotalPrice += cart.Products[i].Quantity * cart.Products[i].Info.Price
	}

	return &cart, nil
}

//	q := db.Query{
//		Name:     "cart_repositry.GetByUserId " + userId,
//		QueryRaw: query,
//	}
//
//	var cart model.Cart
//	rows, err := r.db.DB().QueryContext(ctx, q, args...)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	cart.Products = []model.CartProduct{}
//
//	var totalPrice uint32
//	for rows.Next() {
//		var product model.CartProduct
//		var productInfo model.CartProductInfo
//
//		err := rows.Scan(
//			&cart.CartID,
//			&cart.UserID,
//			&cart.CreatedAt,
//			&cart.UpdatedAt,
//			&product.ID,
//			&productInfo.ProductId,
//			&productInfo.Quantity,
//			&productInfo.Name,
//			&productInfo.Slug,
//			&productInfo.Image,
//			&productInfo.Price,
//			&product.CreatedAt,
//			&product.UpdatedAt,
//		)
//
//		if err != nil {
//			return nil, err
//		}
//
//		product.Info = productInfo
//		cart.Products = append(cart.Products, product)
//
//		totalPrice += product.Info.Price * product.Info.Quantity
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	cart.TotalPrice = totalPrice
//
//	return &cart, nil
//}
