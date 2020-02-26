package slack

import (
	"github.com/slack-go/slack"
)

func CheckReaction(ts string, channelID string) ([]string, error) {
	api := Api()

	var reactedUser []string

	history, err := api.GetChannelHistory(channelID, slack.HistoryParameters{"0", "0", 100, false, false})
	if err != nil {
		return reactedUser, err
	}

	var targetMsg slack.Message

	for _, e := range history.Messages {
		if e.Timestamp == ts {
			targetMsg = e
			break
		}
	}

	for _, e := range targetMsg.Reactions {
		reactedUser = append(reactedUser, e.Users...)
	}

	return reactedUser, nil
}
