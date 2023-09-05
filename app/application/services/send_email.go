package services

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail(to string, subject string, body string) error {
	//configuração
	servername := "smtp.gmail.com:465"                                    //servidor SMTP e PORTA
	host := "smtp.gmail.com"                                              //host
	pass := "lgivvrjvybnnipti"                                            //senha
	auth := smtp.PlainAuth("Curi", "shopper.tiago@gmail.com", pass, host) //autenticação
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	//body := fmt.Sprintf("http://localhost:3000/confirmed?apnmnt=%s", appointmentId)
	msg := "From: " + "shopper.tiago@gmail.com" + "\n" + "To: " + to + "\n" + subject + body

	//conecta com o servidor SMTP
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	if checkErr(err) {
		return err
	}

	//retorna client SMTP
	c, err := smtp.NewClient(conn, host)
	if checkErr(err) {
		return err
	}

	//autentica
	err = c.Auth(auth)
	if checkErr(err) {
		return err
	}

	//adiciona remetente
	err = c.Mail("shopper.tiago@gmail.com")
	if checkErr(err) {
		return err
	}

	to = strings.Trim(to, " ")

	fmt.Println("To: ", to)
	//adiciona destinatários
	err = c.Rcpt(to)
	if checkErr(err) {
		return err
	}

	//prepara corpo do email
	w, err := c.Data()
	if checkErr(err) {
		return err
	}

	//adiciona corpo do e-mail
	_, err = w.Write([]byte(msg))
	if checkErr(err) {
		return err
	}

	//fecha corpo do e-mail
	err = w.Close()
	if checkErr(err) {
		return err
	}

	//encerra conexão
	c.Quit()
	return nil
}

func checkErr(err error) bool {
	if err != nil {
		return true
	}

	return false
}
