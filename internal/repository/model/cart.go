package model

import (
	"database/sql"
	"time"
)

type Cart struct {
	CartID     string       `db:"cart_id"`
	UserID     string       `db:"user_id"`
	CreatedAt  time.Time    `db:"created_at"`
	UpdatedAt  sql.NullTime `db:"updated_at"`
	TotalPrice uint32
	Products   []CartProduct
}

type CartProduct struct {
	ID        string          `db:"id"`
	Info      CartProductInfo `db:""`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"updated_at"`
}

type CartProductInfo struct {
	ProductId string `db:"product_id"`
	Name      string `db:"name"`
	Slug      string `db:"slug"`
	Image     string `db:"image"`
	Price     uint32 `db:"price"`
	Quantity  uint32 `db:"quantity"`
}

type DeleteCartProductInfo struct {
	ProductId string `db:"product_id"`
	Quantity  uint32 `db:"quantity"`
}
