package web

import (
	"araha/models"
	"araha/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login() interface{} {

	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		result := services.LoginServices()

	}
}
