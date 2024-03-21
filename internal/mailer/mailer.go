package mailer

import (
	"btc-network-monitor/internal/logger"
	"bytes"
	"crypto/tls"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	FirstName string
	Subject string
	MailTo string
	Confirmations int64
	TxId string
}


// 👇 Email template parser
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logger.Error("Could not parse template" + err.Error())
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		logger.Error("Could not parse template" + err.Error())
		return nil, err
	}

	return template.ParseFiles(paths...)
}

// ? Email template parser
func SendEmail(data *EmailData) error {
	// Sender data.
	from := os.Getenv("EMAIL_FROM")
	smtpPass := os.Getenv("SMTP_PASS")
	smtpUser := os.Getenv("SMTP_USER")
	smtpHost := os.Getenv("SMTP_HOST")
	portString := os.Getenv("SMTP_PORT")

	smtpPort, err := strconv.Atoi(portString)
	if err != nil {
        // ... handle error
		logger.Error("Could not covert port string" + err.Error())
    }

	var body bytes.Buffer
	template, err := ParseTemplateDir("internal/views/templates")
	if err != nil {
		logger.Error("Could not parse template" + err.Error())
	}

	template.ExecuteTemplate(&body, "notification.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", data.MailTo)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	logger.Info(d)

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		logger.Error("Could not send mail" + err.Error())
		return err
	}
	return nil
}