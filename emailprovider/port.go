package emailprovider

type EmailProvider interface {
	SendHTML(from, to, subject, body string) error
	ProviderName() string
}
