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

type UserMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewUserRepo(pg *pgxpool.Pool, c context.Context) *UserMP {
	return &UserMP{pgx: pg, ctx: c}

}
func (pgx *UserMP) CreateUser(user model.User) (pgconn.CommandTag, error) {
	commands, err := pgx.pgx.Exec(pgx.ctx, query.UserInsert,
		 user.Username, user.Password, user.FirstName,
		user.LastName, user.Gender, user.Email, user.Address.Country, user.Address.City, user.Address.State, user.Address.Zip,
		user.Address.Street, user.Address.Latitude, user.Address.Longitude, user.Phone, user.RoleID)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}

func (pgx *UserMP) GetUserByID(id uuid.UUID) (*model.User, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, query.UserSelectOne, id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.FirstName,
		&user.LastName,&user.Gender,&user.Email,&user.Address.Country,&user.Address.City,&user.Address.State,&user.Address.Zip,
		&user.Address.Street,&user.Address.Latitude,&user.Address.Longitude,&user.Phone,&user.RoleID,&user.CreatedOn,&user.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &user, nil
}

func (pgx *UserMP) GetUsers() ([]model.User, error) {
      fmt.Println("user repo get started")
	rows, err := pgx.pgx.Query(pgx.ctx, query.UserSelectAll)
	if err != nil {
		fmt.Println("line 53",err)
		return nil,pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()
    fmt.Println("row obj ",rows)
	users := []model.User{}
	for rows.Next() {
		custom := model.User{}
		err = rows.Scan(&custom.ID, &custom.Username, &custom.Password, &custom.FirstName,
			&custom.LastName,&custom.Gender,&custom.Email,&custom.Address.Country,&custom.Address.City,&custom.Address.State,&custom.Address.Zip,
			&custom.Address.Street,&custom.Address.Latitude,&custom.Address.Longitude,&custom.Phone,&custom.RoleID,&custom.CreatedOn,&custom.UpdatedOn)
		if err != nil {
			fmt.Println("line 65",err)
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		users = append(users, custom)
	}
	fmt.Println("data repo -", users)
	return users, nil
}
func (pgx *UserMP) UpdateUser(user *model.User) (pgconn.CommandTag, error) {

	tag_command, err := pgx.pgx.Exec(pgx.ctx, query.UserUpdate, user.ID, user.Username, user.Password, user.FirstName,
		user.LastName, user.Gender, user.Email, user.Address.Country, user.Address.City, user.Address.State, user.Address.Zip, user.Address.Street,
		user.Address.Latitude, user.Address.Longitude, user.Phone, user.RoleID, user.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *UserMP) DeleteUser(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, query.UserDelete, id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
