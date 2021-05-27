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

type CustomerIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewCustomerRepo(pg *pgxpool.Pool, c context.Context) *CustomerIMP {
	return &CustomerIMP{pgx: pg, ctx: c}

}
func (pgx *CustomerIMP) Create(custom model.Customer) (pgconn.CommandTag, error) {
	query := "INSERT INTO customer(username,password,first_name,last_name,gender,email,country,city,state,zipcode," +
		"street,latitude,longitude,phone,role_id)" +
		" values($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13,$14,$15)"
	commands, err := pgx.pgx.Exec(pgx.ctx, query,
		 custom.Username, custom.Password, custom.FirstName,
		custom.LastName,custom.Gender,custom.Email,custom.Address.Country,custom.Address.City,custom.Address.State,custom.Address.Zip,
		custom.Address.Street,custom.Address.Latitude,custom.Address.Longitude,custom.Phone,custom.RoleID)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}

func (pgx *CustomerIMP) GetById(id uuid.UUID) (*model.Customer, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, "SELECT * FROM customer WHERE id = $1", id)
	custom := model.Customer{}
	err := row.Scan(&custom.ID, &custom.Username, &custom.Password, &custom.FirstName,
		&custom.LastName,&custom.Gender,&custom.Email,&custom.Address.Country,&custom.Address.City,&custom.Address.State,&custom.Address.Zip,
		&custom.Address.Street,&custom.Address.Latitude,&custom.Address.Longitude,&custom.Phone,&custom.RoleID,&custom.CreatedOn,&custom.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &custom, nil
}

func (pgx *CustomerIMP) Get() ([]model.Customer, error) {
      fmt.Println("cust repo get started")
	rows, err := pgx.pgx.Query(pgx.ctx, "SELECT * FROM customer;")
	if err != nil {
		fmt.Println("line 53",err)
		return nil,pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()
    fmt.Println("row obj ",rows)
	customers := []model.Customer{}
	for rows.Next() {
		custom := model.Customer{}
		err = rows.Scan(&custom.ID, &custom.Username, &custom.Password, &custom.FirstName,
			&custom.LastName,&custom.Gender,&custom.Email,&custom.Address.Country,&custom.Address.City,&custom.Address.State,&custom.Address.Zip,
			&custom.Address.Street,&custom.Address.Latitude,&custom.Address.Longitude,&custom.Phone,&custom.RoleID,&custom.CreatedOn,&custom.UpdatedOn)
		if err != nil {
			fmt.Println("line 65",err)
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		customers = append(customers, custom)
	}
	fmt.Println("data repo -",customers)
	return customers, nil
}
func (pgx *CustomerIMP) Update(custom *model.Customer) (pgconn.CommandTag, error) {
	query := "UPDATE customer SET id=$1, username=$2,password=$3,first_name=$4,last_name=$5," +
		"gender=$6,email=$7,country=$8,city=$9,state=$10,zipcode=$11,street=$12,latitude=$13,longitude=$14,phone=$15,role_id=$16 WHERE id=$17"

	tag_command, err := pgx.pgx.Exec(pgx.ctx, query, custom.ID, custom.Username, custom.Password, custom.FirstName,
		custom.LastName,custom.Gender,custom.Email,custom.Address.Country,custom.Address.City,custom.Address.State,custom.Address.Zip,custom.Address.Street,
		custom.Address.Latitude,custom.Address.Longitude,custom.Phone,custom.RoleID,custom.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *CustomerIMP) Delete(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, "DELETE FROM customer WHERE id=$1", id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
