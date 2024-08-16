package repositories

import "github.com/matiaseiglesias/storiChallenge/internal/database"

type Repositories struct {
	TransactionsSummary TransactionsSummaryRepository
}

func CreateRepositories(db *database.DataBase) *Repositories {
	return &Repositories{
		TransactionsSummary: CreateTransactionsSummaryRepository(db),
	}

}
