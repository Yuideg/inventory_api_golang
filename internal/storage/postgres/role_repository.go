package postgres

import (
	"context"
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/constant/query"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

type RolePgxIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewRoleRepo(pg *pgxpool.Pool, c context.Context) *RolePgxIMP {
	return &RolePgxIMP{pgx: pg, ctx: c}

}
func (pgx *RolePgxIMP) CreateRole(role model.Role) (pgconn.CommandTag, error) {
	commands, err := pgx.pgx.Exec(pgx.ctx, query.RoleInsert, role.ID, role.Name)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *RolePgxIMP) GetRoleByID(id uuid.UUID) (*model.Role, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, query.RoleSelectOne, id)
	fmt.Println("id =", id)
	role := model.Role{}
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}

	return &role, nil
}

func (pgx *RolePgxIMP) GetRoles() ([]model.Role, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, query.RoleSelectAll)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()
	roles := []model.Role{}

	for rows.Next() {
		role := model.Role{}
		err = rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		roles = append(roles, role)
	}
	return roles, nil
}

func (pgx *RolePgxIMP) UpdateRole(role *model.Role) (pgconn.CommandTag, error) {
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query.RoleUpdate, role.ID, role.Name,role.ID)
	if err != nil {
		return nil,pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}


func (pgx *RolePgxIMP) DeleteRole(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, query.RoleDelete, id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
