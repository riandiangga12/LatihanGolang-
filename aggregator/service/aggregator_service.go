package service

import (
	"LatihanGolang/entity"
	userService "LatihanGolang/user/proto/user_service/v1"
	walletService "LatihanGolang/wallet/proto/wallet_service/v1"
	"context"
)

// IAggregatorService is an interface that defines all methods that should be implemented by AggregatorService
type IAggregatorService interface {
	CreateUserAndWallet(ctx context.Context, user *entity.User) (entity.User, entity.Wallet, error)
	GetUserAndWallet(ctx context.Context, userID int) (entity.User, entity.Wallet, error)
	TopUpWallet(ctx context.Context, id int, amount float32) (walletService.WalletMutationResponse, error)
	TransferWallet(ctx context.Context, from int, to int, amount float32) (walletService.WalletMutationResponse, error)
	GetTransactions(ctx context.Context, id int) ([]entity.Transaction, error)
}

type aggregatorService struct {
	userService   userService.UserServiceClient
	walletService walletService.WalletServiceClient
}

// NewAggregatorService creates a new instance of AggregatorService
func NewAggregatorService(userService userService.UserServiceClient, walletService walletService.WalletServiceClient) IAggregatorService {
	return &aggregatorService{
		userService:   userService,
		walletService: walletService,
	}
}

// CreateUserAndWallet is a method to create user and wallet
func (s *aggregatorService) CreateUserAndWallet(ctx context.Context, user *entity.User) (entity.User, entity.Wallet, error) {
	// Create User
	userResponse, err := s.userService.CreateUser(ctx, &userService.CreateUserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if err != nil {
		return entity.User{}, entity.Wallet{}, err
	}

	// Create Wallet
	_, err = s.walletService.CreateWallet(ctx, &walletService.CreateWalletRequest{
		UserId: userResponse.Id,
	})

	if err != nil {
		return entity.User{}, entity.Wallet{}, err
	}

	// Get User and Wallet
	users, wallet, err := s.GetUserAndWallet(ctx, int(userResponse.Id))
	if err != nil {
		return entity.User{}, entity.Wallet{}, err
	}

	return users, wallet, nil

}

// GetUserAndWallet is a method to get user and wallet by user ID
func (s *aggregatorService) GetUserAndWallet(ctx context.Context, userID int) (entity.User, entity.Wallet, error) {
	// Get User
	userResponse, err := s.userService.GetUserByID(ctx, &userService.GetUserByIDRequest{
		Id: int32(userID),
	})
	if err != nil {
		return entity.User{}, entity.Wallet{}, err
	}

	// Get Wallet
	walletResponse, err := s.walletService.GetWalletByID(ctx, &walletService.GetWalletByIDRequest{
		Id: int32(userResponse.User.Id),
	})
	if err != nil {
		return entity.User{}, entity.Wallet{}, err
	}

	return entity.User{
			ID:    int(userResponse.User.Id),
			Name:  userResponse.User.Name,
			Email: userResponse.User.Email,
		}, entity.Wallet{
			ID:      int(walletResponse.Wallet.Id),
			UserID:  int(walletResponse.Wallet.UserId),
			Balance: walletResponse.Wallet.Balance,
		}, nil
}

// TopUpWallet is a method to top up wallet by wallet ID
func (s *aggregatorService) TopUpWallet(ctx context.Context, id int, amount float32) (walletService.WalletMutationResponse, error) {
	// Top Up Wallet
	_, err := s.walletService.TopUpWallet(ctx, &walletService.TopUpWalletRequest{
		UserId: int32(id),
		Amount: amount,
	})

	if err != nil {
		return walletService.WalletMutationResponse{
			Message: "Failed to top up wallet",
		}, err
	}

	// Create Transaction
	_, err = s.walletService.CreateTransaction(ctx, &walletService.CreateTransactionRequest{
		WalletId: int32(id),
		Amount:   amount,
		Type:     "topup",
	})

	if err != nil {
		return walletService.WalletMutationResponse{
			Message: "Failed to create transaction",
		}, err
	}

	return walletService.WalletMutationResponse{
		Message: "Wallet has been topped up",
	}, nil

}

// TransferWallet is a method to transfer wallet balance from one wallet to another
func (s *aggregatorService) TransferWallet(ctx context.Context, from int, to int, amount float32) (walletService.WalletMutationResponse, error) {
	// Transfer Wallet
	_, err := s.walletService.TransferWallet(ctx, &walletService.TransferWalletRequest{
		From:   int32(from),
		To:     int32(to),
		Amount: amount,
	})

	if err != nil {
		return walletService.WalletMutationResponse{
			Message: "Failed to transfer wallet",
		}, err
	}

	// Create Transaction Transfer
	_, err = s.walletService.CreateTransaction(ctx, &walletService.CreateTransactionRequest{
		WalletId: int32(from),
		WalletTo: int32(to),
		Amount:   amount * -1,
		Type:     "transfer",
	})

	if err != nil {
		return walletService.WalletMutationResponse{
			Message: "Failed to create transaction",
		}, err
	}

	// Create Transaction Receive
	_, err = s.walletService.CreateTransaction(ctx, &walletService.CreateTransactionRequest{
		WalletId: int32(to),
		WalletTo: int32(from),
		Amount:   amount,
		Type:     "receive",
	})

	if err != nil {
		return walletService.WalletMutationResponse{
			Message: "Failed to create transaction",
		}, err
	}

	return walletService.WalletMutationResponse{
		Message: "Wallet has been transferred",
	}, nil

}

// GetTransactions is a method to get all transactions by wallet ID
func (s *aggregatorService) GetTransactions(ctx context.Context, id int) ([]entity.Transaction, error) {
	// Get Transactions
	transactions, err := s.walletService.GetTransactions(ctx, &walletService.GetTransactionsRequest{
		UserId: int32(id),
	})
	if err != nil {
		return nil, err
	}

	var transactionsResponse []entity.Transaction
	for _, transaction := range transactions.Transactions {
		transactionsResponse = append(transactionsResponse, entity.Transaction{
			ID:       int(transaction.Id),
			WalletID: int(transaction.WalletId),
			WalletTo: int(transaction.WalletTo),
			Amount:   transaction.Amount,
			Type:     transaction.Type,
		})

	}

	return transactionsResponse, nil

}
