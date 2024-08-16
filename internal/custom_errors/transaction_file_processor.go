package customerrors

type FileOpeningError struct {
	Message string
}

func (e *FileOpeningError) Error() string {
	return e.Message
}

type FileReadingError struct {
	Message string
}

func (e *FileReadingError) Error() string {
	return e.Message
}
