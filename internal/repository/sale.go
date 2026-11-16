package repository

import (
	"database/sql"
	"stockgo/internal/domain"
	"time"
)

type SaleRepository struct {
	db *sql.DB
}

func NewSaleRepo(db *sql.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

func (r *SaleRepository) Insert(sale *domain.Sale) error {
	res, err := r.db.Exec(`
		INSERT INTO sales (
    		product_id,
    		product_name,
    		unit_cost,
    		unit_price
       	) VALUES (?, ?, ?, ?)`,
		sale.StockID(), sale.ProductName(), sale.UnitCost(), sale.SoldPrice(),
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	sale.AssignID(int(id))
	return nil
}

func (r *SaleRepository) GetAll() ([]domain.Sale, error) {
	rows, err := r.db.Query(
		`SELECT id, product_id, product_name, unit_cost, unit_price, created_at FROM sales ORDER BY created_at DESC, id DESC`,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sales []domain.Sale

	for rows.Next() {
		var (
			id, stockID, unitCost, soldPrice int
			productName                      string
			createdAt                        time.Time
		)

		if err := rows.Scan(&id, &stockID, &productName, &unitCost, &soldPrice, &createdAt); err != nil {
			return nil, err
		}

		sales = append(sales, domain.RestoreSale(id, stockID, productName, unitCost, soldPrice, createdAt))
	}

	return sales, rows.Err()
}

func (r *SaleRepository) GetByStockID(stockID int) ([]domain.Sale, error) {
	rows, err := r.db.Query(`
		SELECT
    		id,
    		product_id,
    		product_name,
    		unit_cost,
    		unit_price,
    		created_at
		FROM sales
		WHERE product_id = ?
		ORDER BY created_at DESC, id DESC`,
		stockID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sales []domain.Sale

	for rows.Next() {
		var (
			id, sid, cost, price int
			name                 string
			createdAt            time.Time
		)

		if err := rows.Scan(&id, &sid, &name, &cost, &price, &createdAt); err != nil {
			return nil, err
		}

		sales = append(sales, domain.RestoreSale(id, sid, name, cost, price, createdAt))
	}

	return sales, rows.Err()
}
