package main

import (
	"email_service/emailprovider"
	"email_service/fes"
	"os"
)

func main() {
	sendgrid := emailprovider.NewSendGridAdapter("username", "password", "smtp.sendgrid.net", "587")
	aws := emailprovider.NewAwsAdapter("username", "password", "email-smtp.eu-north-1.amazonaws.com", "587")

	content, _ := os.ReadFile("email.html")

	emailsender := fes.NewFallbackEmailSender([]emailprovider.EmailProvider{
		sendgrid,
		aws,
	})

	err := emailsender.SendEmail("from@example.com", "to@example.com", "ExampleSubject", string(content))
	if err != nil {
		panic(err)
	}
}
