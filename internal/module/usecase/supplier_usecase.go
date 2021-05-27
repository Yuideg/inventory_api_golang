package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type SupplierUsecaseMP struct {
	repo module.SupplierRepository
}

func NewSupplierUsecase(r module.SupplierRepository) module.SupplierUsecase {
	return &SupplierUsecaseMP{repo: r}

}

func (s *SupplierUsecaseMP) Get() ([]model.Supplier, error) {
	roles, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return roles, nil
}

func (s *SupplierUsecaseMP) GetById(id uuid.UUID) (*model.Supplier, error) {
	role, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}

func (s *SupplierUsecaseMP) Create(sup model.Supplier) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.Create(sup)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *SupplierUsecaseMP) Update(sup *model.Supplier) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(sup)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *SupplierUsecaseMP) Delete(id uuid.UUID) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
