package postgres

import (
	"context"
	"fmt"
	"github.com/Yideg/inventory_app/internal/constant/model"
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
func (pgx *RolePgxIMP) Create(role model.Role) (pgconn.CommandTag, error) {
	query := "INSERT INTO role (id,name) values($1, $2)"
	commands, err := pgx.pgx.Exec(pgx.ctx, query, role.ID, role.Name)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *RolePgxIMP) GetById(id uuid.UUID) (*model.Role, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, "SELECT * FROM role WHERE id = $1", id)
	fmt.Println("id =", id)
	role := model.Role{}
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}

	return &role, nil
}

func (pgx *RolePgxIMP) Get() ([]model.Role, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, "SELECT * FROM role;")
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

func (pgx *RolePgxIMP) Update(role *model.Role) (pgconn.CommandTag, error) {
	query := "UPDATE role SET id=$1, id=$2 WHERE id=$3"

	tag_command, err := pgx.pgx.Exec(pgx.ctx, query, role.ID, role.Name,role.ID)
	if err != nil {
		return nil,pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}


func (pgx *RolePgxIMP) Delete(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, "DELETE FROM role WHERE id=$1", id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
