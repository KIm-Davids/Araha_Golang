package mapper

import (
	"araha/araha/constants"
	"araha/araha/exceptions"
	"araha/araha/models"
	"net/http"
	"time"
)

func FindSubscriptionTypes(subscription models.Subscription) models.Subscription {
	var noSubscription models.Subscription

	if subscription.SubscriptionType == constants.NETFLIX {
		return netflixMapper(subscription)
	}

	if subscription.SubscriptionType == constants.JUMIA {
		return jumiaMapper(subscription)
	}

	if subscription.SubscriptionType == constants.GOTV {
		return gotvMapper(subscription)
	}
	return noSubscription
}

func jumiaMapper(subscription models.Subscription) models.Subscription {
	subscription.Amount = 1000
	subscription.Discount = 10 / 100 * 1000
	subscription.Date = time.Now().String()
	return subscription
}

func netflixMapper(subscription models.Subscription) models.Subscription {
	subscription.Amount = 2000
	subscription.Discount = 10 / 100 * 2000
	subscription.Date = time.Now().String()
	return subscription
}

func gotvMapper(subscription models.Subscription) models.Subscription {
	subscription.Amount = 8000
	subscription.Discount = 10 / 100 * 8000
	subscription.Date = time.Now().String()
	return subscription
}

func subscriptionNotAvailableErrorMessage() exceptions.MyException {
	err := exceptions.MyException{Data: http.StatusExpectationFailed, Object: "This service is currently not available\nPlease try again later"}
	return err
}
