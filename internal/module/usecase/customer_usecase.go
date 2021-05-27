package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type CustomerUsecaseIMP struct {
	repo module.CustomerRepository
}

func  NewCustomerUsecase(r module.CustomerRepository) module.CustomerUsecase {
	return &CustomerUsecaseIMP{repo: r}

}

func (s *CustomerUsecaseIMP) Get() ([]model.Customer, error) {
	customers, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return customers, nil
}

func (s *CustomerUsecaseIMP) GetById(id uuid.UUID) (*model.Customer, error) {
	customer, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  customer, nil
}

func (s *CustomerUsecaseIMP) Create(customer model.Customer) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.Create(customer)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *CustomerUsecaseIMP) Update(customer *model.Customer) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(customer)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *CustomerUsecaseIMP) Delete(id uuid.UUID) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
