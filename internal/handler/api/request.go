package handler

// Stock Requests

type AddProductRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type UpdateProductRequest struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"unit_cost"`
}

type AdjustStockRequest struct {
	Delta int `json:"delta"`
}

// Sale Requests

type CreateSaleRequest struct {
	StockID   int `json:"product_id"`
	SoldPrice int `json:"unit_price"`
}
