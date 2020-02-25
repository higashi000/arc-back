package slack

import (
	"github.com/slack-go/slack"
)

func PostMsg(target []slack.User, channelID string, sendText string) error {
	api := Api()
	text := ""

	for _, e := range target {
		text += "<@" + e.ID + "> "
	}

	text += sendText

	_, _, err := api.PostMessage(channelID,
		slack.MsgOptionAsUser(true),
		slack.MsgOptionText(text, false),
		slack.MsgOptionAttachments(slack.Attachment{}))

	if err != nil {
		return err
	}

	return nil
}
