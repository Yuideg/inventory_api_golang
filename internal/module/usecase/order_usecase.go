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

func (s *OrderUsecaseIMP) GetOrders() ([]model.Order, error) {
	orders, errs := s.repo.GetOrders()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return orders, nil
}

func (s *OrderUsecaseIMP) GetOrderByID(id uuid.UUID) (*model.Order, error) {
	order, errs := s.repo.GetOrderByID(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  order, nil
}

func (s *OrderUsecaseIMP) CreateOrder(order model.Order) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.CreateOrder(order)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *OrderUsecaseIMP) UpdateOrder(order *model.Order) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.UpdateOrder(order)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *OrderUsecaseIMP) DeleteOrder(id uuid.UUID) error {
	errs := s.repo.DeleteOrder(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
