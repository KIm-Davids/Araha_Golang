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

<<<<<<< HEAD
		subscriptionCost = 2000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = float64(int(userBalance.Balance) - balance)
		subscription.Discount = 10 / 100 * 1000
=======
		subscriptionCost = 1000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = int(userBalance.Balance) - balance
>>>>>>> 9ad16d30b17e84e71ab15e2abf0e7a71ce354023
		subscription.Date = time.Now().String()
		return subscription
	}

	if subscription.SubscriptionType == constants.JUMIA {
		var userBalance models.User
		var subscriptionCost float64

<<<<<<< HEAD
		subscriptionCost = 4000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = float64(int(userBalance.Balance) - balance)
		subscription.Discount = 10 / 100 * 1000
=======
		subscriptionCost = 2000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = int(userBalance.Balance) - balance
>>>>>>> 9ad16d30b17e84e71ab15e2abf0e7a71ce354023
		subscription.Date = time.Now().String()
		return subscription
	}

	if subscription.SubscriptionType == constants.GOTV {
		var userBalance models.User
		var subscriptionCost float64

<<<<<<< HEAD
		subscriptionCost = 6000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = float64(int(userBalance.Balance) - balance)
		subscription.Discount = 10 / 100 * 1000
=======
		subscriptionCost = 4000
		userBalance.Balance = 10000
		discountAmount := int((10.0 / 100.0) * subscriptionCost)
		balance := int(subscriptionCost) - discountAmount
		subscription.Amount = int(userBalance.Balance) - balance
>>>>>>> 9ad16d30b17e84e71ab15e2abf0e7a71ce354023
		subscription.Date = time.Now().String()
		return subscription
	}
	return noSubscription
}
<<<<<<< HEAD
=======

func subscriptionNotAvailableErrorMessage() exceptions.MyException {
	err := exceptions.MyException{Data: http.StatusExpectationFailed, Object: "This service is currently not available\nPlease try again later"}
	return err
}
>>>>>>> 9ad16d30b17e84e71ab15e2abf0e7a71ce354023
