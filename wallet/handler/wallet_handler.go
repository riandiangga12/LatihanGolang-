package handler

import (
	"LatihanGolang/entity"
	pb "LatihanGolang/wallet/proto/wallet_service/v1"
	"LatihanGolang/wallet/service"
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WalletHandler struct {
	pb.UnimplementedWalletServiceServer
	walletService service.IWalletService
}

func NewWalletHandler(walletService service.IWalletService) *WalletHandler {
	return &WalletHandler{
		walletService: walletService,
	}
}

func (h *WalletHandler) GetWallets(ctx context.Context, _ *emptypb.Empty) (*pb.GetWalletsResponse, error) {
	wallets, err := h.walletService.GetWallets(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var pbWallets []*pb.Wallet
	for _, wallet := range wallets {
		pbWallets = append(pbWallets, &pb.Wallet{
			Id:        int32(wallet.ID),
			UserId:    int32(wallet.UserID),
			Balance:   wallet.Balance,
			CreatedAt: timestamppb.New(wallet.CreatedAt),
			UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		})
	}

	return &pb.GetWalletsResponse{
		Wallets: pbWallets,
	}, nil
}

func (h *WalletHandler) CreateWallet(ctx context.Context, req *pb.CreateWalletRequest) (*pb.WalletMutationResponse, error) {
	createdWallet, err := h.walletService.CreateWallet(ctx, &entity.Wallet{
		UserID: int(req.UserId),
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.WalletMutationResponse{
		Message: fmt.Sprintf("Wallet with ID %d has been created", createdWallet.ID),
	}, nil
}

func (h *WalletHandler) GetWalletByID(ctx context.Context, req *pb.GetWalletByIDRequest) (*pb.GetWalletByIDResponse, error) {
	wallet, err := h.walletService.GetWalletByID(ctx, int(req.Id))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetWalletByIDResponse{
		Wallet: &pb.Wallet{
			Id:        int32(wallet.ID),
			UserId:    int32(wallet.UserID),
			Balance:   wallet.Balance,
			CreatedAt: timestamppb.New(wallet.CreatedAt),
			UpdatedAt: timestamppb.New(wallet.UpdatedAt),
		},
	}, nil
}

func (h *WalletHandler) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.WalletMutationResponse, error) {
	createdTransaction, err := h.walletService.CreateTransaction(ctx, &entity.Transaction{
		WalletID: int(req.WalletId),
		WalletTo: int(req.WalletTo),
		Amount:   req.Amount,
		Type:     req.Type,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.WalletMutationResponse{
		Message: fmt.Sprintf("Transaction with ID %d has been created", createdTransaction.ID),
	}, nil
}

func (h *WalletHandler) TopUpWallet(ctx context.Context, req *pb.TopUpWalletRequest) (*pb.WalletMutationResponse, error) {
	topupWallet, err := h.walletService.TopUpWallet(ctx, int(req.UserId), req.Amount)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.WalletMutationResponse{
		Message: fmt.Sprintf("Wallet with ID %d has been topped up", topupWallet.ID),
	}, nil
}

func (h *WalletHandler) TransferWallet(ctx context.Context, req *pb.TransferWalletRequest) (*pb.WalletMutationResponse, error) {
	transferWallet, err := h.walletService.TransferWallet(ctx, int(req.From), int(req.To), req.Amount)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.WalletMutationResponse{
		Message: fmt.Sprintf("Wallet with ID %d has been transferred", transferWallet.ID),
	}, nil
}

func (h *WalletHandler) GetWalletBalance(ctx context.Context, req *pb.GetWalletBalanceRequest) (*pb.GetWalletBalanceResponse, error) {
	wallet, err := h.walletService.GetWalletBalance(ctx, int(req.UserId))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetWalletBalanceResponse{
		Balance: wallet.Balance,
	}, nil

}

func (h *WalletHandler) GetTransactions(ctx context.Context, req *pb.GetTransactionsRequest) (*pb.GetTransactionsResponse, error) {
	transactions, err := h.walletService.GetTransactions(ctx, int(req.UserId))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var pbTransactions []*pb.Transaction
	for _, transaction := range transactions {
		pbTransactions = append(pbTransactions, &pb.Transaction{
			Id:        int32(transaction.ID),
			WalletId:  int32(transaction.WalletID),
			WalletTo:  int32(transaction.WalletTo),
			Amount:    transaction.Amount,
			Type:      transaction.Type,
			CreatedAt: timestamppb.New(transaction.CreatedAt),
		})
	}

	return &pb.GetTransactionsResponse{
		Transactions: pbTransactions,
	}, nil
}
