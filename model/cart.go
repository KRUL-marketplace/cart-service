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
	ItemID    string          `db:"item_id"`
	ProductID string          `db:"product_id"`
	Info      CartProductInfo `db:"info"`
	Quantity  uint32          `db:"quantity"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt sql.NullTime    `db:"updated_at"`
}

type CartProductInfo struct {
	Name  string `db:"name"`
	Slug  string `db:"slug"`
	Image string `db:"image"`
	Price uint32 `db:"price"`
	Brand Brand  `db:"brand"`
}

type AddProductRequest struct {
	UserID    string `db:"user_id"`
	ProductID string `db:"product_id"`
	Quantity  uint32 `db:"quantity"`
}

type DeleteProductRequest struct {
	UserID    string `db:"user_id"`
	ProductID string `db:"product_id"`
	Quantity  uint32 `db:"quantity"`
}

type Brand struct {
	ID        uint32       `db:"id"`
	Info      BrandInfo    `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type BrandInfo struct {
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	Description string `db:"description"`
}

type DeleteCartProductInfo struct {
	ProductId string `db:"product_id"`
	Quantity  uint32 `db:"quantity"`
}
