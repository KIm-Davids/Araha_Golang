package services

import "araha/models"

type SubscriptionServices interface {
	CreateSubscription(subscription models.Subscription) (int, error)
	UpdateSubscription(updateSubscription models.Subscription) (int, error)
	DeleteSubscription(subscription models.Subscription) (int, error)
	GetAllSubscription() interface{}
}
