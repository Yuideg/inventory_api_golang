package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type UserUsecaseIMP struct {
	repo module.UserRepository
}

func NewUserUsecase(r module.UserRepository) module.UserUsecase {
	return &UserUsecaseIMP{repo: r}

}

func (s *UserUsecaseIMP) GetUsers() ([]model.User, error) {
	customers, errs := s.repo.GetUsers()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return customers, nil
}

func (s *UserUsecaseIMP) GetUserByID(id uuid.UUID) (*model.User, error) {
	customer, errs := s.repo.GetUserByID(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  customer, nil
}

func (s *UserUsecaseIMP) CreateUser(customer model.User) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.CreateUser(customer)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *UserUsecaseIMP) UpdateUser(customer *model.User) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.UpdateUser(customer)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *UserUsecaseIMP) DeleteUser(id uuid.UUID) error {
	errs := s.repo.DeleteUser(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
