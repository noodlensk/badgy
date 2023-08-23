package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Provider struct {
	// token and cookie are extracted from the slack web app. We have to do that since currently it's not possible to
	// create user tokens, only app tokens and app needs to be installed into workspace.
	token  string
	cookie string
}

func NewProvider(token, cookie string) (*Provider, error) {
	return &Provider{token: token, cookie: cookie}, nil
}

// NotificationsCount returns the number of unread messages in slack.
func (p *Provider) NotificationsCount(ctx context.Context) (int, error) {
	// /api/client.counts is undocumented, but it's used by slack internally.
	req, err := p.makeQuery(ctx, "/api/client.counts", map[string][]string{})
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	r := clientsCountResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return 0, fmt.Errorf("unable to decode response: %w", err)
	}

	if !r.Ok {
		return 0, fmt.Errorf("response error: %s", r.Error)
	}

	return r.ChannelBadges.Channels + r.ChannelBadges.Dms + r.ChannelBadges.ThreadMentions, nil
}

func (p *Provider) makeQuery(ctx context.Context, path string, values map[string][]string) (*http.Request, error) {
	v := url.Values(values)
	v.Set("token", p.token)

	u := url.URL{
		Scheme:   "https",
		Host:     "slack.com",
		Path:     path,
		RawQuery: v.Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "d", Value: p.cookie})

	return req, nil
}

type clientsCountResponseSub struct {
	ID             string `json:"id"`
	LastRead       string `json:"last_read"`
	Latest         string `json:"latest"`
	Updated        string `json:"updated"`
	HistoryInvalid string `json:"history_invalid"`
	MentionCount   int    `json:"mention_count"`
	HasUnreads     bool   `json:"has_unreads"`
}

type clientsCountResponse struct {
	Ok      bool   `json:"ok"`
	Error   string `json:"error"`
	Threads struct {
		HasUnreads   bool `json:"has_unreads"`
		MentionCount int  `json:"mention_count"`
	} `json:"threads"`
	Channels      []clientsCountResponseSub `json:"channels"`
	MPIMS         []clientsCountResponseSub `json:"mpims"`
	IMS           []clientsCountResponseSub `json:"ims"`
	ChannelBadges struct {
		Channels       int `json:"channels"`
		Dms            int `json:"dms"`
		AppDms         int `json:"app_dms"`
		ThreadMentions int `json:"thread_mentions"`
		ThreadUnreads  int `json:"thread_unreads"`
	} `json:"channel_badges"`
}
