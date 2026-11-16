package web

import (
	"net/http"
)

func (h *Handler) render(w http.ResponseWriter, name string, data any) error {
	return h.tmpl.ExecuteTemplate(w, name, data)
}
