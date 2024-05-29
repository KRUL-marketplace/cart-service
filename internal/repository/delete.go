package repository

import (
	"cart-service/client/db"
	"cart-service/internal/repository/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"log"
	"time"
)

func (r *repo) Delete(ctx context.Context, userId string, cartProductInfo *model.DeleteCartProductInfo) (string, error) {
	builder := sq.Select("cart_id").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{userIdColumn: userId}).
		From(tableName)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	var cartID string
	q := db.Query{
		Name:     "cart_repository.FindCart",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&cartID)
	if err != nil {
		log.Printf("[DeleteProduct] error finding cart for user %s: %v", userId, err)
		return "", err
	}

	var currentQuantity uint32
	builderSelect := sq.Select("quantity").
		PlaceholderFormat(sq.Dollar).
		From("cart_items").
		Where(sq.Eq{"cart_id": cartID, "product_id": cartProductInfo.ProductId})

	query, args, err = builderSelect.ToSql()
	if err != nil {
		return "", err
	}

	q = db.Query{
		Name:     "cart_repository.GetCurrentQuantity",
		QueryRaw: query,
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&currentQuantity)
	if err != nil {
		log.Printf("[DeleteProduct] error finding product %s in cart %s: %v", cartProductInfo.ProductId, cartID, err)
		return "", err
	}

	if currentQuantity > cartProductInfo.Quantity {
		newQuantity := currentQuantity - cartProductInfo.Quantity
		builderUpdate := sq.Update("cart_items").
			PlaceholderFormat(sq.Dollar).
			Set("quantity", newQuantity).
			Set("updated_at", time.Now()).
			Where(sq.Eq{"cart_id": cartID, "product_id": cartProductInfo.ProductId})

		query, args, err = builderUpdate.ToSql()
		if err != nil {
			return "", err
		}

		q = db.Query{
			Name:     "cart_repository.UpdateQuantity",
			QueryRaw: query,
		}

		_, err = r.db.DB().ExecContext(ctx, q, args...)
		if err != nil {
			return "", err
		}
	} else {
		// Delete the product from the cart
		builderDelete := sq.Delete("cart_items").
			PlaceholderFormat(sq.Dollar).
			Where(sq.Eq{"cart_id": cartID, "product_id": cartProductInfo.ProductId})

		query, args, err = builderDelete.ToSql()
		if err != nil {
			return "", err
		}

		q = db.Query{
			Name:     "cart_repository.DeleteProduct",
			QueryRaw: query,
		}

		_, err = r.db.DB().ExecContext(ctx, q, args...)
		if err != nil {
			return "", err
		}
	}

	return "", nil
}
