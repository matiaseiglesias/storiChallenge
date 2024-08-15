package models

type TransactionsCount struct {
	Month  string
	Amount uint
}

type Summary struct {
	TotalBalance      float64
	TransactionsCount []TransactionsCount
	AverageCredit     float64
	AverageDebit      float64
}
