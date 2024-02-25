package models

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/volatiletech/null/v9"
)

type Order struct {
	Id          uint            `json:"id"`
	CreatedOn   time.Time       `json:"created_on"`
	ModifiedOn  null.Time       `json:"modified_on"`
	Items       []OrderItem     `json:"items"`
	Code        string          `json:"code"`
	Total       decimal.Decimal `json:"total"`
	Description null.String     `json:"description"`
}

type OrderItem struct {
	Id          uint            `json:"id"`
	CreatedOn   time.Time       `json:"created_on"`
	ModifiedOn  null.Time       `json:"modified_on"`
	Product     Product         `json:"product"`
	Quantity    uint            `json:"quantity"`
	Price       decimal.Decimal `json:"price"`
	Description null.String     `json:"description"`
}
