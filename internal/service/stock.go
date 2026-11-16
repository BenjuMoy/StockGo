package service

import (
	"errors"
	"fmt"
	"stockgo/internal/domain"
)

type StockRepository interface {
	Save(*domain.Product) error
	GetAll() ([]domain.Product, error)
	GetByID(int) (domain.Product, error)
	Delete(int) error
}

// StockService handles inventory operations.
type StockService struct {
	stockRepo StockRepository
	saleRepo  SaleRepository
}

func NewStockService(stockRepo StockRepository, saleRepo SaleRepository) *StockService {
	return &StockService{stockRepo: stockRepo, saleRepo: saleRepo}
}

func (s *StockService) AddStock(name string, quantity int, cost int) error {
	product, err := domain.NewProduct(name, quantity, cost)
	if err != nil {
		return fmt.Errorf(
			"creating stock: %w",
			err,
		)
	}
	return s.stockRepo.Save(product)
}

func (s *StockService) GetAll() ([]domain.Product, error) {
	return s.stockRepo.GetAll()
}

func (s *StockService) GetByID(id int) (domain.Product, error) {
	return s.stockRepo.GetByID(id)
}

func (s *StockService) Update(id int, name string, quantity int, cost int) error {
	product, err := s.stockRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf(
			"getting stock: %w",
			err,
		)
	}

	if err := product.Update(name, cost, quantity); err != nil {
		return fmt.Errorf(
			"Updating stock: %w",
			err,
		)
	}
	return s.stockRepo.Save(&product)
}

func (s *StockService) Delete(id int) error {
	// Business rule: can't delete if there are sales
	sales, err := s.saleRepo.GetByStockID(id)
	if err != nil {
		return fmt.Errorf("checking sales for stock %d: %w", id, err)
	}
	if len(sales) > 0 {
		return domain.ErrHasSales
	}
	return s.stockRepo.Delete(id)
}

func (s *StockService) AdjustQuantity(id int, delta int) error {
	stock, err := s.stockRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf(
			"getting stock: %w",
			err,
		)
	}

	if delta == 0 {
		return errors.New("delta cannot be zero")
	}

	if err := stock.AdjustQuantity(delta); err != nil {
		return fmt.Errorf(
			"adjusting stock: %w",
			err,
		)
	}

	return s.stockRepo.Save(&stock)
}
