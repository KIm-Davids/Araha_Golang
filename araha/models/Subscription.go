package models

import (
	"araha/constants"
)

type Subscription struct {
	ID               string                     `json:"id"`
	SubscriptionType constants.SubscriptionType `json:"SubscriptionType"`
	Description      string                     `json:"description"`
	Amount           float64
	Discount         float64
	Date             string
	UserId           int
}
