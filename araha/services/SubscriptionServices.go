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

func (nss *NewSubscriptionServices) UpdateSubscription(updateSubscription models.Subscription) (int, error) {
	var subscription models.Subscription

	db, err := repository.SubscriptionRepo()
	result := db.Where("id = ?", updateSubscription.ID).First(&subscription)
	fmt.Print(result)

	if err != nil {
		log.Fatalf("An error occured in fetching from database")
	}

	if updateSubscription.SubscriptionType != " " && subscription.SubscriptionType == updateSubscription.SubscriptionType {
		foundSubscription := mapper.FindSubscriptionTypes(updateSubscription)
		updateSubscription.Amount = foundSubscription.Amount + subscription.Amount
		updateSubscription.Date = time.Now().String()

		db, err = repository.SubscriptionRepo()
		db.Save(updateSubscription)
		return http.StatusAccepted, nil
	}

	return http.StatusNotModified, nil

}

type DeleteSubscriptionServices struct{}

func (nss *NewSubscriptionServices) DeleteSubscription(subscription models.Subscription) (int, error) {

	var sub models.Subscription

	db, err := repository.SubscriptionRepo()
	foundSub := db.Where("subscription_type = ?", subscription.SubscriptionType).First(&sub)
	if foundSub.Error != nil {
		log.Fatalf("Couldn't retrieve values from db Error: %v", foundSub.Error)
	}
	result := db.Delete(&sub)
	if result.RowsAffected >= 1 {
		return http.StatusOK, nil
	}
	if err != nil {
		return http.StatusNotModified, err
	}
	return http.StatusInternalServerError, err
}

func (nss *NewSubscriptionServices) GetAllSubscription() (interface{}, error) {
	var allSubscription models.Subscription

	db, err := repository.SubscriptionRepo()
	foundSub := db.Find(&allSubscription)
	if foundSub.Error != nil {
		return http.StatusInternalServerError, err
	}

	//if err != nil {
	//	log.Fatalf("Could'nt retrieve all the values from the database %v", unableToGetAllValuesException)
	//}

	return allSubscription, nil

}
