package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/higashi000/arc-back/addmsg"
	"github.com/higashi000/arc-back/check"
)

func main() {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: os.Getenv("PROJECT_ID")}
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
	check.Check(r, client, ctx)

	r.Run()
}

func Test(r *gin.Engine, client *firestore.Client, ctx context.Context) {
	r.GET("arc/test/", func(c *gin.Context) {
		c.String(200, "hogehoge")
	})
}
