package services

import (
	"encoding/csv"
	"os"

	customerrors "github.com/matiaseiglesias/storiChallenge/internal/custom_errors"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
)

// TransactionFileProcessorService defines the interface for processing transaction files service.
type TransactionFileProcessorService interface {
	ProcessTransactionFile(filename string) ([]models.TransactionCsv, error)
}

type TransactionFileProcessorImpl struct {
}

func CreateTransactionFileProcessor() *TransactionFileProcessorImpl {
	return &TransactionFileProcessorImpl{}
}

// ProcessTransactionFile reads and processes a CSV file containing transaction data.
// The function returns a slice of TransactionCsv models, each representing a transaction
// record from the file.
//
// Parameters:
//   - filename: The path to the CSV file containing the transactions.
//
// Returns:
//   - []models.TransactionCsv: A slice of TransactionCsv structs representing the transactions.
//   - error: An error is returned if there is an issue opening or reading the file.
//
// Possible Errors:
//   - *customerrors.FileOpeningError: Returned when there is an error opening the specified file.
//   - *customerrors.FileReadingError: Returned when there is an error reading the file or processing its contents.

func (s *TransactionFileProcessorImpl) ProcessTransactionFile(filename string) ([]models.TransactionCsv, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, &customerrors.FileOpeningError{Message: "couldnt open file"}
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields

	// Read and skip the header
	if _, err := reader.Read(); err != nil {
		return nil, &customerrors.FileReadingError{Message: "error while reading"}
	}

	data, err := reader.ReadAll()
	if err != nil {
		return nil, &customerrors.FileReadingError{Message: "error while reading"}
	}

	transactions := []models.TransactionCsv{}

	// Parse the CSV data
	for _, row := range data {
		transactions = append(transactions, models.TransactionCsv{
			Id:          row[0],
			Date:        row[1],
			Transaction: row[2],
		})
	}
	return transactions, nil
}
