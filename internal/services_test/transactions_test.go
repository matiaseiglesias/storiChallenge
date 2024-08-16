package servicestest

import (
	"testing"

	"github.com/matiaseiglesias/storiChallenge/config"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
	"github.com/matiaseiglesias/storiChallenge/internal/repositories"
	"github.com/matiaseiglesias/storiChallenge/internal/services"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
)

var transactionDataset = []models.TransactionCsv{
	{
		Id:          "1",
		Date:        "13/08",
		Transaction: "+75.00",
	},
	{
		Id:          "2",
		Date:        "12/08",
		Transaction: "-35.50",
	},
	{
		Id:          "3",
		Date:        "11/08",
		Transaction: "+120.25",
	},
	{
		Id:          "4",
		Date:        "10/08",
		Transaction: "-10.00",
	},
	{
		Id:          "5",
		Date:        "09/07",
		Transaction: "+150.75",
	},
	{
		Id:          "6",
		Date:        "08/07",
		Transaction: "-60.25",
	},
	{
		Id:          "7",
		Date:        "07/07",
		Transaction: "+90.50",
	},
	{
		Id:          "8",
		Date:        "06/06",
		Transaction: "-45.00",
	},
	{
		Id:          "9",
		Date:        "05/06",
		Transaction: "+130.80",
	},
	{
		Id:          "10",
		Date:        "04/06",
		Transaction: "-25.75",
	},
	{
		Id:          "11",
		Date:        "03/08",
		Transaction: "+100.75",
	},
	{
		Id:          "12",
		Date:        "02/08",
		Transaction: "-50.00",
	},
	{
		Id:          "13",
		Date:        "01/08",
		Transaction: "+120.30",
	},
	{
		Id:          "14",
		Date:        "31/07",
		Transaction: "-30.20",
	},
	{
		Id:          "15",
		Date:        "30/07",
		Transaction: "+200.00",
	},
	{
		Id:          "16",
		Date:        "29/07",
		Transaction: "-15.00",
	},
	{
		Id:          "17",
		Date:        "28/07",
		Transaction: "+75.25",
	},
	{
		Id:          "18",
		Date:        "27/07",
		Transaction: "-40.75",
	},
	{
		Id:          "19",
		Date:        "26/06",
		Transaction: "+180.50",
	},
	{
		Id:          "20",
		Date:        "25/06",
		Transaction: "-25.25",
	},
}

type MockTransactionFileProcessor struct {
	mock.Mock
}

func (m *MockTransactionFileProcessor) ProcessTransactionFile(filename string) ([]models.TransactionCsv, error) {
	args := m.Called(filename)
	return args.Get(0).([]models.TransactionCsv), args.Error(1)
}

