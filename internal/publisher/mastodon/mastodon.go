package mastodon

import (
	"context"

	"github.com/mattn/go-mastodon"
)

type Config struct {
	MastodonServer      string
	MastodonID          string
	MastodonSecret      string
	MastodonAccessToken string
	MastodonEmail       string
	MastodonPassword    string
}

func New(config *Config) (*Mastodon, error) {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       config.MastodonServer,
		ClientID:     config.MastodonID,
		ClientSecret: config.MastodonSecret,
		AccessToken:  config.MastodonAccessToken,
	})
	err := c.Authenticate(context.Background(), config.MastodonEmail, config.MastodonPassword)
	if err != nil {
		return nil, err
	}
	return &Mastodon{
		client: c,
	}, nil
}

type Mastodon struct {
	client *mastodon.Client
}

func (m *Mastodon) SendMessage(string) error {
	return nil
}

func (m *Mastodon) UpdateStatus(string) error {
	return nil
}
