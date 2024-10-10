package main

import (
	"LatihanGolang/aggregator/config"
	"LatihanGolang/aggregator/handler"
	router "LatihanGolang/aggregator/routes"
	"LatihanGolang/aggregator/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	//Connect to user service
	userClient := config.InitUser()
	if userClient == nil {
		log.Fatalln("Failed to connect to user service")
	}

	//Connect to wallet service
	walletClient := config.InitWallet()
	if walletClient == nil {
		log.Fatalln("Failed to connect to wallet service")
	}

	aggregatorService := service.NewAggregatorService(userClient, walletClient)
	aggregatorHandler := handler.NewAggregatorHandler(aggregatorService)

	router.SetupRouter(r, aggregatorHandler)

	log.Println("Server is running on port 8080")
	r.Run(":8080")
}
