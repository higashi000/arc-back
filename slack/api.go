package slack

import (
	"os"

	"github.com/slack-go/slack"
)

func Api() *slack.Client {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	return api
}
