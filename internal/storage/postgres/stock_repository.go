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

type StockIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewStockRepo(pg *pgxpool.Pool, c context.Context) *StockIMP {
	return &StockIMP{pgx: pg, ctx: c}

}
func (pgx *StockIMP) CreateStock(stock model.Stock) (pgconn.CommandTag, error) {
	query := "INSERT INTO stock(products_common_code,products_common_name, number_of_products,supplier_id,country,city,state,zipcode,street,latitude,longitude)" +
		" values($1, $2, $3,$4,$5,$6,$7,$8,$9,$10,$11)"
	commands, err := pgx.pgx.Exec(pgx.ctx, query,
		stock.ID, stock.ProductsCode,stock.ProductsName,stock.No_products,stock.SupplierID,
		stock.Address.Country,stock.Address.City,stock.Address.State,stock.Address.Street,stock.Address.Latitude,stock.Address.Longitude)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *StockIMP) GetStockByID(id uuid.UUID) (*model.Stock, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, "SELECT * FROM stock WHERE id = $1", id)
	stock := model.Stock{}
	err := row.Scan(&stock.ID, &stock.ProductsCode,&stock.ProductsName,&stock.No_products,&stock.SupplierID,
		&stock.Address.Country,&stock.Address.City,&stock.Address.State,&stock.Address.Zip,&stock.Address.Street,&stock.Address.Latitude,&stock.Address.Longitude,&stock.CreatedOn,&stock.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &stock, nil
}

func (pgx *StockIMP) GetStocks() ([]model.Stock, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, "SELECT * FROM stock;")
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()

	stocks := []model.Stock{}

	for rows.Next() {
		stock := model.Stock{}
		err = rows.Scan(&stock.ID, &stock.ProductsCode,&stock.ProductsName,&stock.No_products,&stock.SupplierID,
			&stock.Address.Country,&stock.Address.City,&stock.Address.State,&stock.Address.Zip,&stock.Address.Street,&stock.Address.Latitude,&stock.Address.Longitude,&stock.CreatedOn,&stock.UpdatedOn)
		if err != nil {
			fmt.Println("line 59",err)
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		stocks = append(stocks, stock)
	}
	return stocks, nil
}
func (pgx *StockIMP) UpdateStock(stock *model.Stock) (pgconn.CommandTag, error) {
	query := "UPDATE stock SET id=$1 ,products_code=$2,products_name=$3, number_of_products=$4,supplier_id=$5,country=$6,city=$7,state=$8," +
		"zipcode=$9,street=$10,latitude=$11,longitude=$12 WHERE id=$13"
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query, stock.ID, stock.ProductsCode,stock.ProductsName,stock.No_products,stock.SupplierID,
		stock.Address.Country,stock.Address.City,stock.Address.State,stock.Address.Street,stock.Address.Latitude,stock.Address.Longitude,stock.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *StockIMP) DeleteStock(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, "DELETE FROM stock WHERE id=$1", id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
