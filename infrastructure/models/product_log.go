package models

import (
	"time"

	"github.com/volatiletech/null/v9"
)

type ProductLog struct {
	Id          uint        `json:"id"`
	CreatedOn   time.Time   `json:"created_on"`
	ProductId   uint        `json:"product_id"`
	Event       string      `json:"event"`
	Description null.String `json:"description"`
}
