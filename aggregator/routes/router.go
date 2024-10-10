package router

import (
	"LatihanGolang/aggregator/handler"
	"LatihanGolang/aggregator/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, aggregator handler.IAggregatorHandler) {
	userEndpoint := r.Group("/user")
	userEndpoint.Use(middleware.AuthMiddleware())
	userEndpoint.GET("/:id", aggregator.GetUserAndWallet)
	userEndpoint.POST("/create", aggregator.CreateUserAndWallet)

	walletEndpoint := r.Group("/wallet")
	walletEndpoint.Use(middleware.AuthMiddleware())
	walletEndpoint.POST("/topup", aggregator.TopUpWallet)
	walletEndpoint.POST("/transfer", aggregator.TransferWallet)
	walletEndpoint.GET("/transactions/:id", aggregator.GetTransactions)
}
