package email

import (
	"bytes"
	"fmt"
	"io"

	"gopkg.in/mail.v2"
)

type EmailInfo struct {
	EmailSender    string
	Password       string
	EmailRecipient string
	EmailSubject   string
}

const (
	gmailSmtpHost = "smtp.gmail.com"
	gmailSmtpPort = 587
)

func NewEmailInfo(sender, password, recipient, subject string) EmailInfo {
	return EmailInfo{
		EmailSender:    sender,
		Password:       password,
		EmailRecipient: recipient,
		EmailSubject:   subject,
	}
}
func SendGmail(imageBuffer *bytes.Buffer, emailInfo EmailInfo) {
	msg := mail.NewMessage()
	msg.SetHeader("From", emailInfo.EmailSender)
	msg.SetHeader("To", emailInfo.EmailRecipient)
	msg.SetHeader("Subject", emailInfo.EmailSubject)
	msg.SetBody("text/html", "Here is your screenshot")
	msg.Attach("screenshot.png", mail.SetCopyFunc(func(writer io.Writer) error {
		_, err := writer.Write(imageBuffer.Bytes())
		return err
	}))

	dialer := mail.NewDialer(gmailSmtpHost, gmailSmtpPort, emailInfo.EmailSender, emailInfo.Password)
	if err := dialer.DialAndSend(msg); err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
	} else {
		fmt.Println("Screenshot sent successfully")
	}
}
