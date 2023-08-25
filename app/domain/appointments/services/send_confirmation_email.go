package services

import (
	"crypto/tls"
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func SendConfirmationEmail(email string) bool {
	m := gomail.NewMessage()

	m.SetHeader("From", "shopper.tiago@gmail.com")
	m.SetHeader("to", email)
	m.SetHeader("Subject", "Link de confirmação: Bem-vindo a WA")

	m.SetBody("text/plain", "Segue link de confirmação: ")

	d := gomail.NewDialer("smtp.gmail.com", 587, "shopper.tiago@gmail.com", "$N57ntlctl")

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {

		fmt.Println("Err: ", err.Error())

		fmt.Println("Email not sent")

		return false
	}

	return true
}
