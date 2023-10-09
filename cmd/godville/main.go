package main

import (
	"context"

	"github.com/sattellite/godville-publisher/internal/config"
	"github.com/sattellite/godville-publisher/internal/godville"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	cfg, cfgErr := config.Load()
	if cfgErr != nil {
		panic(cfgErr)
	}

	ctx := context.Background()

	game := godville.New(cfg.GodvilleGodname, cfg.GodvilleToken)
	info, err := game.Info(ctx)
	if err != nil {
		panic(err)
	}

	spew.Dump(info)
}
