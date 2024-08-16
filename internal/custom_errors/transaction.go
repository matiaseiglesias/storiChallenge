package customerrors

type EmptyFieldError struct {
	Message string
}

func (e *EmptyFieldError) Error() string {
	return e.Message
}

type SummaryError struct {
	Message string
}

func (e *SummaryError) Error() string {
	return e.Message
}

type ProcessTransactionError struct {
	Message string
}

func (e *ProcessTransactionError) Error() string {
	return e.Message
}
