package usecase

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

type StaffUsecaseIMP struct {
	repo module.StaffRepository
}

func  NewStaffUsecase(r module.StaffRepository) module.StaffUsecase {
	return &StaffUsecaseIMP{repo: r}

}

func (s *StaffUsecaseIMP) Get() ([]model.Staff, error) {
	staffs, errs := s.repo.Get()

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return staffs, nil
}

func (s *StaffUsecaseIMP) GetById(id uuid.UUID) (*model.Staff, error) {
	role, errs := s.repo.GetById(id)

	if errs != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(errs.Error())
	}

	return  role, nil
}

func (s *StaffUsecaseIMP) Create(staff model.Staff) (pgconn.CommandTag, error) {
	pgcom, errs := s.repo.Create(staff)
	if errs != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(errs.Error())
	}

	return pgcom, nil
}

func (s *StaffUsecaseIMP) Update(staff *model.Staff) (pgconn.CommandTag, error) {
	commadtag, errs := s.repo.Update(staff)

	if errs != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(errs.Error())
	}
	return commadtag, nil
}

func (s *StaffUsecaseIMP) Delete(id uuid.UUID) error {
	errs := s.repo.Delete(id)

	if errs != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(errs.Error())
	}

	return nil
}
