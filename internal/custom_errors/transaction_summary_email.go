package customerrors

type DiractoryError struct {
	Message string
}

func (e *DiractoryError) Error() string {
	return e.Message
}

type TemplateError struct {
	Message string
}

func (e *TemplateError) Error() string {
	return e.Message
}
