package services_test

import (
	"reflect"
	"testing"

	"atm-api.com/models"
	"atm-api.com/repositories/withdrawal"
	"atm-api.com/services"
)

func TestWithdrawal(t *testing.T) {
	repository := withdrawal.DefaultRepository{}
	banknoteDataService := services.BanknoteDataService{WithdrawalRepository: repository}
	expectedBanknotes := []models.BanknoteData{
		{
			Value:    50,
			Quantity: 3,
		},
		{
			Value:    10,
			Quantity: 2,
		},
		{
			Value:    1,
			Quantity: 2,
		},
	}

	banknotes := banknoteDataService.Withdrawal(172)

	if !reflect.DeepEqual(expectedBanknotes, banknotes) {
		t.Errorf("Error on Withdrawal method; expected %v, got %v", expectedBanknotes, banknotes)
	}
}
