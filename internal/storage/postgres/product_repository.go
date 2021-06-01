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

type ProductIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewProductRepo(pg *pgxpool.Pool, c context.Context) *ProductIMP {
	return &ProductIMP{pgx: pg, ctx: c}

}
func (pgx *ProductIMP) CreateProduct(product model.Product) (pgconn.CommandTag, error) {
	commands, err := pgx.pgx.Exec(pgx.ctx, query.ProductInsert,product.ID,product.Barcode,product.ProductName,
		product.StockId,product.Brand, product.Image,product.Cost, product.Price,product.Unit,
		product.Tax,product.Discount,product.MfgDate,product.ExpiredOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *ProductIMP) GetProductsByID(id uuid.UUID) (*model.Product, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, query.ProductSelectOne, id)
	product := model.Product{}
	err := row.Scan(&product.ID,&product.Barcode,&product.ProductName,
		&product.StockId,&product.Brand, &product.Image,&product.Cost, &product.Price,&product.Unit,
		&product.Tax,&product.Discount,&product.MfgDate,&product.ExpiredOn,&product.CreatedOn,&product.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &product, nil
}

func (pgx *ProductIMP) GetProducts() ([]model.Product, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, query.ProductSelectAll)
	if err != nil {
		return nil,pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()

	products := []model.Product{}

	for rows.Next() {
		product := model.Product{}
		err = rows.Scan(&product.ID,&product.Barcode,&product.ProductName,
			&product.StockId,&product.Brand, &product.Image,&product.Cost, &product.Price,&product.Unit,
			&product.Tax,&product.Discount,&product.MfgDate,&product.ExpiredOn,&product.CreatedOn,&product.UpdatedOn)
		if err != nil {
			fmt.Println("line 60 ",err)
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		products = append(products, product)
	}
	return products, nil
}
func (pgx *ProductIMP) UpdateProduct(product *model.Product) (pgconn.CommandTag, error) {
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query.ProductUpdate, product.ID,product.Barcode,product.ProductName,
		product.StockId,product.Brand, product.Image,product.Cost, product.Price,product.Unit,
		product.Tax,product.Discount,product.MfgDate,product.ExpiredOn,product.Barcode)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *ProductIMP) DeleteProduct(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, query.ProductDelete, id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
