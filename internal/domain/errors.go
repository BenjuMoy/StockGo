package domain

import "errors"

var (
	ErrNotFound          = errors.New("resource not found")
	ErrEmptyName         = errors.New("product name cannot be empty")
	ErrInvalidQuantity   = errors.New("quantity cannot be negative")
	ErrInvalidPrice      = errors.New("price must be positive")
	ErrInsufficientStock = errors.New("insufficient stock")
	ErrInvalidProductID  = errors.New("invalid product ID")
	ErrHasSales          = errors.New("product has associated sales")
	ErrInvalidInput      = errors.New("Input invalido")
)
