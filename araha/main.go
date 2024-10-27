package main

import (
	"araha/web"
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
	r.DELETE("/delete-sub", web.DeleteSubscriptionController())
	r.GET("/get-all-sub", web.GetAllSubscriptionController())
	r.Run()

}

// go mod init your_project_name (Use this to initialize your go project)
// go get -u github.com/gin-gonic/gin (to install gin)
// go get -u gorm.io/gorm (to install gorm)
//
