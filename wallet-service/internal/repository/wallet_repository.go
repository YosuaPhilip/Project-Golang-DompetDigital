package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	errX "wallet-service/internal/errors"
)

type WalletRepository struct {
	db *sqlx.DB
}

func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db: db}
}

func (r *WalletRepository) GetBalance(ctx context.Context, userID int64) (int64, error) {
	var bal int64
	err := r.db.QueryRowContext(ctx, `SELECT balance FROM wallets WHERE user_id=$1`, userID).Scan(&bal)
	if err == sql.ErrNoRows {
		return 0, errX.ErrUserNotFound
	}
	return bal, err
}

type WithdrawResult struct {
	UserID         int64 `json:"user_id"`
	WithdrawAmount int64 `json:"withdraw_amount"`
	Balance        int64 `json:"balance"`
}

func (r *WalletRepository) WithdrawTx(ctx context.Context, userID, amount int64) (*WithdrawResult, error) {
	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var balance int64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM wallets WHERE user_id=$1 FOR UPDATE`, userID).Scan(&balance)
	if err == sql.ErrNoRows {
		return nil, errX.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	status := "SUCCESS"
	if balance < amount {
		status = "FAILED"
	}

	_, _ = tx.ExecContext(ctx,
		`INSERT INTO transactions(id,user_id,type,amount,status)
		 VALUES($1,$2,'WITHDRAW',$3,$4)`,
		uuid.New().String(), userID, amount, status,
	)

	if status == "FAILED" {
		return nil, errX.ErrInsufficientBalance
	}

	_, err = tx.ExecContext(ctx,
		`UPDATE wallets SET balance = balance - $1, updated_at=NOW() WHERE user_id=$2`,
		amount, userID,
	)
	if err != nil {
		return nil, err
	}

	var newBalance int64
	err = tx.QueryRowContext(ctx, `SELECT balance FROM wallets WHERE user_id=$1`, userID).Scan(&newBalance)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &WithdrawResult{
		UserID:         userID,
		WithdrawAmount: amount,
		Balance:        newBalance,
	}, nil
}
