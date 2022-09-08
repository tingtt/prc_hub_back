package webhook

import (
	"errors"
	"prc_hub_back/domain/model/webhook"
)

// Singleton fields
var (
	initialized      = false
	frontUrl         string
	webHookProviders []Provider
)

type Provider struct {
	Provider webhook.Provider
	Token    string
}

// Errors
var (
	ErrProviderNotInitialized = errors.New("webhook provider not initialized")
)

func InitApplication(frontEndUrl string, p ...Provider) {
	initialized = true
	frontUrl = frontEndUrl
	webHookProviders = append(webHookProviders, p...)
}

func NotifyToAllProviders(msg string) error {
	if !initialized {
		return ErrProviderNotInitialized
	}

	for _, p := range webHookProviders {
		err := p.Provider.Notify(p.Token, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
