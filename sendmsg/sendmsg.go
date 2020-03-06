package sendmsg

import (
	"context"

	"cloud.google.com/go/firestore"
	arcslack "github.com/higashi000/arc-back/slack"
	arctwitter "github.com/higashi000/arc-back/tweet"
	"google.golang.org/api/iterator"
)

type Msg struct {
	Mention []struct {
		SlackRN   string `json:"SlackRN"`
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
		var slackrn []string
		var twitterid []string

		for _, user := range e.Mention {
			slackrn = append(slackrn, user.SlackRN)
			twitterid = append(twitterid, user.SlackRN)
		}

		notReaction, err := arcslack.CheckReaction(e.Timestamp, "CEVCQUGAJ", slackrn)
		if err != nil {
			return err
		}

		var notReactionTwitter []string

		for _, user := range e.Mention {
			for _, notreaction := range notReaction {
				if notreaction == user.SlackRN {
					notReactionTwitter = append(notReactionTwitter, user.TwitterID)
					break
				}
			}
		}

		err = arctwitter.Tweet("slackにメッセージが届いています", notReactionTwitter)
		if err != nil {
			return err
		}
	}

	return nil
}
