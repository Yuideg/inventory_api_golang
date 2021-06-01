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

func (s *SupplierUsecaseMP) GetSupplierBySupplierID(supplierId uuid.UUID) ([]model.Supplier, error) {
	suppliers, errs := s.repo.GetSupplierBySupplierID(supplierId)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return suppliers, nil
}

func NewSupplierUsecase(r module.SupplierRepository) module.SupplierUsecase {
	return &SupplierUsecaseMP{repo: r}

}

func (s *SupplierUsecaseMP) GetSuppliers() ([]model.Supplier, error) {
	roles, errs := s.repo.GetSuppliers()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return roles, nil
}

func (s *SupplierUsecaseMP) GetSupplierByID(id uuid.UUID) (*model.Supplier, error) {
	role, errs := s.repo.GetSupplierByID(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}
func (s *SupplierUsecaseMP) GetOrderBySupplierID(supplierId uuid.UUID) ([]model.Supplier, error) {
	panic("implement me")
}

func (s *SupplierUsecaseMP) CreateSupplier(sup model.Supplier) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.CreateSupplier(sup)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *SupplierUsecaseMP) UpdateSupplier(sup *model.Supplier) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.UpdateSupplier(sup)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *SupplierUsecaseMP) DeleteSupplier(id uuid.UUID) error {
	errs := s.repo.DeleteSupplier(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
