package models

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v9"
)

type Product struct {
	ID          uint            `json:"id"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   null.Time       `json:"updated_at"`
	Name        string          `json:"name"`
	Sku         string          `json:"sku"`
	Price       decimal.Decimal `json:"price"`
	Description null.String     `json:"description"`
}
