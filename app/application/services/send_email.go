package services

import (
	"crypto/tls"
	"net/smtp"
)

// func SendConfirmationEmail(email string, appointmentId string) bool {
// 	m := gomail.NewMessage()

// 	m.SetHeader("From", "shopper.tiago@gmail.com")
// 	m.SetHeader("to", email)

// 	confirmationLinkMessage := fmt.Sprintf("Link de confirmação: http://localhost:3000/%s", appointmentId)
// 	m.SetHeader("Subject", confirmationLinkMessage)

// 	m.SetBody("text/plain", "Segue link de confirmação: ")

// 	d := gomail.NewDialer("smtp.gmail.com", 587, "shopper.tiago@gmail.com", "lgivvrjvybnnipti")

// 	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

// 	if err := d.DialAndSend(m); err != nil {

// 		fmt.Println("Err: ", err.Error())

// 		fmt.Println("Email not sent")

// 		return false
// 	}

// 	return true
// }

func SendEmail(to string, subject string, body string) bool {
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
		return false
	}

	//retorna client SMTP
	c, err := smtp.NewClient(conn, host)
	if checkErr(err) {
		return false
	}

	//autentica
	err = c.Auth(auth)
	if checkErr(err) {
		return false
	}

	//adiciona remetente
	err = c.Mail("shopper.tiago@gmail.com")
	if checkErr(err) {
		return false
	}

	//adiciona destinatários
	err = c.Rcpt(to)
	if checkErr(err) {
		return false
	}

	//prepara corpo do email
	w, err := c.Data()
	if checkErr(err) {
		return false
	}

	//adiciona corpo do e-mail
	_, err = w.Write([]byte(msg))
	if checkErr(err) {
		return false
	}

	//fecha corpo do e-mail
	err = w.Close()
	if checkErr(err) {
		return false
	}

	//encerra conexão
	c.Quit()
	return true
}

func checkErr(err error) bool {
	if err != nil {
		return true
	}

	return false
}
