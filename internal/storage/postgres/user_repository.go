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

type UserMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewUserRepo(pg *pgxpool.Pool, c context.Context) *UserMP {
	return &UserMP{pgx: pg, ctx: c}

}
func (pgx *UserMP) CreateUser(user model.User) (pgconn.CommandTag, error) {
	query := "INSERT INTO user_tb(username,password,first_name,last_name,gender,email,country,city,state,zipcode,street,latitude,longitude,phone,role_id)" +
		" values($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13,$14,$15)"
	commands, err := pgx.pgx.Exec(pgx.ctx, query,
		 user.Username, user.Password, user.FirstName,
		user.LastName, user.Gender, user.Email, user.Address.Country, user.Address.City, user.Address.State, user.Address.Zip,
		user.Address.Street, user.Address.Latitude, user.Address.Longitude, user.Phone, user.RoleID)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}

func (pgx *UserMP) GetUserByID(id uuid.UUID) (*model.User, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, "SELECT * FROM user_tb WHERE id = $1", id)
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
      fmt.Println("cust repo get started")
	rows, err := pgx.pgx.Query(pgx.ctx, "SELECT * FROM user_tb;")
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
	query := "UPDATE user_tb SET id=$1, username=$2,password=$3,first_name=$4,last_name=$5," +
		"gender=$6,email=$7,country=$8,city=$9,state=$10,zipcode=$11,street=$12,latitude=$13,longitude=$14,phone=$15,role_id=$16 WHERE id=$17"

	tag_command, err := pgx.pgx.Exec(pgx.ctx, query, user.ID, user.Username, user.Password, user.FirstName,
		user.LastName, user.Gender, user.Email, user.Address.Country, user.Address.City, user.Address.State, user.Address.Zip, user.Address.Street,
		user.Address.Latitude, user.Address.Longitude, user.Phone, user.RoleID, user.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *UserMP) DeleteUser(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, "DELETE FROM user_tb WHERE id=$1", id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
