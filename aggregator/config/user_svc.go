package config

import (
	user_service "LatihanGolang/user/proto/user_service/v1"
	"log"

	"google.golang.org/grpc"
)

func InitUser() user_service.UserServiceClient {
	// Code for initializing user service

	conn, err := grpc.Dial("user-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial user service: %v", err)
	}
	return user_service.NewUserServiceClient(conn)
}
