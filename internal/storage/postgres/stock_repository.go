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

type StockIMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewStockRepo(pg *pgxpool.Pool, c context.Context) *StockIMP {
	return &StockIMP{pgx: pg, ctx: c}

}
func (pgx *StockIMP) CreateStock(stock model.Stock) (pgconn.CommandTag, error) {
	commands, err := pgx.pgx.Exec(pgx.ctx, query.StockInsert,
		stock.ID, stock.ProductsCode,stock.ProductsName,stock.No_products,stock.SupplierID,
		stock.Address.Country,stock.Address.City,stock.Address.State,stock.Address.Street,stock.Address.Latitude,stock.Address.Longitude)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *StockIMP) GetStockByID(id uuid.UUID) (*model.Stock, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, query.StockSelectOne, id)
	stock := model.Stock{}
	err := row.Scan(&stock.ID, &stock.ProductsCode,&stock.ProductsName,&stock.No_products,&stock.SupplierID,
		&stock.Address.Country,&stock.Address.City,&stock.Address.State,&stock.Address.Zip,&stock.Address.Street,&stock.Address.Latitude,&stock.Address.Longitude,&stock.CreatedOn,&stock.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &stock, nil
}

func (pgx *StockIMP) GetStocks() ([]model.Stock, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, query.StockSelectAll)
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
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query.StockUpdate, stock.ID, stock.ProductsCode,stock.ProductsName,stock.No_products,stock.SupplierID,
		stock.Address.Country,stock.Address.City,stock.Address.State,stock.Address.Street,stock.Address.Latitude,stock.Address.Longitude,stock.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *StockIMP) DeleteStock(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, query.StockDelete, id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
