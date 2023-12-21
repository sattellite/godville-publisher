package mastodon

import (
	"context"
	"fmt"

	"github.com/sattellite/godville-publisher/internal/godville"

	strip "github.com/grokify/html-strip-tags-go"
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
		return nil, fmt.Errorf("failed authenticate: %+w", err)
	}

	acc, accErr := c.GetAccountCurrentUser(context.Background())
	if accErr != nil {
		return nil, fmt.Errorf("failed get account info: %+w", accErr)
	}
	return &Mastodon{
		client: c,
		accID:  acc.ID,
	}, nil
}

type Mastodon struct {
	client *mastodon.Client
	accID  mastodon.ID
}

func (m *Mastodon) SendMessage(ctx context.Context, message string) error {
	_, err := m.client.PostStatus(ctx, &mastodon.Toot{
		Status: message,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Mastodon) SendMessageWithoutDuplicate(ctx context.Context, message string) error {
	c, err := m.client.GetAccountStatuses(ctx, m.accID, &mastodon.Pagination{
		Limit: 1,
	})
	if err != nil {
		return err
	}
	// if no statuses or last status message does not equal with new message
	if len(c) != 0 || strip.StripTags(c[0].Content) != message {
		return m.SendMessage(ctx, message)
	}

	return fmt.Errorf("have same message already")
}

func (m *Mastodon) UpdateProfile(ctx context.Context, info *godville.Info) error {
	loc := info.TownName
	if loc == "" {
		loc = info.Location
	}
	fields := []mastodon.Field{
		{Name: "Здоровье", Value: fmt.Sprintf("%d/%d", info.Health, info.MaxHealth)},
		{Name: "Квест", Value: info.Quest},
		{Name: "Местоположение", Value: loc},
		{Name: "Характер", Value: info.Alignment},
	}

	note := fmt.Sprintf(`Меня постоянно шпуняет мой главный бог %s. В печёнках уже сидит.
Еще и сюда обязал отправлять записи из дневника.
Его профиль в Годвиле, я туда чаще пишу по привычке: https://godville.net/gods/%s`, info.Godname, info.Godname)

	_, err := m.client.AccountUpdate(ctx, &mastodon.Profile{
		Note:   &note,
		Fields: &fields,
	})
	return err
}
