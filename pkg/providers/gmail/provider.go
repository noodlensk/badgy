package gmail

import (
	"context"
	"encoding/json"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type Provider struct{ svc *gmail.Service }

func NewProvider(ctx context.Context, token, credentials []byte) (*Provider, error) {
	config, err := google.ConfigFromJSON(credentials, gmail.GmailReadonlyScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %w", err)
	}

	tok := &oauth2.Token{}
	if err = json.Unmarshal(token, &tok); err != nil {
		return nil, fmt.Errorf("unable to decode token file: %w", err)
	}

	svc, err := gmail.NewService(ctx, option.WithHTTPClient(config.Client(ctx, tok)))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Gmail client: %w", err)
	}

	return &Provider{svc: svc}, nil
}

// NotificationsCount returns the number of unread emails in the inbox.
func (p *Provider) NotificationsCount(ctx context.Context) (int, error) {
	r, err := p.svc.Users.Messages.List("me").Q("in:inbox and is:unread").Context(ctx).Do()
	if err != nil {
		return 0, fmt.Errorf("unable to retrieve labels: %w", err)
	}

	return len(r.Messages), nil
}
