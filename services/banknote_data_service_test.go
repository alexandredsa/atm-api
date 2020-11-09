package services_test

import (
	"reflect"
	"testing"

	"atm-api.com/exceptions"
	mock_withdrawal "atm-api.com/mocks"
	"atm-api.com/models"
	"atm-api.com/repositories/withdrawal"
	"atm-api.com/services"
	"github.com/golang/mock/gomock"
)

func TestWithdrawalWithSuccess(t *testing.T) {
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

	banknotes, _ := banknoteDataService.Withdrawal(172)

	if !reflect.DeepEqual(&expectedBanknotes, banknotes) {
		t.Errorf("Error on Withdrawal method; expected %v, got %v", expectedBanknotes, banknotes)
	}
}

func TestWithdrawalWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mock_withdrawal.NewMockBaseRepository(ctrl)
	mockRepository.
		EXPECT().
		GetAvailableBanknotesValues().
		Return([]int{50, 2})

	banknoteDataService := services.BanknoteDataService{WithdrawalRepository: mockRepository}

	_, err := banknoteDataService.Withdrawal(53)

	if err == nil {
		t.Errorf("Expected Error")
	}

	_, ok := err.(*exceptions.UnsupportedValueError)

	if !ok {
		t.Errorf("Expected: UnsupportedValueError, got %v", err)
	}
}
