package domain

import "time"

type Sale struct {
	id        int
	ProductID int

	productName string
	unitCost    int
	unitPrice   int

	createdAt time.Time
}

func NewSale(
	productID int,
	productName string,
	unitCost,
	unitPrice int,
) (*Sale, error) {
	if productID <= 0 {
		return nil, ErrInvalidProductID
	}

	if unitPrice < 0 || unitCost < 0 {
		return nil, ErrInvalidPrice
	}

	return &Sale{
		ProductID:   productID,
		productName: productName,
		unitCost:    unitCost,
		unitPrice:   unitPrice,
	}, nil
}

func RestoreSale(id, stockID int, productName string, unitCost, soldPrice int, createdAt time.Time) Sale {
	return Sale{
		id:          id,
		ProductID:   stockID,
		productName: productName,
		unitCost:    unitCost,
		unitPrice:   soldPrice,
		createdAt:   createdAt,
	}
}

func (s *Sale) AssignID(id int) {
	if s.id != 0 {
		panic("AssignID called twice")
	}

	s.id = id
}

func (s Sale) Profit() int {
	return s.unitPrice - s.unitCost
}

func (s Sale) ID() int              { return s.id }
func (s Sale) StockID() int         { return s.ProductID }
func (s Sale) ProductName() string  { return s.productName }
func (s Sale) UnitCost() int        { return s.unitCost }
func (s Sale) SoldPrice() int       { return s.unitPrice }
func (s Sale) CreatedAt() time.Time { return s.createdAt }
