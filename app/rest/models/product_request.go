package rest_models

import (
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v9"
)

type ProductRequest struct {
	Name        string          `json:"name"`
	Sku         string          `json:"sku"`
	Price       decimal.Decimal `json:"price"`
	Description null.String     `json:"description"`
}
