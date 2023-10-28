package config

import (
	"errors"
	"os"
)

type Config struct {
	GodvilleGodname string
	GodvilleToken   string

	MastodonServer      string
	MastodonID          string
	MastodonSecret      string
	MastodonAccessToken string
	MastodonEmail       string
	MastodonPassword    string
}

func Load() (*Config, error) {
	godvilleGodname := os.Getenv("GODVILLE_GODNAME")
	if godvilleGodname == "" {
		return nil, errors.New("GODVILLE_GODNAME is not set")
	}
	godvilleToken := os.Getenv("GODVILLE_TOKEN")
	if godvilleToken == "" {
		return nil, errors.New("GODVILLE_TOKEN is not set")
	}

	// get mastodon config
	mastodonServer := os.Getenv("MASTODON_SERVER")
	mastodonID := os.Getenv("MASTODON_ID")
	mastodonSecret := os.Getenv("MASTODON_SECRET")
	mastodonAccessToken := os.Getenv("MASTODON_ACCESS_TOKEN")
	mastodonEmail := os.Getenv("MASTODON_EMAIL")
	mastodonPassword := os.Getenv("MASTODON_PASSWORD")

	if mastodonServer != "" && (mastodonID == "" || mastodonSecret == "" || mastodonEmail == "" || mastodonPassword == "" || mastodonAccessToken == "") {
		return nil, errors.New("MASTODON_ID, MASTODON_SECRET, MASTODON_ACCESS_TOKEN, MASTODON_EMAIL and MASTODON_PASSWORD must be set if MASTODON_SERVER is set")
	}

	// return error if no publishers are configured
	if mastodonServer == "" {
		return nil, errors.New("no publishers are configured")
	}

	return &Config{
		GodvilleGodname: godvilleGodname,
		GodvilleToken:   godvilleToken,

		MastodonServer:      mastodonServer,
		MastodonID:          mastodonID,
		MastodonSecret:      mastodonSecret,
		MastodonAccessToken: mastodonAccessToken,
		MastodonEmail:       mastodonEmail,
		MastodonPassword:    mastodonPassword,
	}, nil
}
