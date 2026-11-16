package repository

import (
	"database/sql"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepo(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// Consider passing periods

func (r *ReportRepository) GetRevenueToday() (int, error) {
	var total int

	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(unit_price), 0)
		FROM sales
		WHERE DATE(created_at) = DATE('now', 'localtime');
	`).Scan(&total)

	return total, err
}

func (r *ReportRepository) GetRevenueThisMonth() (int, error) {
	var total int

	err := r.db.QueryRow(`
		SELECT COALESCE(SUM(unit_price),0)
		FROM sales
		WHERE strftime('%Y-%m', created_at)
		      = strftime('%Y-%m','now','localtime');
	`).Scan(&total)

	return total, err
}

func (r *ReportRepository) GetTotalRevenue() (int, error) {
	var total int

	err := r.db.QueryRow(
		`SELECT COALESCE(SUM(unit_price), 0) FROM sales;`,
	).Scan(&total)

	return total, err
}

func (r *ReportRepository) CountSalesToday() (int, error) {
	var count int

	err := r.db.QueryRow(`
		SELECT COUNT(*)
		FROM sales
		WHERE DATE(created_at) = DATE('now','localtime');
	`).Scan(&count)

	return count, err
}

func (r *ReportRepository) CountSalesThisMonth() (int, error) {
	var count int

	err := r.db.QueryRow(
		`
		SELECT COUNT(*)
		FROM sales
		WHERE strftime('%Y-%m', sales.created_at)
		      = strftime('%Y-%m','now','localtime');
		`,
	).Scan(&count)

	return count, err
}

func (r *ReportRepository) GetInventoryValue() (int, error) {
	var total int

	err := r.db.QueryRow(`
		SELECT COALESCE(
			SUM(quantity * unit_cost),
			0
		)
		FROM products;
	`).Scan(&total)

	return total, err
}

func (r *ReportRepository) CountProducts() (int, error) {
	var count int

	err := r.db.QueryRow(`
		SELECT COUNT(*)
		FROM products;
	`).Scan(&count)

	return count, err
}

func (r *ReportRepository) GetProfitThisMonth() (int, error) {
	var profit int

	err := r.db.QueryRow(`
		SELECT COALESCE(
		    SUM(unit_price - unit_cost),
		    0
		)
		FROM sales

		WHERE strftime('%Y-%m', created_at)
		      = strftime('%Y-%m','now','localtime');
	`).Scan(&profit)

	return profit, err
}
