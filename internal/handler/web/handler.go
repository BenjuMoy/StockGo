package web

import (
	"html/template"
	"stockgo/internal/service"
)

type Handler struct {
	tmpl  *template.Template
	stock *service.StockService
	sale  *service.SaleService
}

func NewHandler(
	tmpl *template.Template,
	stock *service.StockService,
	sale *service.SaleService,
) *Handler {
	return &Handler{
		tmpl:  tmpl,
		stock: stock,
		sale:  sale,
	}
}
