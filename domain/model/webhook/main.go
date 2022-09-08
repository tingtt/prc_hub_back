package webhook

type Provider interface {
	Notify(token string, msg string) error
}
