package router

import (
	"github.com/gin-gonic/gin"

	"wallet-service/internal/handler"
)

func New(h *handler.WalletHandler) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		wallet := v1.Group("/wallet")
		{
			wallet.GET("/balance", h.Balance)
			wallet.POST("/withdraw", h.Withdraw)
		}
	}

	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })
	return r
}
