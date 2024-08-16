package dto

// Response
//
// This is a generic response structure.
//
// swagger:model Response
type Response struct {
	// The status of the response, indicating success or error.
	// Example: "success"
	// in: body
	Status string `json:"status"`

	// A message providing additional details about the response.
	// Example: "Email sent successfully"
	// in: body
	Message string `json:"message"`
}
