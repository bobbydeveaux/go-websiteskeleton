package common

import (
    "fmt"
    "net/smtp"
    "strings"
)

type Message struct {
    Email string
    Name string
    Content string
    Contact string
    Heardvia string
    Require string
}

func (msg *Message) Deliver() error {
  to := []string{"info@dvomedia.net"}

  message := []string{msg.Name, msg.Email, msg.Contact, msg.Heardvia, msg.Require, msg.Content};
  content := fmt.Sprintf(strings.Join(message, " \r\n "));

  body := fmt.Sprintf("Reply-To: %v\r\nSubject: New Enquiry\r\n%v", msg.Email, content)

  username := "user@gmail.com"
  password := "password"
  auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")

  return smtp.SendMail("smtp.gmail.com:587", auth, msg.Email, to, []byte(body))
}