package services

import (
	"fmt"

	"atm-api.com/exceptions"
	"atm-api.com/helpers"
	"atm-api.com/models"
	"atm-api.com/repositories/withdrawal"
	log "github.com/sirupsen/logrus"
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
func (b *BanknoteDataService) Withdrawal(value int) (*[]models.BanknoteData, error) {
	banknoteDatas := make([]models.BanknoteData, 0)
	valueLeftOver := value
	availableBankNotesValues := b.WithdrawalRepository.GetAvailableBanknotesValues()
	log.Debugf("Retrieved available bank notes: %v", availableBankNotesValues)
	helpers.Slice{}.SortDesc(availableBankNotesValues)
	for _, availableBanknote := range availableBankNotesValues {
		residual := valueLeftOver % availableBanknote
		if residual != valueLeftOver {
			quantity := valueLeftOver / availableBanknote
			banknoteDatas = append(banknoteDatas, models.BanknoteData{Value: int16(availableBanknote), Quantity: int16(quantity)})
			valueLeftOver = residual
		}
	}

	if valueLeftOver > 0 {
		log.Errorf("Value not supported. Value: %v | Left Over: %v", value, valueLeftOver)
		return nil, &exceptions.UnsupportedValueError{
			Reason: fmt.Sprintf("Withdrawal not supported for value, left over: %v", value),
		}
	}

	return &banknoteDatas, nil
}
