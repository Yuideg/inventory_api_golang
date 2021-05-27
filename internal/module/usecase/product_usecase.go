package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
)

type ProductUsecaseIMP struct {
	repo module.ProductRepository
}

func  NewProductUsecase(r module.ProductRepository) module.ProductUsecase {
	return &ProductUsecaseIMP{repo: r}

}

func (s *ProductUsecaseIMP) Get() ([]model.Product, error) {
	products, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return products, nil
}

func (s *ProductUsecaseIMP) GetById(id string) (*model.Product, error) {
	product, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  product, nil
}

func (s *ProductUsecaseIMP) Create(product model.Product) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.Create(product)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *ProductUsecaseIMP) Update(product *model.Product) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(product)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *ProductUsecaseIMP) Delete(id string) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
