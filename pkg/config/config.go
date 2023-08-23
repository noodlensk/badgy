package config

import (
	"fmt"
	"os"
)

type Config struct {
	Providers struct {
		Gmail struct {
			Token       []byte
			Credentials []byte
		}
		Slack struct {
			Token  string
			Cookie string
		}
	}
}

func NewConfigFromEnv() (*Config, error) {
	expectedEnvVars := []string{
		"BADGY_GMAIL_TOKEN",
		"BADGY_GMAIL_CREDENTIALS",
		"BADGY_SLACK_TOKEN",
		"BADGY_SLACK_COOKIE",
	}

	var errors []string

	for _, envVar := range expectedEnvVars {
		if os.Getenv(envVar) == "" {
			errors = append(errors, fmt.Sprintf("missing environment variable %s", envVar))
		}
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("config: %s", errors)
	}

	cfg := &Config{}

	cfg.Providers.Gmail.Token = []byte(os.Getenv("BADGY_GMAIL_TOKEN"))
	cfg.Providers.Gmail.Credentials = []byte(os.Getenv("BADGY_GMAIL_CREDENTIALS"))
	cfg.Providers.Slack.Token = os.Getenv("BADGY_SLACK_TOKEN")
	cfg.Providers.Slack.Cookie = os.Getenv("BADGY_SLACK_COOKIE")

	return cfg, nil
}
