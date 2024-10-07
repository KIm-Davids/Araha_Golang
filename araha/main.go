package main

import (
	"araha/araha/web"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create-sub", web.CreateSubscriptionController())
	r.PATCH("/update-sub", web.UpdateSubscriptionController())
	r.Run()

}
