package models

import (
	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v9"
)

type Product struct {
	Id          uint            `json:"id"`
	CreatedOn   null.Time       `json:"created_on"`
	ModifiedOn  null.Time       `json:"modified_on"`
	Name        string          `json:"name"`
	Sku         string          `json:"sku"`
	Price       decimal.Decimal `json:"price"`
	Description null.String     `json:"description"`
}
