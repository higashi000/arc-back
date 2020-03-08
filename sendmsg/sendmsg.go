package sendmsg

import (
	"context"

	"cloud.google.com/go/firestore"
	arcslack "github.com/higashi000/arc-back/slack"
	arctweet "github.com/higashi000/arc-back/tweet"
	"google.golang.org/api/iterator"
)

type Msg struct {
	Mention []struct {
		SlackRN   string `json:"SlackRN"`
		SlackID   string `json:"SlackID"`
		TwitterID string `json:"TwitterID"`
	} `json:"mention"`
	Text      string `json:"text"`
	Channel   string `json:"channel"`
	Timestamp string `json:"timestamp"`
}

func SendMsg(client *firestore.Client, ctx context.Context) error {
	iter := client.Collection("messages").Documents(ctx)

	var msgs []Msg

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		var tmp Msg
		doc.DataTo(&tmp)

		msgs = append(msgs, tmp)
	}

	for _, e := range msgs {
		var tmp []string
		for _, user := range e.Mention {
			tmp = append(tmp, user.SlackID)
		}
		reactedUser, err := arcslack.CheckReaction(e.Timestamp, e.Channel, tmp)
		if err != nil {
			return err
		}

		var sendTarget []string

		for _, target := range e.Mention {
			flg := false
			for i := 0; i < len(reactedUser); i++ {
				if target.SlackID == reactedUser[i] {
					flg = true
					break
				}
			}

			if !flg {
				sendTarget = append(sendTarget, target.TwitterID)
			}
		}

		var attachmentText string
		if len([]rune(e.Text)) > 20 {
			attachmentText = string([]rune(e.Text)[:20]) + "..."
		} else {
			attachmentText = e.Text
		}

		if len(sendTarget) != 0 {
			err = arctweet.Tweet("Slackに「"+attachmentText+"」というメッセージが届いています。\nリアクションをしてください", sendTarget)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
