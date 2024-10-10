package main

import (
	"LatihanGolang/entity"
	grpcHandler "LatihanGolang/wallet/handler"
	pb "LatihanGolang/wallet/proto/wallet_service/v1"
	walletgorm "LatihanGolang/wallet/repository"
	"LatihanGolang/wallet/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "postgresql://postgres:password@pg-db_wallet:5432/go_db_wallet"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalln(err)
	}

	//run migration
	if err = gormDB.AutoMigrate(&entity.Wallet{}); err != nil {
		log.Fatalln("Failed to migrate:", err)
	}

	if err = gormDB.AutoMigrate(&entity.Transaction{}); err != nil {
		log.Fatalln("Failed to migrate:", err)
	}

	//setup service
	walletRepo := walletgorm.NewWalletRepository(gormDB)
	walletService := service.NewWalletService(walletRepo)
	walletHandler := grpcHandler.NewWalletHandler(walletService)

	//run the grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterWalletServiceServer(grpcServer, walletHandler)

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Server is running on port 50052")
	_ = grpcServer.Serve(lis)
}
