package services

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/matiaseiglesias/storiChallenge/config"
	customerrors "github.com/matiaseiglesias/storiChallenge/internal/custom_errors"
	"github.com/matiaseiglesias/storiChallenge/internal/models"
)

type TransactionSummaryEmailService interface {
	CreateSummaryTemplate(data *models.Summary) ([]byte, error)
}

type TransactionSummaryEmailServiceImpl struct {
	templatePath string
}

func CreateTransactionSummaryEmailService(config config.EmailTemplate) *TransactionSummaryEmailServiceImpl {
	return &TransactionSummaryEmailServiceImpl{
		templatePath: config.Path,
	}
}

func (s *TransactionSummaryEmailServiceImpl) CreateSummaryTemplate(data *models.Summary) ([]byte, error) {
	name := s.templatePath
	if name == "" {
		return nil, &customerrors.DiractoryError{Message: "mail template diractory not configured"}
	}
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		log.Println(err)
		return nil, &customerrors.DiractoryError{Message: "error while parsing template"}
	}

	// Execute the template with data
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	body.Write([]byte(fmt.Sprintf("Subject: Your Transaction\n%s\n\n", mimeHeaders)))
	err = tmpl.Execute(&body, data)
	if err != nil {
		return nil, &customerrors.TemplateError{Message: "error while filling template data"}
	}
	return body.Bytes(), nil
}
