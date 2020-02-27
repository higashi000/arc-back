package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	Test(r)

	r.Run()
}

func Test(r *gin.Engine) {
	r.GET("arc/Test", func(c *gin.Context) {
		c.String(200, "TEST!?!??!?!?!?!??!?!")
	})
}
