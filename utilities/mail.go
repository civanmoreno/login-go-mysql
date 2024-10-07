package utilities

import (
	"fmt"
	"os"

	"net/smtp"

	"github.com/joho/godotenv"
)

func SendMail(to string, subject string, body string) error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	// Retrieve SMTP configuration from environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpFrom := os.Getenv("SMTP_FROM")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	message := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", to, subject, body))

	recipient := []string{to}

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpFrom, recipient, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
