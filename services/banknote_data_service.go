package services

import (
	"atm-api.com/helpers"
	"atm-api.com/models"
	"atm-api.com/repositories/withdrawal"
)

//BanknoteDataService handle ATM basic operations
type BanknoteDataService struct {
	WithdrawalRepository withdrawal.BaseRepository
}

//Withdrawal receive value and calculate which and how many notes will return
func (b *BanknoteDataService) Withdrawal(value int) (*[]models.BanknoteData, *error) {
	bankNotesValues := make([]int16, 0)
	availableBankNotesValues := b.WithdrawalRepository.GetAvailableBanknotesValues()
	helpers.Slice{}.SortDesc(availableBankNotesValues)

	i := 0
	for value > 0 {
		for value >= int(availableBankNotesValues[i]) {
			bankNotesValues = append(bankNotesValues, availableBankNotesValues[i])
			value = value - int(availableBankNotesValues[i])
		}

		i = i + 1
	}

	return b.ConvertValuesToBanknotes(bankNotesValues), nil
}

//ConvertValuesToBanknotes receives an slice of values and convert it to slice of BanknoteData
func (b *BanknoteDataService) ConvertValuesToBanknotes(values []int16) *[]models.BanknoteData {
	banknotes := make([]models.BanknoteData, 0)
	for i := 0; i < len(values); i++ {
		for j := 0; j <= len(banknotes); j++ {
			if j == len(banknotes) {
				banknotes = append(banknotes, models.BanknoteData{Value: values[i], Quantity: 1})
				break
			}
			if banknotes[j].Value == values[i] {
				banknotes[j].Quantity = banknotes[j].Quantity + 1
				break
			}

		}
	}

	return &banknotes
}
