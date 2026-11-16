package service

import (
	"stockgo/internal/domain"
)

type ReportRepository interface {
	GetRevenueToday() (int, error)
	GetRevenueThisMonth() (int, error)
	GetTotalRevenue() (int, error)
	CountSalesToday() (int, error)
	CountSalesThisMonth() (int, error)
	GetInventoryValue() (int, error)
	CountProducts() (int, error)
	GetProfitThisMonth() (int, error)
}

type ReportService struct {
	reportRepo ReportRepository
}

func NewReportService(reportRepo ReportRepository) *ReportService {
	return &ReportService{reportRepo: reportRepo}
}

type Dashboard struct {
	InventoryValue int
	ProductCount   int
	SalesToday     int
	RevenueToday   int
	RevenueMonth   int
	ProfitMonth    int
	LowStock       []domain.Product
	//RecentSales []SaleItem
}

func (s *ReportService) GetDashboard() (Dashboard, error) {
	var d Dashboard
	var err error

	if d.RevenueToday, err = s.reportRepo.GetRevenueToday(); err != nil {
		return Dashboard{}, err
	}

	if d.RevenueMonth, err = s.reportRepo.GetRevenueThisMonth(); err != nil {

		return Dashboard{}, err
	}

	if d.SalesToday, err = s.reportRepo.CountSalesToday(); err != nil {
		return Dashboard{}, err
	}

	if d.InventoryValue, err = s.reportRepo.GetInventoryValue(); err != nil {
		return Dashboard{}, err
	}

	if d.ProductCount, err = s.reportRepo.CountProducts(); err != nil {
		return Dashboard{}, err
	}

	if d.ProfitMonth, err = s.reportRepo.GetProfitThisMonth(); err != nil {
		return Dashboard{}, err
	}

	return d, nil
}

func (s *ReportService) GetRevenueToday() (int, error) {
	return s.reportRepo.GetRevenueToday()
}

func (s *ReportService) GetRevenueThisMonth() (int, error) {
	return s.reportRepo.GetRevenueThisMonth()
}

func (s *ReportService) GetTotalRevenue() (int, error) {
	return s.reportRepo.GetTotalRevenue()
}

func (s *ReportService) CountSalesToday() (int, error) {
	return s.reportRepo.CountSalesToday()
}

func (s *ReportService) GetSalesThisMonth() (int, error) {
	return s.reportRepo.CountSalesThisMonth()
}
