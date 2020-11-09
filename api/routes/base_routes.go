package routes

import "github.com/gin-gonic/gin"

//BaseRoutes defines basic methods for a **Routes
type BaseRoutes interface {
	Setup(router *gin.Engine) *gin.RouterGroup
}
