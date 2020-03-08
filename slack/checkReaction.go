package slack

import (
	"fmt"
	"log"

	"github.com/slack-go/slack"
)

func CheckReaction(ts, channelID string, slackrn []string) ([]string, error) {
	api := Api()

	var reactedUser []string

	history, err := api.GetChannelHistory(channelID, slack.HistoryParameters{"0", "0", 50, false, false})
	if err != nil {
		fmt.Println(channelID)
		return []string{}, err
	}

	var targetMsg slack.Msg

	for _, e := range history.Messages {
		if e.Msg.Timestamp == ts {
			targetMsg = e.Msg
			break
		}
	}

	if targetMsg.Text == "" {
		log.Println("error: can not found message")
		return []string{}, nil
	}

	for _, e := range targetMsg.Reactions {
		reactedUser = append(reactedUser, e.Users...)
	}

	return reactedUser, nil
}
