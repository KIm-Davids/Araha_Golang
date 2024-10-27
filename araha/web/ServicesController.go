package web

import (
	"araha/models"
	"araha/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSubscriptionController() gin.HandlerFunc {
	var createSubscription services.NewSubscriptionServices

	//if http.MethodPost {
	//
	//}

	return func(c *gin.Context) {
		var Subscription models.Subscription

		if err := c.ShouldBindJSON(&Subscription); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error creating subscription": err.Error()})
			return
		}

		result, err := createSubscription.CreateSubscription(Subscription)

		if result == 409 {
			c.JSON(http.StatusConflict, gin.H{"message": "Already on this subscription"})
		} else {
			c.JSON(http.StatusCreated, gin.H{"data": "Created Subscription successfully !!!",
				"Object": result})
		}

		if err != nil {
			fmt.Printf("Couldnt create a subscription %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	}

}

func UpdateSubscriptionController() gin.HandlerFunc {
	var updateSubscription services.NewSubscriptionServices
	return func(c *gin.Context) {
		var subscription models.Subscription
		if err := c.ShouldBindJSON(&subscription); err != nil {
			c.JSON(http.StatusNotModified, gin.H{"Error updating subscription": err})
		}
		result, err := updateSubscription.UpdateSubscription(subscription)
		if err != nil {
			c.JSON(http.StatusNotModified, gin.H{"Error": err})
		}
		if result == http.StatusNotModified {
			c.JSON(http.StatusNotModified, gin.H{"message": "Could not update the subscription\nPlease try again later"})
		}
		c.JSON(http.StatusOK, gin.H{"data": "Updated Successfully",
			"Object": result})
	}
}

func DeleteSubscriptionController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var subscription models.Subscription
		var deleteSubscription services.NewSubscriptionServices
		if err := c.ShouldBindJSON(&subscription); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "An error occurred"})
		}

		response, err := deleteSubscription.DeleteSubscription(subscription)

		if response == http.StatusOK {
			c.JSON(response, gin.H{"message": "Subscription deleted successfully"})
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error Could not delete subscription"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Error Could not delete subscription"})

	}
}

func GetAllSubscriptionController() gin.HandlerFunc {
	return func(c *gin.Context) {
		var subscription models.Subscription
		var getAllSubServices services.NewSubscriptionServices

		if err := c.ShouldBindJSON(&subscription); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "An error occurred"})
		}

		result, err := getAllSubServices.GetAllSubscription()

		if result == http.StatusFound {
			c.JSON(http.StatusFound, gin.H{"Fetched Data from the database": result})
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	}
}
