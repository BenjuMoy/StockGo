package handler

import "stockgo/internal/service"

type Handler struct {
	stock *service.StockService
	sale  *service.SaleService
}

func NewHandler(
	stock *service.StockService,
	sale *service.SaleService,
) *Handler {
	return &Handler{
		stock: stock,
		sale:  sale,
	}
}
