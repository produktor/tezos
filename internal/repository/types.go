package repository

import "github.com/shopspring/decimal"

type TgGroupDB struct {
	ID          int64
	Title       string
	Description string
	Price       decimal.Decimal
}
