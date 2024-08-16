package services

import (
	"fmt"
	"net/smtp"

	"github.com/matiaseiglesias/storiChallenge/config"
	customerrors "github.com/matiaseiglesias/storiChallenge/internal/custom_errors"
)

type EmailSenderService interface {
	Send(to string, message []byte) error
}

type EmailSenderServiceImpl struct {
	From        string
	Credentials smtp.Auth
	Port        string
	Host        string
}

func CreateEmailSenderService(smtServerConfig config.Smtpserver) *EmailSenderServiceImpl {

	// Sender data
	from := smtServerConfig.From
	password := smtServerConfig.Password

	// SMTP server configuration
	smtpHost := smtServerConfig.Host
	smtpPort := smtServerConfig.Port

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	return &EmailSenderServiceImpl{
		From:        from,
		Credentials: auth,
		Port:        smtpPort,
		Host:        smtpHost,
	}
}

func (s *EmailSenderServiceImpl) Send(to string, message []byte) error {
	err := smtp.SendMail(s.Host+":"+s.Port, s.Credentials, s.From, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return &customerrors.EmailError{Message: "error while sending email, try later"}
	}
	return nil
}
