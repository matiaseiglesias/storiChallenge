package models

import "time"

type TransactionType int

const (
	Debit TransactionType = iota
	Credit
)

type Transaction struct {
	Id     uint32
	Date   time.Time
	Amount float32
	Type   TransactionType
}
