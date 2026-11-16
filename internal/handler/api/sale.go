package handler

import (
	"net/http"
	"strconv"
)

func (h *Handler) CreateSale(w http.ResponseWriter, r *http.Request) {
	var req CreateSaleRequest

	if err := decodeJSON(r, &req); err != nil {
		handleServiceError(w, err)
		return
	}

	sale, err := h.sale.CreateSale(req.StockID, req.SoldPrice)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, sale)
}

func (h *Handler) GetSalesByStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	sales, err := h.sale.GetSalesByStock(id)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, sales)
}

//func (h *Handler) GetRevenue(w http.ResponseWriter, r *http.Request) {
//	revenue, err := h.sale.GetTotalRevenue()
//	if err != nil {
//		handleServiceError(w, err)
//		return
//	}
//
//	writeJSON(w, http.StatusOK, map[string]int{
//		"revenue": revenue,
//	})
//}

func (h *Handler) GetSales(w http.ResponseWriter, r *http.Request) {
	sales, err := h.sale.GetAll()
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, sales)
}
