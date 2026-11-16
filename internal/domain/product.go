package domain

import (
	"strings"
)

const iva = 0.21 // Fixed taxes

type Product struct {
	id       int
	name     string
	quantity int
	unitCost int // This means bought price
}

func NewProduct(name string, quantity int, cost int) (*Product, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, ErrEmptyName
	}
	if quantity < 0 {
		return nil, ErrInvalidQuantity
	}
	if cost <= 0 {
		return nil, ErrInvalidPrice
	}

	return &Product{
		name:     name,
		quantity: quantity,
		unitCost: cost,
	}, nil
}

func RestoreProduct(id int, name string, quantity int, cost int) Product {
	return Product{
		id:       id,
		name:     name,
		quantity: quantity,
		unitCost: cost,
	}
}

func (p *Product) AssignID(id int) {
	if p.id != 0 {
		panic("AssignID called twice")
	}
	if id <= 0 {
		panic("AssignID called with invalid ID")
	}

	p.id = id
}

func (p *Product) Sell(qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}

	if p.quantity < qty {
		return ErrInsufficientStock
	}

	p.quantity -= qty
	return nil
}

func (p *Product) Update(name string, cost int, quantity int) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return ErrEmptyName
	}

	if cost <= 0 {
		return ErrInvalidPrice
	}

	if quantity < 0 {
		return ErrInvalidQuantity
	}

	p.name = name
	p.unitCost = cost
	p.quantity = quantity

	return nil
}

func (p *Product) AdjustQuantity(delta int) error {
	if p.quantity+delta < 0 {
		return ErrInsufficientStock
	}

	p.quantity += delta
	return nil
}

func (p Product) ID() int            { return p.id }
func (p Product) Name() string       { return p.name }
func (p Product) Quantity() int      { return p.quantity }
func (p Product) Cost() int          { return p.unitCost }
func (p Product) TaxedCost() float64 { return float64(p.unitCost) * (1 + iva) }
