package repository

import (
	"cart-service/client/db"
	"cart-service/internal/repository/model"
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (r *repo) GetUserCart(ctx context.Context, userId string) (*model.Cart, error) {
	builder := sq.Select(
		"c.cart_id",
		"c.user_id",
		"c.created_at AS cart_created_at",
		"c.updated_at AS cart_updated_at",
		"ci.item_id",
		"ci.product_id",
		"ci.quantity",
		"ci.name AS product_name",
		"ci.slug AS product_slug",
		"ci.image AS product_image",
		"ci.price AS product_price",
		"ci.created_at AS item_created_at",
		"ci.updated_at AS item_updated_at",
	).
		From("carts c").
		LeftJoin("cart_items ci ON c.cart_id = ci.cart_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"c.user_id": userId})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "cart_repositry.GetByUserId " + userId,
		QueryRaw: query,
	}

	var cart model.Cart
	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cart.Products = []model.CartProduct{}

	var totalPrice uint32
	for rows.Next() {
		var product model.CartProduct
		var productInfo model.CartProductInfo

		err := rows.Scan(
			&cart.CartID,
			&cart.UserID,
			&cart.CreatedAt,
			&cart.UpdatedAt,
			&product.ID,
			&productInfo.ProductId,
			&productInfo.Quantity,
			&productInfo.Name,
			&productInfo.Slug,
			&productInfo.Image,
			&productInfo.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		product.Info = productInfo
		cart.Products = append(cart.Products, product)

		totalPrice += product.Info.Price * product.Info.Quantity
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	cart.TotalPrice = totalPrice

	return &cart, nil
}
