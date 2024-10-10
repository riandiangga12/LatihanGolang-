package repository

import (
	"LatihanGolang/entity"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type GormDBIface interface {
	WithContext(ctx context.Context) *gorm.DB
	Create(value interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
}

type IWalletRepository interface {
	GetWallets(ctx context.Context) ([]entity.Wallet, error)
	GetWalletByID(ctx context.Context, id int) (entity.Wallet, error)
	CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error)
	CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error)
	TopUpWallet(ctx context.Context, id int, amount float32) (entity.Wallet, error)
	GetWalletBalance(ctx context.Context, UserID int) (entity.Wallet, error)
	TransferWallet(ctx context.Context, from int, to int, amount float32) (entity.Wallet, error)
	GetTransactions(ctx context.Context, id int) ([]entity.Transaction, error)
}

type walletRepository struct {
	db GormDBIface
}

func NewWalletRepository(db GormDBIface) IWalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) GetWallets(ctx context.Context) ([]entity.Wallet, error) {
	var wallets []entity.Wallet
	if err := r.db.WithContext(ctx).Find(&wallets).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		log.Printf("Error getting wallets: %v\n", err)
		return wallets, nil
	}

	return wallets, nil
}

func (r *walletRepository) GetWalletByID(ctx context.Context, id int) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.WithContext(ctx).First(&wallet, "user_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Wallet{}, err
		}
		log.Printf("Error getting wallet by user id: %v\n", err)
		return entity.Wallet{}, err
	}

	return wallet, nil
}

func (r *walletRepository) CreateWallet(ctx context.Context, wallet *entity.Wallet) (entity.Wallet, error) {
	if err := r.db.WithContext(ctx).Create(wallet).Error; err != nil {
		log.Printf("Error creating wallet: %v\n", err)
		return entity.Wallet{}, err
	}

	return *wallet, nil
}

func (r *walletRepository) CreateTransaction(ctx context.Context, transaction *entity.Transaction) (entity.Transaction, error) {
	if err := r.db.WithContext(ctx).Create(transaction).Error; err != nil {
		log.Printf("Error creating transaction: %v\n", err)
		return entity.Transaction{}, err
	}

	return *transaction, nil
}

func (r *walletRepository) TopUpWallet(ctx context.Context, id int, amount float32) (entity.Wallet, error) {
	var existingWallet entity.Wallet
	if err := r.db.WithContext(ctx).First(&existingWallet, id).Error; err != nil {
		log.Printf("Error getting wallet by id: %v\n", err)
		return entity.Wallet{}, err
	}

	existingWallet.Balance += amount
	if err := r.db.WithContext(ctx).Save(&existingWallet).Error; err != nil {
		log.Printf("Error topping up wallet: %v\n", err)
		return entity.Wallet{}, err
	}

	return existingWallet, nil
}

func (r *walletRepository) GetWalletBalance(ctx context.Context, UserID int) (entity.Wallet, error) {
	var wallet entity.Wallet
	if err := r.db.WithContext(ctx).First(&wallet, UserID).Error; err != nil {
		log.Printf("Error getting wallet balance: %v\n", err)
		return entity.Wallet{}, err
	}

	return wallet, nil
}

func (r *walletRepository) TransferWallet(ctx context.Context, from int, to int, amount float32) (entity.Wallet, error) {
	fromWallet, err := r.GetWalletBalance(ctx, from)
	if err != nil {
		return entity.Wallet{}, err
	}

	toWallet, err := r.GetWalletBalance(ctx, to)
	if err != nil {
		return entity.Wallet{}, err
	}

	if fromWallet.Balance < amount {
		return entity.Wallet{}, errors.New("insufficient balance")
	}

	fromWallet.Balance -= amount
	toWallet.Balance += amount

	if err := r.db.WithContext(ctx).Save(&fromWallet).Error; err != nil {
		log.Printf("Error updating wallet: %v\n", err)
		return entity.Wallet{}, err
	}

	if err := r.db.WithContext(ctx).Save(&toWallet).Error; err != nil {
		log.Printf("Error updating wallet: %v\n", err)
		return entity.Wallet{}, err
	}

	return fromWallet, nil
}

func (r *walletRepository) GetTransactions(ctx context.Context, id int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := r.db.WithContext(ctx).Find(&transactions, "wallet_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		log.Printf("Error getting transactions: %v\n", err)
		return transactions, nil
	}

	return transactions, nil
}
