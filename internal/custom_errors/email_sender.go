package customerrors

type EmailError struct {
	Message string
}

func (e *EmailError) Error() string {
	return e.Message
}
