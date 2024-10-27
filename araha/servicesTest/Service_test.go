package ServiceTest

import (
	"araha/models"
	"araha/services"
	"gotest.tools/v3/assert"
	"log"
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
	subscription.ID = "1"

	result, err := subServices.CreateSubscription(subscription)
	expected, err := http.StatusCreated, err

	if result != expected {
		t.Errorf("Wanted %v, got %v", expected, result)
	}

	assert.Equal(t, result, expected)
}

func TestToUpdateSubscription(t *testing.T) {
	var subscription models.Subscription
	var subServices services.NewSubscriptionServices

	subscription.Amount = 8000
	subscription.Date = time.Now().String()
	subscription.SubscriptionType = "Netflix"
	subscription.Description = "Changed sub to gotv"
	subscription.ID = "1"

	result, err := subServices.UpdateSubscription(subscription)
	expected, err := http.StatusAccepted, err

	if result != expected {
		t.Errorf("Wanted %v Got %v", expected, result)
	}

	assert.Equal(t, result, expected)
}

func TestDeleteSubscription(t *testing.T) {
	var subscription models.Subscription
	var subService services.NewSubscriptionServices

	subscription.SubscriptionType = "Netflix"
	result, err := subService.DeleteSubscription(subscription)
	expected, err := http.StatusOK, err

	if result != expected {
		t.Errorf("Wanted %v, Got %v", result, expected)
	}
	assert.Equal(t, result, expected)
}

func TestThatWeCanRetrieveValuesFromTheDB(t *testing.T) {
	var subService services.NewSubscriptionServices

	result, err := subService.GetAllSubscription()
	expected := http.StatusFound

	if result != expected {
		t.Errorf("Wanted %v, Got %v", expected, result)
	}
	assert.Equal(t, result, expected)

	if err != nil {
		log.Fatalln("An error occurred in test")
	}

}
