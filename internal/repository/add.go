package repository

import (
	"cart-service/client/db"
	"cart-service/internal/repository/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"log"
)

func (r *repo) Add(ctx context.Context, req *model.AddProductRequest) (string, error) {
	builderSelectCartListID := sq.Select("cart_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"user_id": req.UserID}).
		From(tableName)

	query, args, err := builderSelectCartListID.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "cart_repository.FindCart",
		QueryRaw: query,
	}

	var cartID string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&cartID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Создаём новую корзину, если не существует
			log.Printf("[Add] test 4")
			cartID, err = r.createCart(ctx, req.UserID)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}

	_, err = r.productCatalogServiceClient.GetById(ctx, []string{req.ProductID})
	if err != nil {
		return "", err
	}

	builderInsert := sq.Insert("cart_items").
		PlaceholderFormat(sq.Dollar).
		Columns("cart_id", "product_id", "quantity").
		Values(cartID, req.ProductID, req.Quantity).
		Suffix("ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity, updated_at = CURRENT_TIMESTAMP RETURNING item_id")

	query, args, err = builderInsert.ToSql()
	if err != nil {
		return "", err
	}

	var itemID string
	q = db.Query{
		Name:     "cart_repository.Add",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&itemID)
	if err != nil {
		return "", err
	}

	return itemID, nil
}

//		// Запрос для добавления или обновления товара в корзине
//		builderInsert := sq.Insert("cart_items").
//			PlaceholderFormat(sq.Dollar).
//			Columns("cart_id", "product_id", "quantity", "name", "slug", "image", "price").
//			Values(cartID, cartProductInfo.ProductId, cartProductInfo.Quantity, cartProductInfo.Name, cartProductInfo.Slug, cartProductInfo.Image, cartProductInfo.Price).
//			Suffix("ON CONFLICT (cart_id, product_id) DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity, updated_at = CURRENT_TIMESTAMP RETURNING item_id")
//
//		query, args, err = builderInsert.ToSql()
//		if err != nil {
//			return "", err
//		}
//
//		var itemID string
//		q = db.Query{
//			Name:     "cart_repository.Add",
//			QueryRaw: query,
//		}
//
//		err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&itemID)
//		if err != nil {
//			return "", err
//		}
//
//		return itemID, nil
//	}
func (r *repo) createCart(ctx context.Context, userId string) (string, error) {
	log.Printf("[Add] createCart 1")

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(userIdColumn).
		Values(userId).
		Suffix("RETURNING  cart_id")

	log.Printf("[Add] createCart 2")
	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	log.Printf("[Add] createCart 3")
	var cartID string
	q := db.Query{
		Name:     "cart_repository.createCart",
		QueryRaw: query,
	}

	log.Printf("[Add] createCart 4")
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&cartID)
	if err != nil {
		return "", err
	}
	log.Printf("[Add] createCart 5")

	return cartID, nil
}
