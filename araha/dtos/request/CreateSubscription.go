package dtos

import "araha/araha/constants"

type CardDetails struct {
	CardNumber string
	CardCvv    string
	CardPin    string
}

type CreateSubscription struct {
	SubscriptionType constants.SubscriptionType `json:"subscriptionType"`
	PaymentCard      CardDetails                `json:"paymentCard"`
}
