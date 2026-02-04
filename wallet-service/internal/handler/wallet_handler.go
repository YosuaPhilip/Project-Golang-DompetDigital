package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"wallet-service/internal/service"
)

type WalletHandler struct {
	svc *service.WalletService
}

func NewWalletHandler(svc *service.WalletService) *WalletHandler {
	return &WalletHandler{svc: svc}
}

func (h *WalletHandler) Balance(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id"})
		return
	}

	bal, err := h.svc.GetBalance(c.Request.Context(), userID)
	if err != nil {
		code, msg := MapErrorToHTTP(err)
		c.JSON(code, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"balance": bal,
	})
}

type WithdrawRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

func (h *WalletHandler) Withdraw(c *gin.Context) {
	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id must be > 0"})
		return
	}
	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "amount must be > 0"})
		return
	}

	res, err := h.svc.Withdraw(c.Request.Context(), req.UserID, req.Amount)
	if err != nil {
		code, msg := MapErrorToHTTP(err)
		c.JSON(code, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "withdraw success",
		"user_id":         res.UserID,
		"withdraw_amount": res.WithdrawAmount,
		"balance":         res.Balance,
	})
}
