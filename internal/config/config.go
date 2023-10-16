package config

import (
	"errors"
	"os"
)

type Config struct {
	GodvilleGodname string
	GodvilleToken   string

	MastodonServer string
	MastodonID     string
	MastodonSecret string
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

	if mastodonServer != "" && (mastodonID == "" || mastodonSecret == "") {
		return nil, errors.New("MASTODON_ID and MASTODON_SECRET must be set if MASTODON_SERVER is set")
	}

	return &Config{
		GodvilleGodname: godvilleGodname,
		GodvilleToken:   godvilleToken,

		MastodonServer: mastodonServer,
		MastodonID:     mastodonID,
		MastodonSecret: mastodonSecret,
	}, nil
}
