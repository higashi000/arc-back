package slack

import (
	"github.com/slack-go/slack"
)

func PostMsg(target []slack.User, channelID string) error {
	api := Api()
	text := ""

	for _, e := range target {
		text += "<@" + e.ID + "> "
	}

	_, _, err := api.PostMessage(channelID,
		slack.MsgOptionText(text, false),
		slack.MsgOptionAttachments(slack.Attachment{}))

	if err != nil {
		return err
	}

	return nil
}
