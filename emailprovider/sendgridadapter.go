package emailprovider

import (
	"net/smtp"
)

type sendGridAdapter struct {
	uname  string
	pass   string
	server string
	port   string
}

func NewSendGridAdapter(username string, password string, host string, port string) EmailProvider {
	return &sendGridAdapter{
		uname:  username,
		pass:   password,
		server: host,
		port:   port,
	}
}

func (s *sendGridAdapter) SendHTML(from string, to string, subject string, body string) error {

	auth := smtp.PlainAuth("", s.uname, s.pass, s.server)
	

	msg := []byte(
		"From: " + from + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	err := smtp.SendMail(s.server+":"+s.port, auth, "", []string{to}, msg)

	return err
}

func (s *sendGridAdapter) ProviderName() string {
	return "SendGrid"
}
