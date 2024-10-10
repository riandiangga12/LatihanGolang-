package config

import (
	wallet_service "LatihanGolang/wallet/proto/wallet_service/v1"
	"log"

	"google.golang.org/grpc"
)

func InitWallet() wallet_service.WalletServiceClient {
	conn, err := grpc.Dial("wallet-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial wallet service: %v", err)
	}
	return wallet_service.NewWalletServiceClient(conn)
}
