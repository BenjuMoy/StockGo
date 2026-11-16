package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"stockgo/internal/domain"
)

func handleServiceError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		writeError(w, http.StatusNotFound, "resource not found")

	case errors.Is(err, domain.ErrHasSales):
		writeError(w, http.StatusConflict, "cannot delete a product with sales")

	case errors.Is(err, domain.ErrInsufficientStock):
		writeError(w, http.StatusConflict, "insufficient stock")

	case errors.Is(err, domain.ErrInvalidProductID):
		writeError(w, http.StatusBadRequest, "invalid stock id")

	case errors.Is(err, domain.ErrInvalidPrice),
		errors.Is(err, domain.ErrEmptyName),
		errors.Is(err, domain.ErrInvalidQuantity):
		writeError(w, http.StatusBadRequest, "invalid request")

	default:
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func decodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must contain a single JSON object")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("encoding JSON response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}
