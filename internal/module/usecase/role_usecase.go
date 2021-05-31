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

func (s *RoleUsercaseIMP) GetRoles() ([]model.Role, error) {
	roles, errs := s.repo.GetRoles()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return roles, nil
}

func (s *RoleUsercaseIMP) GetRoleByID(id uuid.UUID ) (*model.Role, error) {
	role, errs := s.repo.GetRoleByID(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}

func (s *RoleUsercaseIMP) CreateRole(role model.Role) (pgconn.CommandTag, error) {
	commands, errs := s.repo.CreateRole(role)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return commands, nil
}

func (s *RoleUsercaseIMP) UpdateRole(role *model.Role) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.UpdateRole(role)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *RoleUsercaseIMP) DeleteRole(id uuid.UUID) error {
	errs := s.repo.DeleteRole(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
