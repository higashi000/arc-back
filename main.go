package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/higashi000/arc-back/addmsg"
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

	Test(r)
	TestFirestore(r, client, ctx)
	addmsg.AddMsg(r, client, ctx)

	r.Run()
}

func Test(r *gin.Engine) {
	r.GET("arc/Test", func(c *gin.Context) {
		c.String(200, "TEST!?!??!?!?!?!??!?!")
	})
}

func TestFirestore(r *gin.Engine, client *firestore.Client, ctx context.Context) {
	r.GET("arc/TestFS", func(c *gin.Context) {
		_, err := client.Collection("messages").Doc(c.Query("Document")).Set(ctx, map[string]interface{}{
			"name":    c.Query("name"),
			"message": c.Query("message"),
		})

		if err != nil {
			log.Fatal(err)
		}
	})
}
