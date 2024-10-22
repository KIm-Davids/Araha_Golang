package ServiceTest

import (
	"araha/models"
	"araha/services"
	"gotest.tools/v3/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateSubscription(t *testing.T) {
	var subServices services.NewSubscriptionServices
	var subscription models.Subscription

	subscription.Amount = 4000
	subscription.SubscriptionType = "Netflix"
	subscription.Date = time.Now().String()
	subscription.Description = "Pay for Netflix"

	result, err := subServices.CreateSubscription(subscription)
	expected, err := http.StatusBadRequest, err

	if result != expected {
		t.Errorf("Wanted %v, got %v", expected, result)
	}

	assert.Equal(t, result, expected)
}
