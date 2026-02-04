package main

import (
	"log"

	"wallet-service/internal/config"
	"wallet-service/internal/db"
	"wallet-service/internal/handler"
	"wallet-service/internal/repository"
	"wallet-service/internal/router"
	"wallet-service/internal/service"
)

func main() {
	cfg := config.Load()
	if cfg.DBDSN == "" {
		log.Fatal("DB_DSN is required")
	}

	dbConn, err := db.NewPostgres(cfg.DBDSN)
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	repo := repository.NewWalletRepository(dbConn)
	svc := service.NewWalletService(repo)
	h := handler.NewWalletHandler(svc)

	r := router.New(h)

	addr := ":" + cfg.Port
	log.Printf("wallet-service listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
