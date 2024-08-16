package services

import (
	"log"
	"strconv"
	"strings"

	"github.com/matiaseiglesias/storiChallenge/config"
	customerrors "github.com/matiaseiglesias/storiChallenge/internal/custom_errors"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
	"github.com/matiaseiglesias/storiChallenge/internal/repositories"
	"github.com/shopspring/decimal"
)

type TransactionService interface {
	NotifyTransactionSummary(account string, email string) error
}

type TransactionServiceImpl struct {
	transactionDirectory            string
	summaryRepository               repositories.TransactionsSummaryRepository
	transactionFileProcessorService TransactionFileProcessorService
	emailSenderService              EmailSenderService
	transactionSummaryEmailService  TransactionSummaryEmailService
}

func CreateTransactionsService(config config.TransactionFile, ransactionsSummaryRepository repositories.TransactionsSummaryRepository, transactionFileProcessor TransactionFileProcessorService, emailSender EmailSenderService, transactionSummary TransactionSummaryEmailService) *TransactionServiceImpl {
	return &TransactionServiceImpl{
		transactionDirectory:            config.Directory,
		summaryRepository:               ransactionsSummaryRepository,
		transactionFileProcessorService: transactionFileProcessor,
		emailSenderService:              emailSender,
		transactionSummaryEmailService:  transactionSummary,
	}
}

// NotifyTransactionSummary generates a transaction summary for a given account
// and sends it via email.
//
// Parameters:
//   - account: The account identifier for which the transaction summary will be calculated.
//   - email: The recipient's email address where the summary will be sent.
//
// Returns:
//   - error: An error is returned if there is an issue with input validation,
//     summary calculation, email template creation, or email sending.
func (s *TransactionServiceImpl) NotifyTransactionSummary(account string, email string) error {

	if account == "" || email == "" {
		return &customerrors.EmptyFieldError{Message: "Empty field"}
	}
	summary, err := s.CalculateSummary(account)
	if err != nil {
		log.Println(err)
		return err
	}
	emailContent, err := s.transactionSummaryEmailService.CreateSummaryTemplate(summary)
	if err != nil {
		log.Println(err)
		return err
	}
	s.summaryRepository.SaveTransactionSummary(summary)

	err = s.emailSenderService.Send(email, emailContent)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CalculateSummary processes a transaction file associated with the given account
// and calculates a summary of the transactions. The summary includes the total balance,
// average credit, average debit, and the count of transactions per month.
//
// Parameters:
//   - account: The account identifier used to locate the transaction file.
//
// Returns:
//   - *models.Summary: A pointer to the Summary struct containing the calculated summary details.
//   - error: An error is returned if there is an issue processing the transaction file.
func (s *TransactionServiceImpl) CalculateSummary(account string) (*models.Summary, error) {

	debitsAmount := decimal.NewFromInt(0)
	debitsCount := int32(0)
	creditsAmount := decimal.NewFromInt(0)
	creditsCount := int32(0)
	transactionsByMonth := make(map[string]uint)

	filename := s.transactionDirectory + account + ".csv"

	transactions, err := s.transactionFileProcessorService.ProcessTransactionFile(filename)
	if err != nil {
		return nil, &customerrors.ProcessTransactionError{Message: "Error while processing transactions"}
	}

	for _, transaction := range transactions {
		i, _ := strconv.ParseFloat(strings.TrimSpace(transaction.Transaction[1:]), 64)
		monthName := GetMonthName(transaction.Date)
		transactionsByMonth[monthName] += 1
		if transaction.Transaction[0] == '+' {
			debitsAmount = debitsAmount.Add(decimal.NewFromFloat(i))
			debitsCount += 1
		} else {
			creditsAmount = creditsAmount.Add(decimal.NewFromFloat(i))
			creditsCount += 1
		}
	}
	avarageCredit := decimal.NewFromFloat(0.0)
	if creditsCount > 0 {
		avarageCredit = creditsAmount.Div(decimal.NewFromInt32(creditsCount))
	}
	avarageDebit := decimal.NewFromFloat(0.0)
	if debitsCount > 0 {
		avarageDebit = debitsAmount.Div(decimal.NewFromInt32(debitsCount))
	}
	summary := &models.Summary{
		TotalBalance:      debitsAmount.Sub(creditsAmount),
		AverageCredit:     avarageCredit,
		AverageDebit:      avarageDebit,
		TransactionsCount: mapTransactionsCount(&transactionsByMonth),
	}
	return summary, nil
}

func mapTransactionsCount(transactionsByMonth *map[string]uint) []models.TransactionsCount {
	transactionsCountByMonth := []models.TransactionsCount{}
	for k, v := range *transactionsByMonth {
		transactionsCountByMonth = append(transactionsCountByMonth, models.TransactionsCount{
			Month:  k,
			Amount: v,
		})

	}
	return transactionsCountByMonth
}
