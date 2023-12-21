package publisher

import (
	"context"

	"github.com/sattellite/godville-publisher/internal/godville"
)

type Publisher interface {
	SendMessage(context.Context, string) error
	SendMessageWithoutDuplicate(context.Context, string) error
	UpdateProfile(context.Context, *godville.Info) error
}
