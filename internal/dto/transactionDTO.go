package dto

// SummaryResponse
//
// # Transactions summary response
//
// swagger:response SummaryResponse
type SummaryResponse struct {
	// This text will appear as description of your request body.
	Status string
}

// SummaryResponse
//
// # Transactions summary response
//
// swagger:model
type TransactionCsv struct {
	Id          string
	Date        string
	Transaction string
}
