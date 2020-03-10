package check

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/higashi000/arc-back/sendmsg"
)

func Check(r *gin.Engine, client *firestore.Client, ctx context.Context) {
	r.GET(os.Getenv("UPDATE_TIME"), func(c *gin.Context) {
		err := sendmsg.SendMsg(client, ctx)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, `{"status": "ok"}`)
	})
}
