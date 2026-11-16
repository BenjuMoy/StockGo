package service

import "stockgo/internal/domain"

type Stock interface {
	AddStock(name string, qty, cost int) error
	GetAll() ([]domain.Product, error)
	GetByID(id int) (domain.Product, error)
	Update(id int, name string, cost int) error
	Delete(id int) error
	AdjustQuantity(id int, delta int) error
}

type Sale interface {
	CreateSale(stockID int, soldPrice int) (domain.Sale, error)
	GetSalesByStock(stockID int) ([]domain.Sale, error)
	GetTotalRevenue() (int, error)
	GetAll() ([]domain.Sale, error)
}
