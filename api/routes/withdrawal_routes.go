package routes

import (
	"atm-api.com/api/controllers"
	"github.com/gin-gonic/gin"
)

//WithdrawalRoutes manage withdrawal endpoints
type WithdrawalRoutes struct {
	Controller controllers.WithdrawalController
}

//CreateWithdrawalRoutes returns a new instance of CreateWithdrawalRoutes
func CreateWithdrawalRoutes(controller controllers.WithdrawalController) WithdrawalRoutes {
	return WithdrawalRoutes{Controller: controller}
}

//Setup routes for withdrawal endpoints
func (w WithdrawalRoutes) Setup(router *gin.Engine) *gin.RouterGroup {
	routes := router.Group("/withdrawals")
	{
		routes.POST("", w.Controller.New)
	}

	return routes
}
