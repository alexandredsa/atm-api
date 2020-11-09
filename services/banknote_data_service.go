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

//CreateBanknoteService returns a new instance of BanknoteDataService
func CreateBanknoteService(withdrawalRepository withdrawal.BaseRepository) BanknoteDataService {
	return BanknoteDataService{WithdrawalRepository: withdrawalRepository}
}

//Withdrawal receive value and calculate which and how many notes will return
func (b *BanknoteDataService) Withdrawal(value int) (*[]models.BanknoteData, *error) {
	banknoteDatas := make([]models.BanknoteData, 0)
	availableBankNotesValues := b.WithdrawalRepository.GetAvailableBanknotesValues()
	helpers.Slice{}.SortDesc(availableBankNotesValues)
	for value > 0 {
		for _, availableBanknote := range availableBankNotesValues {
			residual := value % availableBanknote
			if residual != value {
				quantity := value / availableBanknote
				banknoteDatas = append(banknoteDatas, models.BanknoteData{Value: int16(availableBanknote), Quantity: int16(quantity)})
				value = residual
			}
		}
	}

	return &banknoteDatas, nil
}
