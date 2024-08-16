package models

import "github.com/shopspring/decimal"

type TransactionsCount struct {
	Month  string
	Amount uint
}

type Summary struct {
	TotalBalance      decimal.Decimal
	TransactionsCount []TransactionsCount
	AverageCredit     decimal.Decimal
	AverageDebit      decimal.Decimal
}
