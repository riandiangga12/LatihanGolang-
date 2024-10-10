package service

import (
	"LatihanGolang/entity"
	walletGorm "LatihanGolang/wallet/repository"
	"context"
	"fmt"
)

type IWalletService interface {
	GetWallets(ctx context.Context) ([]entity.Wallet, error)
	GetWalletByID(ctx context.Context, id int) (entity.Wallet, error)
	CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error)
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error)
	TopUpWallet(ctx context.Context, id int, amount float32) (entity.Wallet, error)
	TransferWallet(ctx context.Context, from int, to int, amount float32) (entity.Wallet, error)
	GetWalletBalance(ctx context.Context, UserID int) (entity.Wallet, error)
	GetTransactions(ctx context.Context, id int) ([]entity.Transaction, error)
}

type walletService struct {
	walletRepo walletGorm.IWalletRepository
}

func NewWalletService(walletRepo walletGorm.IWalletRepository) IWalletService {
	return &walletService{walletRepo: walletRepo}
}

func (s *walletService) GetWallets(ctx context.Context) ([]entity.Wallet, error) {
	wallets, err := s.walletRepo.GetWallets(ctx)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (s *walletService) GetWalletByID(ctx context.Context, id int) (entity.Wallet, error) {
	wallet, err := s.walletRepo.GetWalletByID(ctx, id)
	if err != nil {
		return entity.Wallet{}, err
	}

	return wallet, nil
}

func (s *walletService) CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	createdWallet, err := s.walletRepo.CreateWallet(ctx, wallet)
	if err != nil {
		return entity.Wallet{}, err
	}

	return createdWallet, nil
}

func (s *walletService) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error) {
	createdTransaction, err := s.walletRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return entity.Transaction{}, fmt.Errorf("error creating transaction: %v", err)
	}

	return createdTransaction, nil
}

func (s *walletService) TopUpWallet(ctx context.Context, id int, amount float32) (entity.Wallet, error) {
	updatedWallet, err := s.walletRepo.TopUpWallet(ctx, id, amount)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error topping up wallet: %v", err)
	}

	return updatedWallet, nil
}

func (s *walletService) TransferWallet(ctx context.Context, from int, to int, amount float32) (entity.Wallet, error) {
	updatedWallet, err := s.walletRepo.TransferWallet(ctx, from, to, amount)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error transferring wallet: %v", err)
	}

	return updatedWallet, nil
}

func (s *walletService) GetWalletBalance(ctx context.Context, UserID int) (entity.Wallet, error) {
	wallet, err := s.walletRepo.GetWalletBalance(ctx, UserID)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error getting wallet balance: %v", err)
	}

	return wallet, nil
}

func (s *walletService) GetTransactions(ctx context.Context, id int) ([]entity.Transaction, error) {
	transactions, err := s.walletRepo.GetTransactions(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting transactions: %v", err)
	}

	return transactions, nil
}
