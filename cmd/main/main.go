package main

import (
	"fmt"
	"log/slog"
	"os"
	"stockgo/internal/config"
	"stockgo/internal/handler/web"
	"stockgo/internal/repository"
	"stockgo/internal/server"
	"stockgo/internal/service"
)

func main() {
	if err := run(); err != nil {
		slog.Error("StockGo stopped", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	db, err := repository.OpenDB(cfg.DBPath())
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	stockRepo := repository.NewStockRepo(db)
	saleRepo := repository.NewSaleRepo(db)

	stockService := service.NewStockService(stockRepo, saleRepo)
	saleService := service.NewSaleService(saleRepo, stockRepo)

	tmpl, err := web.LoadTemplates()
	if err != nil {
		return fmt.Errorf("load templates: %w", err)
	}

	srv := server.New(
		cfg.Addr,
		tmpl,
		stockService,
		saleService,
	)

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("serve HTTP: %w", err)
	}

	return nil
}
