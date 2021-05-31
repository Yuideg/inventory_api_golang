package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type StockUsecaseMP struct {
	repo module.StockRepository
}

func  NewStockUsecase(r module.StockRepository) module.StockUsecae {
	return &StockUsecaseMP{repo: r}

}

func (s *StockUsecaseMP) GetStocks() ([]model.Stock, error) {
	stocks, errs := s.repo.GetStocks()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return stocks, nil
}

func (s *StockUsecaseMP) GetStockByID(id uuid.UUID) (*model.Stock, error) {
	role, errs := s.repo.GetStockByID(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}

func (s *StockUsecaseMP) CreateStock(stock model.Stock) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.CreateStock(stock)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *StockUsecaseMP) UpdateStock(stock *model.Stock) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.UpdateStock(stock)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *StockUsecaseMP) DeleteStock(id uuid.UUID) error {
	errs := s.repo.DeleteStock(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
