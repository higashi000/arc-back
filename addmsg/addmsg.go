package addmsg

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	arcslack "github.com/higashi000/arc-back/slack"
)

type Msg struct {
	Users   []User `json:"mention"`
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

		var slackRN []string

		for _, e := range msg.Users {
			slackRN = append(slackRN, e.SlackRN)
		}
		users := arcslack.UserList(slackRN)

		channelID, ts, err := arcslack.PostMsg(users, msg.Channel, msg.Text)
		if err != nil {
			log.Fatal(err)
		}

		_, _, err = client.Collection("messages").Add(ctx, map[string]interface{}{
			"channel":   channelID,
			"timestamp": ts,
			"mention":   msg.Users,
		})
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, msg)
	})
}
