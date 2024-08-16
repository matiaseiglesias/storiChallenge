package repositories

import (
	"github.com/matiaseiglesias/storiChallenge/internal/database"
	"github.com/matiaseiglesias/storiChallenge/internal/dto"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
)

type TransactionsSummaryRepository interface {
	SaveTransactionSummary(summary *models.Summary)
}

type TransactionsSummaryRepositoryImpl struct {
	db *database.DataBase
}

func CreateTransactionsSummaryRepository(db *database.DataBase) *TransactionsSummaryRepositoryImpl {

	return &TransactionsSummaryRepositoryImpl{
		db: db,
	}
}

func (r *TransactionsSummaryRepositoryImpl) GetTransactionSummary(filename string) {
}

func (r *TransactionsSummaryRepositoryImpl) SaveTransactionSummary(summary *models.Summary) {
	transactionsCount := []dto.TransactionsCountDto{}
	for _, v := range summary.TransactionsCount {
		transactionsCount = append(transactionsCount, dto.TransactionsCountDto{
			Month:  v.Month,
			Amount: v.Amount,
		})
	}
	summary_ := &dto.SummaryDto{
		TotalBalance:      summary.TotalBalance.String(),
		AverageCredit:     summary.AverageCredit.String(),
		AverageDebit:      summary.AverageDebit.String(),
		TransactionsCount: transactionsCount,
	}
	r.db.Db.Create(summary_)
}
