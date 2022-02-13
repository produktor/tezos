package model

import "github.com/shopspring/decimal"

type TelegramGroup struct {
	ID          int64
	Title       *string
	Description *string
	Price       decimal.Decimal
}

type TelegramLinks struct {
	TgGroup *TelegramGroup
	Link string
}