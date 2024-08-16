package dto

// SummaryRequest
//
// This is the request body for the transaction summary notification.
//
// swagger:parameters makeSummary
type SummaryRequestSwagger struct {
	// in: body
	// Account is the identifier for the account for which the transaction summary is requested.
	// Email is the recipient's email address where the summary will be sent.
	// Required: true
	Data SummaryRequest
}

type SummaryRequest struct {
	Account string `json:"account"`
	Email   string `json:"email"`
}
