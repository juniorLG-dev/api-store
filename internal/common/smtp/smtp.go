package smtp

import (
	gomail "gopkg.in/gomail.v2"
	"fmt"
)

const (
	smtpHost = "smtp.gmail.com"
	smtpPort = 587
)

type senderCredentials struct {
	email string
	password string
}

func NewSMTP(email, password string) *senderCredentials {
	return &senderCredentials{
		email: email,
		password: password,
	}
}

type PortSMTP interface {
	SendVerificationEmail(string, string) error
}

func (sc *senderCredentials) SendVerificationEmail(sellerEmail, code string) error {
	message := gomail.NewMessage()

	messageBody := fmt.Sprintf("<h1>Your verification code is: %v</h1>", code)

	message.SetHeader("From", sc.email)
	message.SetHeader("To", sellerEmail)
	message.SetHeader("Subject", "Verification code")
	message.SetBody("text/html", messageBody)

	smtpClient := gomail.NewDialer(smtpHost, smtpPort, sc.email, sc.password)

	if err := smtpClient.DialAndSend(message); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}