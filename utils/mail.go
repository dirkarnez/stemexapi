package utils

import (
	"fmt"
	"net/smtp"
)

func SendActivationHTMLEmail(toEmail, activationLinkURL string) error {
	return sendEmail(toEmail, "Activation", fmt.Sprintf(`<html><body>Please click this <a href="%s" target="_blank">link</a> for activation</body></html>`, activationLinkURL))
}

func sendEmail(to, subject, body string) error {
	// Sender data.
	from := "alex.chan@stemex.org"
	password := "cyqb rllp qhep glnx"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf("To: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s\n", to, subject, body))

	// client, err := smtp.Dial(smtpHost + ":" + smtpPort)
	// client.Auth()

	// Sending email.
	return smtp.SendMail(smtpHost+":"+smtpPort, smtp.PlainAuth("Alex Chan", from, password, smtpHost), from, []string{to}, message)
}
