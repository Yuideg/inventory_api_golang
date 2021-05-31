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

type SupplierIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewSupplierRepo(pg *pgxpool.Pool, c context.Context) *SupplierIMP {
	return &SupplierIMP{pgx: pg, ctx: c}

}
func (pgx *SupplierIMP) CreateSupplier(supplier model.Supplier) (pgconn.CommandTag, error) {
	query := "INSERT INTO supplier(supplier_name,country,city,state,zipcode,street,latitude,longitude,fax,po_box,email,website,phone)" +
		" values($1, $2, $3,$4,$5,$6,$7, $8, $9,$10,$11,$12,$13)"
	commands, err := pgx.pgx.Exec(pgx.ctx, query,
		supplier.Name,supplier.Address.Country,supplier.Address.City,supplier.Address.State,supplier.Address.Zip,supplier.Address.Street,supplier.Address.Latitude,supplier.Address.Longitude,
		supplier.Fax, supplier.PoBox,supplier.Email, supplier.WebSite, supplier.Phone)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *SupplierIMP) GetSupplierByID(id uuid.UUID) (*model.Supplier, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, "SELECT * FROM supplier WHERE id = $1", id)
	supplier := model.Supplier{}
	err := row.Scan(&supplier.ID,&supplier.Name,&supplier.Address.Country,&supplier.Address.City,&supplier.Address.State,&supplier.Address.Zip,&supplier.Address.Street,&supplier.Address.Latitude,&supplier.Address.Longitude,
		&supplier.Fax, &supplier.PoBox,&supplier.Email, &supplier.WebSite, &supplier.Phone,&supplier.CreatedOn,&supplier.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &supplier, nil
}

func (pgx *SupplierIMP) GetSuppliers() ([]model.Supplier, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, "SELECT * FROM supplier;")
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()

	suppliers := []model.Supplier{}

	for rows.Next() {
		supplier := model.Supplier{}
		err = rows.Scan(&supplier.ID,&supplier.Name,&supplier.Address.Country,&supplier.Address.City,&supplier.Address.State,&supplier.Address.Zip,&supplier.Address.Street,&supplier.Address.Latitude,&supplier.Address.Longitude,
			&supplier.Fax, &supplier.PoBox,&supplier.Email, &supplier.WebSite, &supplier.Phone,&supplier.CreatedOn,&supplier.UpdatedOn)
		if err != nil {
			fmt.Println("line 59",err)
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		suppliers = append(suppliers, supplier)
	}
	return suppliers, nil
}
func (pgx *SupplierIMP) UpdateSupplier(supplier *model.Supplier) (pgconn.CommandTag, error) {
	query := "UPDATE supplier SET id=$1 ,supplier_name=$2,country=$3,city=$4,state=$5,zipcode=$6,street=$7,latitude=$8,longitude=$9,fax=$10," +
		"po_box=$11,email=$12,website=$13,phone=$4 WHERE id=$15"
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query, supplier.ID,supplier.ID, supplier.Name,supplier.Address.Country,supplier.Address.City,supplier.Address.State,supplier.Address.Zip,supplier.Address.Street,supplier.Address.Latitude,supplier.Address.Longitude,
		supplier.Fax, supplier.PoBox,supplier.Email, supplier.WebSite, supplier.Phone,supplier.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *SupplierIMP) DeleteSupplier(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, "DELETE FROM supplier WHERE id=$1", id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
