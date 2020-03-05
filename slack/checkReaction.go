package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

func CheckReaction(ts, channelID string, slackrn []string) ([]string, error) {
	api := Api()

	var reactedUser []string

	history, err := api.GetChannelHistory(channelID, slack.HistoryParameters{"0", "0", 50, false, false})
	if err != nil {
		fmt.Println(channelID)
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

	m := make(map[string]bool)
	var delDuplication []string
	var sendMsgUser []string

	for _, e := range reactedUser {
		if !m[e] {
			m[e] = true
			delDuplication = append(delDuplication, e)
		}
	}

	for _, reacted := range delDuplication {
		flg := false
		var targetuser string
		for _, target := range slackrn {
			targetuser = target
			if reacted == target {
				flg = true
			}
		}

		if !flg {
			sendMsgUser = append(sendMsgUser, targetuser)
		}
	}

	return sendMsgUser, nil
}
