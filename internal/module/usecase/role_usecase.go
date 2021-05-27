package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type RoleUsercaseIMP struct {
	repo module.RoleRepository
}

func NewRoleUsecaseServ(r module.RoleRepository) module.RoleServices {
	return &RoleUsercaseIMP{repo: r}

}

func (s *RoleUsercaseIMP) Get() ([]model.Role, error) {
	roles, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return roles, nil
}

func (s *RoleUsercaseIMP) GetById(id uuid.UUID ) (*model.Role, error) {
	role, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}

func (s *RoleUsercaseIMP) Create(role model.Role) (pgconn.CommandTag, error) {
	commands, errs := s.repo.Create(role)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return commands, nil
}

func (s *RoleUsercaseIMP) Update(role *model.Role) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(role)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *RoleUsercaseIMP) Delete(id uuid.UUID) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
