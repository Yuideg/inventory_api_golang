package initializer

import (
	//"context"
	"fmt"
	"github.com/Yideg/inventory_app/internal/glue/routing"
	"github.com/Yideg/inventory_app/internal/handler/rest"
	"github.com/Yideg/inventory_app/internal/module/usecase"
	"github.com/Yideg/inventory_app/internal/storage/postgres"
	"github.com/Yideg/inventory_app/platform/routers"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	_ "google.golang.org/grpc"
	"os"
)

func User(testInit bool) {
	DATABASE_URL := "postgres://postgres:yideg2378@localhost:5432/inventory"
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx,DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer  pool.Close()
	fmt.Println("database books connected successfuly!")
	//order
	order_repo := postgres.NewOrderRepo(pool, ctx)
	order_serv := usecase.NewOrderUsecase(order_repo)
	order_handler := rest.OrderInit(order_serv, nil)

	//product
	product_repo := postgres.NewProductRepo(pool,ctx)
	product_serv := usecase.NewProductUsecase(product_repo)
	product_handler := rest.ProductInit(product_serv, nil)

	//role
	role_repo := postgres.NewRoleRepo(pool, ctx)
	role_serv := usecase.NewRoleUsecaseServ(role_repo)
	role_handler := rest.RoleInit(role_serv, nil)
	//stock
	stock_repo := postgres.NewStockRepo(pool, ctx)
	stock_serv := usecase.NewStockUsecase(stock_repo)
	stock_handler := rest.StockInit(stock_serv, nil)

	//user
	cust_repo := postgres.NewUserRepo(pool, ctx)
	cust_serv := usecase.NewUserUsecase(cust_repo)
	cust_handler := rest.UserInit(cust_serv, role_serv,nil)

	//supplier
	supplier_repo := postgres.NewSupplierRepo(pool, ctx)
	supplier_serv := usecase.NewSupplierUsecase(supplier_repo)
	supplier_handler := rest.SupplierInit(supplier_serv, nil)
    server_router:=routing.AllRouting(cust_handler,order_handler,product_handler,role_handler,stock_handler,supplier_handler)


	server := routers.NewRouting("localhost", "9780", server_router)
	if testInit {
		fmt.Println("Initialize test mode Finished!")
		os.Exit(0)
	}
	server.Serve()

}
