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

func (s *StockUsecaseMP) Get() ([]model.Stock, error) {
	roles, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return roles, nil
}

func (s *StockUsecaseMP) GetById(id uuid.UUID) (*model.Stock, error) {
	role, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}

func (s *StockUsecaseMP) Create(stock model.Stock) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.Create(stock)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *StockUsecaseMP) Update(stock *model.Stock) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(stock)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *StockUsecaseMP) Delete(id uuid.UUID) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
