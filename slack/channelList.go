package slack

import (
	"github.com/slack-go/slack"
)

func GetChannelList() ([]slack.Channel, error) {
	api := Api()

	channels, err := api.GetChannels(false)
	if err != nil {
		return channels, err
	}

	return channels, nil
}
