package server

import (
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	api "stockgo/internal/handler/api"
	"stockgo/internal/handler/web"
	"stockgo/internal/service"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(
	addr string,
	tmpl *template.Template,
	stockService *service.StockService,
	saleService *service.SaleService,
) *Server {
	webHandler := web.NewHandler(tmpl, stockService, saleService)
	apiHandler := api.NewHandler(stockService, saleService) // no template needed

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", webHandler.Home)

	// Web / HTMX
	mux.HandleFunc("GET /products/table", webHandler.Products)
	mux.HandleFunc("GET /products/form", webHandler.ProductForm)
	mux.HandleFunc("GET /products/{id}/row", webHandler.ProductRow)
	mux.HandleFunc("GET /products/{id}/edit", webHandler.EditProductForm)
	mux.HandleFunc("GET /products/{id}/sell", webHandler.SellProductForm)

	mux.HandleFunc("POST /products", webHandler.CreateProduct)
	mux.HandleFunc("PUT /products/{id}", webHandler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", webHandler.DeleteProduct)
	//mux.HandleFunc("POST /products/{id}/restock", webHandler.RestockProduct)

	mux.HandleFunc("GET /sales/table", webHandler.Sales)
	mux.HandleFunc("POST /sales", webHandler.CreateSale)

	// API (JSON)
	mux.HandleFunc("GET /api/products", apiHandler.GetProducts)
	mux.HandleFunc("GET /api/products/{id}", apiHandler.GetProduct)
	mux.HandleFunc("POST /api/products", apiHandler.CreateProduct)
	mux.HandleFunc("PUT /api/products/{id}", apiHandler.UpdateProduct)
	mux.HandleFunc("DELETE /api/products/{id}", apiHandler.DeleteProduct)

	mux.HandleFunc("GET /api/sales", apiHandler.GetSales)
	mux.HandleFunc("POST /api/sales", apiHandler.CreateSale)

	return &Server{
		httpServer: &http.Server{
			Addr:              addr,
			Handler:           mux,
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       15 * time.Second,
			WriteTimeout:      15 * time.Second,
			IdleTimeout:       60 * time.Second,
		},
	}
}

func (s *Server) ListenAndServe() error {
	slog.Info("starting HTTP server", "addr", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
