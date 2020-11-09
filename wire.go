//+build wireinject

package main

import (
	"atm-api.com/api/controllers"
	"atm-api.com/api/routes"
	"atm-api.com/repositories/withdrawal"
	"atm-api.com/services"
	"github.com/google/wire"
)

func InitializeWithdrawalRoutes() routes.WithdrawalRoutes {
	wire.Build(
		withdrawal.CreateDefaultRepository,
		services.CreateBanknoteService,
		controllers.CreateWithdrawalController,
		routes.CreateWithdrawalRoutes,
	)
	return routes.WithdrawalRoutes{}
}
