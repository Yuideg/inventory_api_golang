package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type ProductUsecaseIMP struct {
	repo module.ProductRepository
}

func  NewProductUsecase(r module.ProductRepository) module.ProductUsecase {
	return &ProductUsecaseIMP{repo: r}

}

func (s *ProductUsecaseIMP) GetProducts() ([]model.Product, error) {
	products, errs := s.repo.GetProducts()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return products, nil
}

func (s *ProductUsecaseIMP) GetProductsByID(id uuid.UUID) (*model.Product, error) {
	product, errs := s.repo.GetProductsByID(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  product, nil
}

func (s *ProductUsecaseIMP) CreateProduct(product model.Product) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.CreateProduct(product)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *ProductUsecaseIMP) UpdateProduct(product *model.Product) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.UpdateProduct(product)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *ProductUsecaseIMP) DeleteProduct(id uuid.UUID) error {
	errs := s.repo.DeleteProduct(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
