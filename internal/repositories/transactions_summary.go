package repositories

import (
	"github.com/matiaseiglesias/storiChallenge/internal/database"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
)

type TransactionsSummaryRepository struct {
	db *database.DataBase
}

func CreateTransactionsSummaryRepository(db *database.DataBase) *TransactionsSummaryRepository {

	return &TransactionsSummaryRepository{
		db: db,
	}
}

func (r *TransactionsSummaryRepository) GetTransactionSummary(filename string) {
}

func (r *TransactionsSummaryRepository) SaveTransactionSummary(summary models.Transaction) {

}
