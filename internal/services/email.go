package services

import (
	"fmt"
	"html/template"
	"net/smtp"
	"os"

	"github.com/matiaseiglesias/storiChallenge/internal/models"
)

type EmailSenderService struct {
	From        string
	Credentials smtp.Auth
	Port        string
	Host        string
}

func CreateEmailSenderService() *EmailSenderService {

	// Sender data
	from := ""
	password := ""
	// password := os.Getenv("EMAIL_PASSWORD")

	// Receiver email address

	// SMTP server configuration
	smtpHost := "smtp.mailgun.org"
	smtpPort := "587"

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	return &EmailSenderService{
		From:        from,
		Credentials: auth,
		Port:        smtpPort,
		Host:        smtpHost,
	}
}

func (s *EmailSenderService) Send(to string) {
	subject := "Subject: Hello World\n"
	body := "Hello, World!"
	message := []byte(subject + "\n" + body)
	err := smtp.SendMail(s.Host+":"+s.Port, s.Credentials, s.From, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}
}

func (s *EmailSenderService) CreateSummaryTemplate() {

	t1 := models.TransactionsCount{
		Month:  "July",
		Amount: 10,
	}

	t2 := models.TransactionsCount{
		Month:  "June",
		Amount: 100,
	}

	st := models.Summary{
		TotalBalance:      10,
		TransactionsCount: []models.TransactionsCount{t1, t2},
		AverageCredit:     210,
		AverageDebit:      110,
	}
	name := "/home/matiasei/Documentos/storiChallenge/mailTemplate/summary.html"
	tmpl, err := template.ParseFiles(name)
	if err != nil {
		panic(err)
	}
	// Execute the template with your data
	err = tmpl.Execute(os.Stdout, st)
	if err != nil {
		panic(err)
	}
}
