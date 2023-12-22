package main

import (
	"context"
	"flag"
	"log"

	"github.com/sattellite/godville-publisher/internal/config"
	"github.com/sattellite/godville-publisher/internal/godville"
	"github.com/sattellite/godville-publisher/internal/publisher/mastodon"
)

func main() {
	var isSendMessage bool
	var isUpdateProfile bool
	flag.BoolVar(&isSendMessage, "m", false, "send last dairy message")
	flag.BoolVar(&isUpdateProfile, "p", false, "update profile from last game info")
	flag.Parse()

	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		log.Fatalln(cfgErr)
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
		log.Fatalln(pErr)
	}

	game := godville.New(cfg.GodvilleGodname, cfg.GodvilleToken)
	info, err := game.Info(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	if isSendMessage {
		msgErr := publisher.SendMessageWithoutDuplicate(ctx, info.DiaryLast)
		if msgErr != nil {
			log.Printf("failed send message: %s\n", msgErr.Error())
		}
	}

	if isUpdateProfile {
		profErr := publisher.UpdateProfile(ctx, info)
		if profErr != nil {
			log.Printf("failed update profile: %s", profErr.Error())
		}
	}
}
