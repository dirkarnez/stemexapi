package utils

import (
	"fmt"
	"net/smtp"
)

func Testing() {
	email := "alex.chan@stemex.org"

	// Sender data.
	from := email
	password := "cyqb rllp qhep glnx"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf("To: %s\nSubject: discount Gophers!\nThis is the email body.\n", email))

	// client, err := smtp.Dial(smtpHost + ":" + smtpPort)
	// client.Auth()

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("Alex Chan", email, password, smtpHost), from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
