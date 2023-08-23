package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/noodlensk/badgy/pkg/config"
	"github.com/noodlensk/badgy/pkg/providers/gmail"
	"github.com/noodlensk/badgy/pkg/providers/slack"
)

func run() error {
	cfg, err := config.NewConfigFromEnv()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	gmailProvider, err := gmail.NewProvider(ctx, cfg.Providers.Gmail.Token, cfg.Providers.Gmail.Credentials)
	if err != nil {
		return fmt.Errorf("gmail: %w", err)
	}

	gmailNotificationsCount, err := gmailProvider.NotificationsCount(ctx)
	if err != nil {
		return fmt.Errorf("get gmail notifications: %w", err)
	}

	slackProvider, err := slack.NewProvider(cfg.Providers.Slack.Token, cfg.Providers.Slack.Cookie)
	if err != nil {
		return fmt.Errorf("slack: %w", err)
	}

	slackNotificationsCount, err := slackProvider.NotificationsCount(ctx)
	if err != nil {
		return fmt.Errorf("get slack notifications: %w", err)
	}

	fmt.Println("Notifications:")
	fmt.Printf("Gmail: %d\n", gmailNotificationsCount)
	fmt.Printf("Slack: %d\n", slackNotificationsCount)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
