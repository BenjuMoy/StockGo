package main

import (
	"log"
	"log/slog"
	"os"

	"fyne.io/fyne/v2/app"

	"stockgo/internal/config"
	"stockgo/internal/repository"
	"stockgo/internal/service"
	"stockgo/ui"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(
				os.Stderr,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		),
	)

	db, err := repository.OpenDB(cfg.DBPath())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stockRepo := repository.NewStockRepo(db)
	saleRepo := repository.NewSaleRepo(db)
	reportRepo := repository.NewReportRepo(db)

	services := service.Services{
		Stock:  service.NewStockService(stockRepo, saleRepo),
		Sale:   service.NewSaleService(saleRepo, stockRepo),
		Report: service.NewReportService(reportRepo),
	}

	ui.New(app.NewWithID("0"), services).Run()
}
