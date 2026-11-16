package service

import (
	"fmt"
	"stockgo/internal/domain"
)

type SaleRepository interface {
	Insert(sale *domain.Sale) error
	GetAll() ([]domain.Sale, error)
	GetByStockID(stockID int) ([]domain.Sale, error)
}

// SaleService handles sale operations.
type SaleService struct {
	saleRepo  SaleRepository
	stockRepo StockRepository
}

func NewSaleService(saleRepo SaleRepository, stockRepo StockRepository) *SaleService {
	return &SaleService{saleRepo: saleRepo, stockRepo: stockRepo}
}

func (s *SaleService) CreateSale(stockID int, soldPrice int) (domain.Sale, error) {
	stock, err := s.stockRepo.GetByID(stockID)
	if err != nil {
		return domain.Sale{}, fmt.Errorf(
			"getting stock: %w",
			err,
		)
	}

	sale, err := domain.NewSale(stockID, stock.Name(), stock.Cost(), soldPrice)
	if err != nil {
		return domain.Sale{}, fmt.Errorf(
			"creating sale: %w",
			err,
		)
	}

	if err := stock.Sell(1); err != nil {
		return domain.Sale{}, err
	}

	// 4. Save both (consider transactions)
	if err := s.stockRepo.Save(&stock); err != nil {
		return domain.Sale{}, fmt.Errorf(
			"saving sale: %w",
			err,
		)
	}

	if err := s.saleRepo.Insert(sale); err != nil {
		return domain.Sale{}, fmt.Errorf(
			"inserting sale: %w",
			err,
		)
	}

	return *sale, nil
}

// GetSalesByStock retrieves all sales for a stock item
func (s *SaleService) GetSalesByStock(stockID int) ([]domain.Sale, error) {
	return s.saleRepo.GetByStockID(stockID)
}

func (s *SaleService) GetAll() ([]domain.Sale, error) {
	return s.saleRepo.GetAll()
}
