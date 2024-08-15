package services

type Services struct {
	EmailSender *EmailSenderService
	Transaction *TransactionsService
}

func CreateServices() *Services {
	var email = CreateEmailSenderService()
	var transaction = CreateTransactionsService()

	return &Services{

		EmailSender: email,
		Transaction: transaction,
	}
}
