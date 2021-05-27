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

type StaffIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewStaffRepo(pg *pgxpool.Pool, c context.Context) *StaffIMP {
	return &StaffIMP{pgx: pg, ctx: c}

}
func (pgx *StaffIMP) Create(staff model.Staff) (pgconn.CommandTag, error) {
	query := "INSERT INTO staff(username,password,first_name,last_name,gender,email,country,city,state,zipcode," +
		"street,latitude,longitude,phone,role_id,salary) values($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13,$14,$15,$16,$17)"
	commands, err := pgx.pgx.Exec(pgx.ctx, query, staff.ID,staff.Username,staff.Password,staff.FirstName,staff.LastName,staff.Gender,
		staff.Email,staff.Address.Country,staff.Address.City,staff.Address.State,staff.Address.Zip,staff.Address.Street,staff.Address.Latitude,
		staff.Address.Longitude,staff.Phone,staff.RoleID,staff.Salary)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *StaffIMP) GetById(id uuid.UUID) (*model.Staff, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, "SELECT * FROM staff WHERE id = $1", id)
	staff := model.Staff{}
	err := row.Scan(&staff.ID,&staff.Username,&staff.Password,&staff.FirstName,&staff.LastName,&staff.Gender,
		&staff.Email,&staff.Address.Country,&staff.Address.City,&staff.Address.State,&staff.Address.Zip,&staff.Address.Street,&staff.Address.Latitude,&staff.Address.Longitude,
		&staff.Phone,&staff.RoleID,&staff.Salary,&staff.CreatedOn,&staff.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &staff, nil
}

func (pgx *StaffIMP) Get() ([]model.Staff, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, "SELECT * FROM staff;")
	if err != nil {
		return nil,pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()

	staffs := []model.Staff{}

	for rows.Next() {
		staff := model.Staff{}
		err = rows.Scan(&staff.ID,&staff.Username,&staff.Password,&staff.FirstName,&staff.LastName,&staff.Gender,
			&staff.Email,&staff.Address.Country,&staff.Address.City,&staff.Address.State,&staff.Address.Zip,
			&staff.Address.Street,&staff.Address.Latitude,&staff.Address.Longitude,
			&staff.Phone,&staff.RoleID,&staff.Salary,&staff.CreatedOn,&staff.UpdatedOn)
		if err != nil {
			panic(err)
			fmt.Println("line 61 ",err)
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		staffs = append(staffs, staff)
	}
	return staffs, nil
}
func (pgx *StaffIMP) Update(staff *model.Staff) (pgconn.CommandTag, error) {
	query := "UPDATE staff SET id=$1, username=$2,password=$3,first_name=$4,last_name=$5," +
		"gender=$6,email=$7,country=$8,city=$9,state=$10,zipcode=$11,street=$12,latitude=$13,longitude=$14,phone=$15,role_id=$16,salary=$17 " +
		"WHERE id=$18"
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query, staff.ID,staff.Username,staff.Password,staff.FirstName,staff.LastName,staff.Gender,
		staff.Email,staff.Address.Country,staff.Address.City,staff.Address.State,staff.Address.Zip,staff.Address.Street,staff.Address.Latitude,
		staff.Address.Longitude,staff.Phone,staff.RoleID,staff.Salary,staff.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *StaffIMP) Delete(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, "DELETE FROM staff WHERE id=$1", id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
