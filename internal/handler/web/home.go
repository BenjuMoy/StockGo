package web

import (
	"log"
	"net/http"

	"stockgo/internal/domain"
)

type HomeData struct {
	Products []domain.Product
	Sales    []domain.Sale
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	products, err := h.stock.GetAll()
	if err != nil {
		log.Printf("load products: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	sales, err := h.sale.GetAll()
	if err != nil {
		log.Printf("load sales: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data := HomeData{Products: products, Sales: sales}

	h.render(w, "layout", data)
}
