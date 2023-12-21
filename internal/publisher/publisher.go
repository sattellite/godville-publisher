package publisher

import "context"

type Publisher interface {
	SendMessage(context.Context, string) error
	SendMessageWithoutDuplicate(context.Context, string) error
	UpdateStatus(context.Context, string) error
}
