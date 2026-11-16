package handler

import (
	"net/http"
	"strconv"
)

// GetProducts returns a list of all products
func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	// 1. Call Service
	products, err := h.stock.GetAll()
	if err != nil {
		handleServiceError(w, err)
		return
	}

	// 2. Set Header
	writeJSON(w, http.StatusOK, products)
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	product, err := h.stock.GetByID(id)
	if err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, product)
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req AddProductRequest

	// 1. Decode JSON Body
	if err := decodeJSON(r, &req); err != nil {
		handleServiceError(w, err)
		return
	}

	// 2. Call Service
	if err := h.stock.AddStock(req.Name, req.Quantity, req.Price); err != nil {
		handleServiceError(w, err)
		return
	}

	// 3. Return Success (201 Created)
	writeJSON(w, http.StatusCreated, map[string]string{
		"message": "product Created",
	})
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req UpdateProductRequest

	if err := decodeJSON(r, &req); err != nil {
		handleServiceError(w, err)
		return
	}

	if err := h.stock.Update(id, req.Name, req.Quantity, req.Price); err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.stock.Delete(id); err != nil {
		handleServiceError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AdjustStock(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req AdjustStockRequest

	if err := decodeJSON(r, &req); err != nil {
		handleServiceError(w, err)
		return
	}

	if err := h.stock.AdjustQuantity(id, req.Delta); err != nil {
		handleServiceError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "stock adjusted",
	})
}
