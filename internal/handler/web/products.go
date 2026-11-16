package web

import (
	"errors"
	"net/http"
	"stockgo/internal/domain"
	"strconv"
)

// VIEWS

func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	products, err := h.stock.GetAll()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.render(w, "product_table", products)
}

func (h *Handler) ProductRow(w http.ResponseWriter, r *http.Request) {
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

	h.render(w, "product_row", product)
}

func (h *Handler) ProductForm(w http.ResponseWriter, r *http.Request) {
	h.render(w, "add_product_form", nil)
}

func (h *Handler) EditProductForm(w http.ResponseWriter, r *http.Request) {
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

	h.render(w, "edit_product_form", product)
}

// ACTIONS

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	qty, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Error(w, "invalid quantity", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	if err = h.stock.AddStock(r.FormValue("name"), qty, price); err != nil {
		switch {
		case errors.Is(err, domain.ErrEmptyName):
			http.Error(w, "product name cannot be empty", http.StatusBadRequest)
			return

		case errors.Is(err, domain.ErrInvalidQuantity):
			http.Error(w, "invalid quantity", http.StatusBadRequest)
			return

		case errors.Is(err, domain.ErrInvalidPrice):
			http.Error(w, "invalid price", http.StatusBadRequest)
			return

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	h.Products(w, r)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	quantity, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "invalid quantity", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}

	if err = h.stock.Update(id, name, quantity, price); err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "product not found", http.StatusNotFound)
			return

		case errors.Is(err, domain.ErrInvalidInput):
			http.Error(w, "invalid product data", http.StatusBadRequest)
			return

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	updated, err := h.stock.GetByID(id)
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

	h.render(w, "product_row", updated)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.stock.Delete(id); err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "product not found", http.StatusNotFound)
			return

		case errors.Is(err, domain.ErrHasSales):
			http.Error(w, "cannot delete a product with sales", http.StatusConflict)
			return

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)

	// h.Products(w, r)
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AdjustQuantity(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	delta, err := strconv.Atoi(r.FormValue("delta"))
	if err != nil {
		http.Error(w, "invalid delta", http.StatusBadRequest)
		return
	}

	if err = h.stock.AdjustQuantity(id, delta); err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			http.Error(w, "product not found", http.StatusNotFound)
			return

		case errors.Is(err, domain.ErrInvalidInput):
			http.Error(w, "invalid quantity adjustment", http.StatusBadRequest)
			return

		case errors.Is(err, domain.ErrInsufficientStock):
			http.Error(w, "insufficient stock", http.StatusConflict)
			return

		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
