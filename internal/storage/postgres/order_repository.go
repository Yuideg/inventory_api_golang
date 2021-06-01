package postgres

import (
	"context"
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/constant/query"
	pkg "github.com/Yideg/inventory_app/pkg/error"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
	uuid "github.com/satori/go.uuid"
)

type OrderMP struct {
	pgx *pgxpool.Pool
	ctx context.Context
}

func NewOrderRepo(pg *pgxpool.Pool, c context.Context) *OrderMP {
	return &OrderMP{pgx: pg, ctx: c}

}
func (pgx *OrderMP) CreateOrder(order model.Order) (pgconn.CommandTag, error) {

	commands, err := pgx.pgx.Exec(pgx.ctx, query.OrderInsert,order.Quantity, order.Unit,
		order.CertifiedBy,order.Customer_ID,order.Product_ID,order.Status,order.ExpiredOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseCreate.FetchErrors(err.Error())
	}

	return commands, nil
}
func (pgx *OrderMP) GetOrderByID(id uuid.UUID) (*model.Order, error) {
	row := pgx.pgx.QueryRow(pgx.ctx, query.OrderSelectOneBy_ID, id)
	order := model.Order{}
	err := row.Scan(&order.ID, &order.Quantity, &order.Unit,
		&order.CertifiedBy,&order.Customer_ID,&order.Product_ID,&order.Status,&order.ExpiredOn,&order.CreatedOn,&order.UpdatedOn)
	if err != nil {
		return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	return &order, nil
}
func (pgx *OrderMP) GetOrderByUserID(user_id uuid.UUID) ([]model.Order, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, query.OrderSelectBy_User_ID,user_id)
	if err != nil {
		return nil,pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()

	orders := []model.Order{}

	for rows.Next() {
		order := model. Order{}
		err = rows.Scan(&order.ID, &order.Quantity, &order.Unit,
			&order.CertifiedBy,&order.Customer_ID,&order.Product_ID,&order.Status,&order.ExpiredOn,&order.CreatedOn,&order.UpdatedOn)
		if err != nil {
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		orders = append(orders, order)
	}
	return orders, nil
}
func (pgx *OrderMP) GetOrders() ([]model.Order, error) {
	rows, err := pgx.pgx.Query(pgx.ctx, query.OrderSelectAll)
	if err != nil {
		return nil,pkg.ErrorDatabaseGet.FetchErrors(err.Error())
	}
	defer rows.Close()

	orders := []model.Order{}

	for rows.Next() {
		order := model. Order{}
		err = rows.Scan(&order.ID, &order.Quantity, &order.Unit,
			&order.CertifiedBy,&order.Customer_ID,&order.Product_ID,&order.Status,&order.ExpiredOn,&order.CreatedOn,&order.UpdatedOn)
		if err != nil {
			return nil, pkg.ErrorDatabaseGet.FetchErrors(err.Error())
		}

		orders = append(orders, order)
	}
	return orders, nil
}
func (pgx *OrderMP) UpdateOrder(order *model.Order) (pgconn.CommandTag, error) {
	tag_command, err := pgx.pgx.Exec(pgx.ctx, query.OrderUpdate, order.ID, order.Quantity, order.Unit,
		order.CertifiedBy,order.Customer_ID,order.Product_ID,order.Status,order.ExpiredOn,order.ID)
	if err != nil {
		return nil, pkg.ErrorDatabaseUpdate.FetchErrors(err.Error())
	}
	return tag_command, nil
}

func (pgx *OrderMP) DeleteOrder(id uuid.UUID) error {
	_, err := pgx.pgx.Exec(pgx.ctx, query.OrderDelete, id)
	if err != nil {
		return pkg.ErrorDatabaseDelete.FetchErrors(err.Error())
	}
	return nil
}