func TestTransaction(t *testing.T) {

	t.Run("Calculate transaction summary with empty file", func(t *testing.T) {
		config := config.TransactionFile{
			Directory: "",
		}
		fileProcessor := new(MockTransactionFileProcessor)
		fileProcessor.On("ProcessTransactionFile", ".csv").Return([]models.TransactionCsv{}, nil)
		r := &repositories.TransactionsSummaryRepositoryImpl{}
		emailSender := &services.EmailSenderServiceImpl{}
		summaryEmail := &services.TransactionSummaryEmailServiceImpl{}

		service := services.CreateTransactionsService(config, r, fileProcessor, emailSender, summaryEmail)

		summary, _ := service.CalculateSummary("")
		if !summary.AverageCredit.Equal(decimal.NewFromInt(0)) {
			t.Error("error while calculating empty file, average credit must be 0, result: ", summary.AverageCredit)
		}
		if !summary.AverageDebit.Equal(decimal.NewFromInt(0)) {
			t.Error("error while calculating empty file, average debit must be 0, result: ", summary.AverageDebit)
		}
		if !summary.TotalBalance.Equal(decimal.NewFromInt(0)) {
			t.Error("error while calculating empty file, total balance must be 0, result: ", summary.TotalBalance)
		}
		if len(summary.TransactionsCount) != 0 {
			t.Error("error while calculating empty file, the number of months must be 0, result: ", len(summary.TransactionsCount))
		}
	})

	t.Run("Calculate transaction summary with only one debit", func(t *testing.T) {
		fileProcessor := new(MockTransactionFileProcessor)
		fileProcessor.On("ProcessTransactionFile", ".csv").Return([]models.TransactionCsv{
			{
				Id:          "1",
				Date:        "02/05",
				Transaction: "+10",
			},
		}, nil)
		config := config.TransactionFile{
			Directory: "",
		}
		r := &repositories.TransactionsSummaryRepositoryImpl{}
		emailSender := &services.EmailSenderServiceImpl{}
		summaryEmail := &services.TransactionSummaryEmailServiceImpl{}

		service := services.CreateTransactionsService(config, r, fileProcessor, emailSender, summaryEmail)

		summary, _ := service.CalculateSummary("")
		if !summary.AverageCredit.Equal(decimal.NewFromInt(0)) {
			t.Error("error while calculating file with one debit, average credit must be 0, result: ", summary.AverageCredit)
		}
		if !summary.AverageDebit.Equal(decimal.NewFromInt(10)) {
			t.Error("error while calculating file with one debit, average debit must be 10, result: ", summary.AverageDebit)
		}
		if !summary.TotalBalance.Equal(decimal.NewFromInt(10)) {
			t.Error("error while calculating file with one debit, total balance must be 10, result: ", summary.TotalBalance)
		}
		if len(summary.TransactionsCount) != 1 {
			t.Error("error while calculating empty file, the number of months must be 1, result: ", len(summary.TransactionsCount))
		}
	})

	t.Run("Calculate transaction summary with only one credit", func(t *testing.T) {
		fileProcessor := new(MockTransactionFileProcessor)
		fileProcessor.On("ProcessTransactionFile", ".csv").Return([]models.TransactionCsv{
			{
				Id:          "1",
				Date:        "02/05",
				Transaction: "-10",
			},
		}, nil)
		config := config.TransactionFile{
			Directory: "",
		}
		r := &repositories.TransactionsSummaryRepositoryImpl{}
		emailSender := &services.EmailSenderServiceImpl{}
		summaryEmail := &services.TransactionSummaryEmailServiceImpl{}

		service := services.CreateTransactionsService(config, r, fileProcessor, emailSender, summaryEmail)

		summary, _ := service.CalculateSummary("")
		if !summary.AverageCredit.Equal(decimal.NewFromInt(10)) {
			t.Error("error while calculating file with one credit, average credit must be 10, result: ", summary.AverageCredit)
		}
		if !summary.AverageDebit.Equal(decimal.NewFromInt(0)) {
			t.Error("error while calculating file with one credit, average debit must be 0, result: ", summary.AverageDebit)
		}
		if !summary.TotalBalance.Equal(decimal.NewFromInt(-10)) {
			t.Error("error while calculating file with one credit, total balance must be -10, result: ", summary.TotalBalance)
		}
		if len(summary.TransactionsCount) != 1 {
			t.Error("error while calculating file with one credit, the number of months must be 1, result: ", len(summary.TransactionsCount))
		}
	})

	t.Run("Calculate transaction summary with credits and debits", func(t *testing.T) {
		fileProcessor := new(MockTransactionFileProcessor)
		fileProcessor.On("ProcessTransactionFile", ".csv").Return(transactionDataset, nil)
		config := config.TransactionFile{
			Directory: "",
		}
		r := &repositories.TransactionsSummaryRepositoryImpl{}
		emailSender := &services.EmailSenderServiceImpl{}
		summaryEmail := &services.TransactionSummaryEmailServiceImpl{}

		service := services.CreateTransactionsService(config, r, fileProcessor, emailSender, summaryEmail)

		summary, _ := service.CalculateSummary("")
		if !summary.AverageCredit.Equal(decimal.NewFromFloat(33.77)) {
			t.Error("error while calculating summary, average credit must be 33.77, result: ", summary.AverageCredit)
		}
		if !summary.AverageDebit.Equal(decimal.NewFromFloat(124.41)) {
			t.Error("error while calculating summary, average debit must be 124.41, result: ", summary.AverageDebit)
		}
		if !summary.TotalBalance.Equal(decimal.NewFromFloat(906.4)) {
			t.Error("error while calculating summary, total balance must be 906.4, result: ", summary.TotalBalance)
		}
		if len(summary.TransactionsCount) != 3 {
			t.Error("error while calculating summary, the number of months must be 3, result: ", len(summary.TransactionsCount))
		}
	})
	t.Run("Checking correct transaction count by month", func(t *testing.T) {
		fileProcessor := new(MockTransactionFileProcessor)
		fileProcessor.On("ProcessTransactionFile", ".csv").Return(transactionDataset, nil)
		config := config.TransactionFile{
			Directory: "",
		}
		r := &repositories.TransactionsSummaryRepositoryImpl{}
		emailSender := &services.EmailSenderServiceImpl{}
		summaryEmail := &services.TransactionSummaryEmailServiceImpl{}

		service := services.CreateTransactionsService(config, r, fileProcessor, emailSender, summaryEmail)

		summary, _ := service.CalculateSummary("")
		if !summary.AverageCredit.Equal(decimal.NewFromFloat(33.77)) {
			t.Error("error while calculating summary, average credit must be 33.77, result: ", summary.AverageCredit)
		}
		if !summary.AverageDebit.Equal(decimal.NewFromFloat(124.41)) {
			t.Error("error while calculating summary, average debit must be 124.41, result: ", summary.AverageDebit)
		}
		if !summary.TotalBalance.Equal(decimal.NewFromFloat(906.4)) {
			t.Error("error while calculating summary, total balance must be 906.4, result: ", summary.TotalBalance)
		}
		if len(summary.TransactionsCount) != 3 {
			t.Error("error while calculating summary, the number of months must be 3, result: ", len(summary.TransactionsCount))
		}

		for _, v := range summary.TransactionsCount {
			if v.Month == "June" && v.Amount != 5 {
				t.Error("error while calculating the number of transaction in June. Must be 5, result: ", v.Amount)
			} else if v.Month == "August" && v.Amount != 7 {
				t.Error("error while calculating the number of transaction in August. Must be 7, result: ", v.Amount)
			} else if v.Month == "July" && v.Amount != 8 {
				t.Error("error while calculating the number of transaction in July. Must be 8, result: ", v.Amount)
			} else if v.Month != "July" && v.Month != "August" && v.Month != "June" {
				t.Error("Unexpected case, result: ", v)
			}
		}
	})
}
