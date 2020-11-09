package controllers

import (
	"atm-api.com/api/dtos"
	"atm-api.com/services"
	"github.com/gin-gonic/gin"
)

//WithdrawalController withdrawals REST bindings
type WithdrawalController struct {
	BanknoteDataService services.BanknoteDataService
}

//CreateWithdrawalController returns a new instance of WithdrawalController
func CreateWithdrawalController(banknoteDataService services.BanknoteDataService) WithdrawalController {
	return WithdrawalController{BanknoteDataService: banknoteDataService}
}

//New func to calculate new Withdrawal
func (w *WithdrawalController) New(ctx *gin.Context) {
	payload := dtos.WithdrawalNewRequest{}
	ctx.ShouldBind(&payload)
	banknotes, err := w.BanknoteDataService.Withdrawal(payload.Value)

	if err != nil {
		ctx.JSON(422, err)
		return
	}

	ctx.JSON(201, banknotes)
}
