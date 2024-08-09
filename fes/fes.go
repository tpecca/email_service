package fes

import (
	"email_service/emailprovider"
	"errors"
	"log"
)

type fallBackEmailSender struct {
	providers []emailprovider.EmailProvider
}

func NewFallbackEmailSender(emailProviders []emailprovider.EmailProvider) fallBackEmailSender {
	return fallBackEmailSender{
		providers: emailProviders,
	}

}

func (f *fallBackEmailSender) SendEmail(from, to, subject, body string) error {
	for _, provider := range f.providers {
		err := provider.SendHTML(from, to, subject, body)
		if err != nil {
			log.Println("Could not send email with provider:", provider.ProviderName())
			log.Println(err)
			continue
		}
		return nil
	}
	return errors.New("all providers failed to send email")
}
