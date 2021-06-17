package main
import (
	"fmt"
	"github.com/Yideg/inventory_app/internal/grpc/order"
	"github.com/Yideg/inventory_app/internal/grpc/product"
	"github.com/Yideg/inventory_app/internal/grpc/role"
	"github.com/Yideg/inventory_app/internal/grpc/stock"
	"github.com/Yideg/inventory_app/internal/grpc/supplier"
	"github.com/Yideg/inventory_app/internal/grpc/user"
	"github.com/Yideg/inventory_app/internal/module"
	"github.com/Yideg/inventory_app/internal/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
if err := RunServer(); err != nil {
_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
os.Exit(1)
}
}

func RunServer() error {
	DATABASE_URL := "postgres://postgres:yideg2378@localhost:5432/inventory"
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx,DATABASE_URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer  pool.Close()
	fmt.Println("database books connected successfuly!")
	//repository connection
	postgresOrder := postgres.NewOrderRepo(pool,ctx)
	postgresProduct := postgres.NewProductRepo(pool,ctx)
	postgresRole := postgres.NewRoleRepo(pool,ctx)
	postgresStock := postgres.NewStockRepo(pool,ctx)
	postgresSupplier := postgres.NewSupplierRepo(pool,ctx)
	postgresUser := postgres.NewUserRepo(pool,ctx)



     //service connection
	osrv := module.OrderUsecase(postgresOrder)
	psrv := module.ProductUsecase(postgresProduct)
	rsrv := module.RoleServices(postgresRole)
	stsrv := module.StockUsecae(postgresStock)
	suppsrv := module.SupplierUsecase(postgresSupplier)
	usrsrv := module.UserUsecase(postgresUser)

	gs := grpc.NewServer()
   //new object calling
	gosrv:=order.NewOrderServer(osrv)
	gpsrv:=product.NewProductServer(psrv)
	grsrv:=role.NewRoleServer(rsrv)
	gstsrv:=stock.NewStockServer(stsrv)
	gsuppsrv:=supplier.NewSupplierServer(suppsrv)
	gusrsrv:=user.NewUserServer(usrsrv)


	order.RegisterOrderServiceServer(gs,gosrv)
	product.RegisterProductServiceServer(gs,gpsrv)
	role.RegisterRoleServiceServer(gs,grsrv)
	stock.RegisterStockServiceServer(gs,gstsrv)
	supplier.RegisterSupplierServiceServer(gs,gsuppsrv)
	user.RegisterUserServiceServer(gs,gusrsrv)

	reflection.Register(gs)

	l, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		return err
	}

	err = gs.Serve(l)
	if err != nil {
		return err
	}
	return nil
}
