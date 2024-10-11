package mapper

import (
	"araha/constants"
	"araha/models"
	"time"
)

func FindSubscriptionTypes(subscription models.Subscription) models.Subscription {
	var noSubscription models.Subscription

	if subscription.SubscriptionType == constants.NETFLIX {
		var userBalance models.User
		var subscriptionCost float64

		subscriptionCost = 2000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = float64(int(userBalance.Balance) - balance)
		subscription.Discount = 10 / 100 * 1000
		subscription.Date = time.Now().String()
		return subscription
	}

	if subscription.SubscriptionType == constants.JUMIA {
		var userBalance models.User
		var subscriptionCost float64

		subscriptionCost = 4000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = float64(int(userBalance.Balance) - balance)
		subscription.Discount = 10 / 100 * 1000
		subscription.Date = time.Now().String()
		return subscription
	}

	if subscription.SubscriptionType == constants.GOTV {
		var userBalance models.User
		var subscriptionCost float64

		subscriptionCost = 6000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = float64(int(userBalance.Balance) - balance)
		subscription.Discount = 10 / 100 * 1000
		subscription.Date = time.Now().String()
		return subscription
	}
	return noSubscription
}
