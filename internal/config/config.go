package config

import (
	"errors"
	"os"
)

type Config struct {
	GodvilleGodname string
	GodvilleToken   string
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
	return &Config{
		GodvilleGodname: godvilleGodname,
		GodvilleToken:   godvilleToken,
	}, nil
}
