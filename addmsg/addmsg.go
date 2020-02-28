package addmsg

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type Msg struct {
	Mention []struct {
		SlackRN   string `json:"slackrn"`
		TwitterID string `json:"twitterid"`
	} `json:"mention"`

	Text string `json:"text"`
}

func AddMsg(r *gin.Engine, client *firestore.Client, ctx context.Context) {
	r.POST("arc/AddMsg", func(c *gin.Context) {
		var msg Msg
		c.BindJSON(msg)

		_, _, err := client.Collection("messages").Add(ctx, map[string]interface{}{
			"text":    msg.Text,
			"mention": msg.Mention,
		})
		if err != nil {
			log.Fatal(err)
		}
	})
}
