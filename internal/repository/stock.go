package repository

import (
	"database/sql"
	"errors"
	"stockgo/internal/domain"
)

var ErrStockNotFound = errors.New("product not found")

type StockRepository struct {
	db *sql.DB
}

func NewStockRepo(db *sql.DB) *StockRepository {
	return &StockRepository{db: db}
}

func (r *StockRepository) Save(stock *domain.Product) error {
	if stock.ID() > 0 {
		res, err := r.db.Exec(
			`UPDATE products SET name = ?, quantity = ?, unit_cost = ? WHERE id = ?`,
			stock.Name(), stock.Quantity(), stock.Cost(), stock.ID(),
		)
		if err != nil {
			return err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			return err
		}

		if rows == 0 {
			return ErrStockNotFound
		}

	} else {

		res, err := r.db.Exec(
			`INSERT INTO products (name, quantity, unit_cost) VALUES (?, ?, ?)`,
			stock.Name(), stock.Quantity(), stock.Cost(),
		)
		if err != nil {
			return err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return err
		}

		stock.AssignID(int(id))
	}

	return nil
}

func (r *StockRepository) GetAll() ([]domain.Product, error) {
	rows, err := r.db.Query(`SELECT id, name, quantity, unit_cost FROM products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []domain.Product

	for rows.Next() {
		var stockID, quantity, cost int
		var name string
		if err := rows.Scan(&stockID, &name, &quantity, &cost); err != nil {
			return nil, err
		}

		s := domain.RestoreProduct(
			stockID, name, quantity, cost,
		)
		stocks = append(stocks, s)
	}

	return stocks, rows.Err()
}

func (r *StockRepository) GetByID(id int) (domain.Product, error) {
	var stockID, quantity, cost int
	var name string

	err := r.db.QueryRow(
		`SELECT id, name, quantity, unit_cost FROM products WHERE id = ?`,
		id,
	).Scan(&stockID, &name, &quantity, &cost)

	if errors.Is(err, sql.ErrNoRows) {
		return domain.Product{}, ErrStockNotFound
	}

	if err != nil {
		return domain.Product{}, err
	}

	return domain.RestoreProduct(stockID, name, quantity, cost), nil
}

func (r *StockRepository) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM products WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrStockNotFound
	}

	return nil
}
