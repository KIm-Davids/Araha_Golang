package services

import (
	"araha/exceptions"
	"araha/mapper"
	"araha/models"
	"araha/repository"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CreateSubscriptionServices interface {
	CreateSubscription(subscription models.Subscription) (int, error)
}

type NewSubscriptionServices struct{}

func (nss *NewSubscriptionServices) CreateSubscription(subscription models.Subscription) (int, error) {
	var invalidDetailsException exceptions.MyException
	var foundSubType models.Subscription

	db, err := repository.SubscriptionRepo()
	db.Find(&foundSubType)

	if foundSubType.SubscriptionType == subscription.SubscriptionType {
		return http.StatusConflict, nil
	}
	//var userBalance models.User
	// userBalance.Balance > 0

	if subscription.SubscriptionType != " " {
		subscription = mapper.FindSubscriptionTypes(subscription)
		db, err := repository.SubscriptionRepo()
		db.Save(subscription)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusCreated, nil
	}

	if err != nil {
		log.Fatalf("An error occurred in creating a new subscription: %v ", err)
	}
	return http.StatusBadRequest, &invalidDetailsException
}

type UpdateSubscriptionServices struct{}

func (nss *NewSubscriptionServices) UpdateSubscription(updateSubscription models.Subscription) (int, error) {
	var subscription models.Subscription

	db, err := repository.SubscriptionRepo()
	result := db.Where("id = ?", updateSubscription.ID).First(&subscription)
	fmt.Print(result)

	if updateSubscription.SubscriptionType != " " && subscription.SubscriptionType == updateSubscription.SubscriptionType {
		foundSubscription := mapper.FindSubscriptionTypes(updateSubscription)
		updateSubscription.Amount = foundSubscription.Amount + subscription.Amount
		subscription.Date = time.Now().String()

		db, err = repository.SubscriptionRepo()
		db.Save(updateSubscription)

		if err != nil {
			log.Fatalf("Couldn't update the user's subscription %v", err)
			return http.StatusAccepted, nil
		}
	}
	return http.StatusNotModified, nil

}

type DeleteSubscriptionServices struct{}

func (nss *NewSubscriptionServices) DeleteSubscription(subscription models.Subscription) (int, error) {

	db, err := repository.SubscriptionRepo()
	db.Delete(subscription)

	if err != nil {
		return http.StatusNotModified, err
	}
	return http.StatusOK, nil

}

func (nss *NewSubscriptionServices) GetAllSubscription() interface{} {
	var unableToGetAllValuesException exceptions.MyException
	var allSubscription models.Subscription

	db, err := repository.SubscriptionRepo()
	foundSub := db.Find(&allSubscription)

	if err != nil {
		log.Fatalf("Could'nt retrieve all the values from the database %v", unableToGetAllValuesException)
	}

	return foundSub

}
