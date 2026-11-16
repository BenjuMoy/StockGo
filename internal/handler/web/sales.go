package web

import (
	"errors"
	"net/http"
	"stockgo/internal/domain"
	"strconv"
)

func (h *Handler) Sales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.sale.GetAll()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.render(w, "sales_table", sales)
}

func (h *Handler) SellProductForm(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	product, err := h.stock.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "product not found", http.StatusNotFound)
			return

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	h.render(w, "sell_product_form", product)
}

func (h *Handler) CreateSale(w http.ResponseWriter, r *http.Request) {
	productID, err := strconv.Atoi(r.FormValue("product_id"))
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	if _, err := h.sale.CreateSale(productID, price); err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "product not found", http.StatusNotFound)
			return

		case errors.Is(err, domain.ErrInvalidProductID):
			http.Error(w, "invalid product id", http.StatusBadRequest)
			return

		case errors.Is(err, domain.ErrInvalidPrice):
			http.Error(w, "invalid price", http.StatusBadRequest)
			return

		case errors.Is(err, domain.ErrInsufficientStock):
			http.Error(w, "insufficient stock", http.StatusConflict)
			return

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	products, err := h.stock.GetAll()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	sales, err := h.sale.GetAll()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	data := struct {
		Products []domain.Product
		Sales    []domain.Sale
	}{
		Products: products,
		Sales:    sales,
	}

	h.render(w, "sale_result", data)
}
