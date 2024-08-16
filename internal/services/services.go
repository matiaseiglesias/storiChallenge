package services

import (
	"github.com/matiaseiglesias/storiChallenge/config"
	"github.com/matiaseiglesias/storiChallenge/internal/repositories"
)

type Services struct {
	EmailSender                    EmailSenderService
	Transaction                    TransactionService
	TransactionFileProcessor       TransactionFileProcessorService
	TransactionSummaryEmailService TransactionSummaryEmailService
}

func CreateServices(config *config.Config, repositories *repositories.Repositories) *Services {
	var email = CreateEmailSenderService(config.SmtpServer)
	var transactionFileProcessor = CreateTransactionFileProcessor()
	var transactionSummaryEmail = CreateTransactionSummaryEmailService(config.EmailTemplate)
	var transaction = CreateTransactionsService(config.TransactionFile, repositories.TransactionsSummary, transactionFileProcessor, email, transactionSummaryEmail)

	return &Services{
		EmailSender:                    email,
		Transaction:                    transaction,
		TransactionFileProcessor:       transactionFileProcessor,
		TransactionSummaryEmailService: transactionSummaryEmail,
	}
}
