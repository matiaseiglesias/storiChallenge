package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/matiaseiglesias/storiChallenge/internal/dto"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
)

type TransactionsService struct {
}

func CreateTransactionsService() *TransactionsService {
	return &TransactionsService{}
}

func (s *TransactionsService) ProccesTransactionFile(filename string) []dto.TransactionCsv {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read and skip the header
	if _, err := reader.Read(); err != nil {
		fmt.Println("Error reading header:", err)
	}

	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	transactions := []dto.TransactionCsv{}

	// Print the CSV data
	for _, row := range data {
		transactions = append(transactions, dto.TransactionCsv{
			Id:          row[0],
			Date:        row[1],
			Transaction: row[2],
		})
		fmt.Println()
	}
	return transactions
}

func (s *TransactionsService) CalculateSummary(filename string) *models.Summary {

	debitsAmount := 0.0
	debitsCount := 0
	creditsAmount := 0.0
	creditsCount := 0
	transactionsByMonth := make(map[string]uint)

	transactions := s.ProccesTransactionFile(filename)
	for _, transaction := range transactions {
		i, _ := strconv.ParseFloat(strings.TrimSpace(transaction.Transaction[1:]), 64)
		monthName := GetMonthName(transaction.Date)
		transactionsByMonth[monthName] += 1
		if transaction.Transaction[0] == '+' {
			debitsAmount += i
			debitsCount += 1
		} else {
			creditsAmount += i
			creditsCount += 1
		}
	}
	summary := models.Summary{
		TotalBalance:      debitsAmount - creditsAmount,
		AverageCredit:     creditsAmount / float64(creditsCount),
		AverageDebit:      debitsAmount / float64(debitsCount),
		TransactionsCount: mapTransactionsCount(&transactionsByMonth),
	}
	log.Println(transactionsByMonth)
	return &summary
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
