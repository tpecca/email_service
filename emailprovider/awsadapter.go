package emailprovider

import (
	"net/smtp"
)

type awsAdapter struct {
	uname  string
	pass   string
	server string
	port   string
}

func NewAwsAdapter(username, password, host, port string) EmailProvider {
	return &awsAdapter{
		uname:  username,
		pass:   password,
		server: host,
		port:   port,
	}
}

func (s *awsAdapter) SendHTML(from string, to string, subject string, body string) error {

	auth := smtp.PlainAuth("", s.uname, s.pass, s.server)

	msg := []byte(
		"Subject: " + subject + "\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
			"\r\n" +
			body + "\r\n",
	)

	err := smtp.SendMail(s.server+":"+s.port, auth, from, []string{to}, msg)

	return err
}

func (s *awsAdapter) ProviderName() string {
	return "aws"
}
