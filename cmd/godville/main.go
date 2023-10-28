package main

import (
	"context"

	"github.com/sattellite/godville-publisher/internal/config"
	"github.com/sattellite/godville-publisher/internal/godville"
	"github.com/sattellite/godville-publisher/internal/publisher/mastodon"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		panic(cfgErr)
	}

	ctx := context.Background()

	publisher, pErr := mastodon.New(&mastodon.Config{
		MastodonServer:      cfg.MastodonServer,
		MastodonID:          cfg.MastodonID,
		MastodonSecret:      cfg.MastodonSecret,
		MastodonAccessToken: cfg.MastodonAccessToken,
		MastodonEmail:       cfg.MastodonEmail,
		MastodonPassword:    cfg.MastodonPassword,
	})
	if pErr != nil {
		panic(pErr)
	}

	game := godville.New(cfg.GodvilleGodname, cfg.GodvilleToken)
	info, err := game.Info(ctx)
	if err != nil {
		panic(err)
	}

	spew.Dump(publisher)
	spew.Dump(info)
}
