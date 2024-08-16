package dto

import "gorm.io/gorm"

type TransactionsCountDto struct {
	gorm.Model
	Month  string
	Amount uint
}

type SummaryDto struct {
	gorm.Model
	TotalBalance      string
	TransactionsCount []TransactionsCountDto `gorm:"foreignKey:Month"`
	AverageCredit     string
	AverageDebit      string
}
