package addmsg

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	arcslack "github.com/higashi000/arc-back/slack"
)

type Msg struct {
	Users   []User `json:"mention"`
	Pass    string `json:"pass"`
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

type User struct {
	SlackRN   string `json:"SlackRN"`
	SlackID   string `json:"SlackID"`
	TwitterID string `json:"TwitterID"`
}

func AddMsg(r *gin.Engine, client *firestore.Client, ctx context.Context) {
	r.POST("arc/AddMsg", func(c *gin.Context) {
		var msg Msg
		c.BindJSON(&msg)

		channelID := arcslack.GetChannelID(msg.Channel)

		if msg.Pass != os.Getenv("PASS") {
			c.JSON(http.StatusOK, `{"status":"false"}`)
			return
		}

		var slackRN []string

		for _, e := range msg.Users {
			slackRN = append(slackRN, e.SlackRN)
		}
		users := arcslack.UserList(slackRN)

		for i, e := range msg.Users {
			for _, user := range users {
				if e.SlackRN == user.Profile.RealName {
					msg.Users[i].SlackID = user.ID
				}
			}
		}

		channelID, ts, err := arcslack.PostMsg(users, channelID, msg.Text)
		if err != nil {
			log.Fatal(err)
		}

		_, _, err = client.Collection("messages").Add(ctx, map[string]interface{}{
			"channel":   channelID,
			"text":      msg.Text,
			"timestamp": ts,
			"mention":   msg.Users,
		})
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, `{"status":"Ok"}`)
	})
}
