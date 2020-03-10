package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/higashi000/arc-back/addmsg"
	"github.com/higashi000/arc-back/sendmsg"
)

func main() {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "higashi-arc"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	addmsg.AddMsg(r, client, ctx)
	Test(r, client, ctx)

	r.Run()
}

func Test(r *gin.Engine, client *firestore.Client, ctx context.Context) {
	r.GET("arc/test/", func(c *gin.Context) {
		err := sendmsg.SendMsg(client, ctx)
		if err != nil {
			log.Fatal(err)
		}

		sendmsg.SendMsg(client, ctx)

		c.String(200, "hogehoge")
	})
}
