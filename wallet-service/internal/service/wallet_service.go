package service

import (
	"context"

	"wallet-service/internal/repository"
)

type WalletService struct {
	repo *repository.WalletRepository
}

func NewWalletService(repo *repository.WalletRepository) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) GetBalance(ctx context.Context, userID int64) (int64, error) {
	return s.repo.GetBalance(ctx, userID)
}

func (s *WalletService) Withdraw(ctx context.Context, userID, amount int64) (*repository.WithdrawResult, error) {
	return s.repo.WithdrawTx(ctx, userID, amount)
}
