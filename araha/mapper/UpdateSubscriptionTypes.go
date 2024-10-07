package mapper

import "araha/araha/models"

func UpdateSubscription(subscription models.Subscription) float64 {
	subscriptionFound := FindSubscriptionTypes(subscription)
	updatedSubscription := subscriptionFound.Amount + subscription.Amount
	return updatedSubscription
}
