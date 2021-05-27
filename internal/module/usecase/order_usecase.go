package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type OrderUsecaseIMP struct {
	repo module.OrderRepository
}

func  NewOrderUsecase(r module.OrderRepository) module.OrderUsecase {
	return &OrderUsecaseIMP{repo: r}

}

func (s *OrderUsecaseIMP) Get() ([]model.Order, error) {
	orders, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return orders, nil
}

func (s *OrderUsecaseIMP) GetById(id uuid.UUID) (*model.Order, error) {
	order, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  order, nil
}

func (s *OrderUsecaseIMP) Create(order model.Order) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.Create(order)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *OrderUsecaseIMP) Update(order *model.Order) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(order)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *OrderUsecaseIMP) Delete(id uuid.UUID) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
