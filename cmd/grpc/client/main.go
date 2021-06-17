package main

import (
	"context"
	"github.com/Yideg/inventory_app/internal/grpc/order"
	"github.com/Yideg/inventory_app/internal/grpc/product"
	"github.com/Yideg/inventory_app/internal/grpc/role"
	"github.com/Yideg/inventory_app/internal/grpc/stock"
	"github.com/Yideg/inventory_app/internal/grpc/supplier"
	"github.com/Yideg/inventory_app/internal/grpc/user"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	order_client :=order.NewOrderServiceClient(conn)
	product_client :=product.NewProductServiceClient(conn)
	role_client :=role.NewRoleServiceClient(conn)
	stock_client :=stock.NewStockServiceClient(conn)
	supplier_client :=supplier.NewSupplierServiceClient(conn)
	user_client :=user.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//Create Order
	createOrderRequest := order.CreateOrderRequest{
		Order: &order.Order{
			Quantity:45,
			Unit: "gram",
			ConfirmBy: "6b8009cf-94ce-4c62-8fc7-486e559be497",
			UserId: "483d9830-ebc6-4e27-863f-ad6d8fa8e68f",
			ProductId: "d5f29826-9d3b-4115-bb18-e6d4afbd4351",
			Status: "paid",
			ExpiredOn: "2022-08-30T00:00:00Z",
		}}
	res1, err := order_client.CreateOrder(ctx, &createOrderRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id := res1.Id

	requestRead := order.OrderByIDRequest{Id: id}
	res2, err := order_client.OrderByID(ctx, &requestRead)
	if err != nil {
		log.Fatalf("Get order failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", res2)

	// Update
	req3 := order.UpdateOrderRequest{
		Order: &order.Order{
			Quantity:25,
			Unit: "gram",
			ConfirmBy: "6b8009cf-94ce-4c62-8fc7-486e559be497",
			UserId: "483d9830-ebc6-4e27-863f-ad6d8fa8e68f",
			ProductId: "d5f29826-9d3b-4115-bb18-e6d4afbd4351",
			Status: "unpaied",
			ExpiredOn: "2022-06-30T00:00:00Z",
		},
	}
	res3, err := order_client.UpdateOrder(ctx, &req3)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", res3)

	// Call List All Order
	req4 := order.OrderListRequest{}
	res4, err := order_client.OrderList(ctx, &req4)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Notes result: <%+v>\n\n", res4)

	// Delete
	req5 := order.DeleteOrderRequest{
		Id: res2.Order.Id,
	}
	res5, err := order_client.DeleteOrder(ctx, &req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", res5)




	//Create Product
	CreateproductRequest := product.CreateProductRequest{
		Product: &product.Product{
			Barcode: "3691-2689",
			ProductName: "plate",
			StockId: "1b6a2b7e-5c19-4823-83da-c8fef4502535",
			Brand: "new",
			Image: "",
			Cost: 20,
			Price: 30.55,
			Unit: "gram",
			Discount: 0,
			Tax: 3,
			MfgDate: "2021-04-30T00:00:00Z",
			ExpiredAt: "2022-04-30T00:00:00Z",
		}}
	p_res1, err := product_client.CreateProduct(ctx, &CreateproductRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

	id = p_res1.Id

	p_requestRead := product.ProductByIDRequest{Id: id}
	p_res2, err := product_client.ProductByID(ctx, &p_requestRead)
	if err != nil {
		log.Fatalf("Get order failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", p_res2)

	// Update
	p_u_req := product.UpdateProductRequest{
		Product: &product.Product{
			Barcode: "3691-2689",
			ProductName: "plate",
			StockId: "1b6a2b7e-5c19-4823-83da-c8fef4502535",
			Brand: "old",
			Image: "",
			Cost: 24,
			Price: 30.35,
			Unit: "gram",
			Discount: 0,
			Tax: 4,
			MfgDate: "2021-04-30T00:00:00Z",
			ExpiredAt: "2022-04-30T00:00:00Z",
		},
	}
	p_u_req3, err := product_client.UpdateProduct(ctx, &p_u_req)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", p_u_req3)

	// Call List All Order
	p_l_req4 := product.ProductListRequest{}
	p_res4, err := product_client.ProductList(ctx, &p_l_req4)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Notes result: <%+v>\n\n", p_res4)

	// Delete
	p_req5 := order.DeleteOrderRequest{
		Id: p_res2.Product.Id,
	}
	p_res5, err := order_client.DeleteOrder(ctx, &p_req5)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", p_res5)




	//Create Role
	createRoleRequest := role.CreateRoleRequest{
		Role: &role.Role{
			Name: "Other",
		}}
	res, err := role_client.CreateRole(ctx, &createRoleRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res)

	id = res.Id

	rest := role.RoleByIDRequest{Id: id}
	re, err := role_client.RoleByID(ctx, &rest)
	if err != nil {
		log.Fatalf("Get order failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", re)

	// Update
	req := role.UpdateRoleRequest{
		Role: &role.Role{
			Name: "Others",
		},
	}
	r, err := role_client.UpdateRole(ctx, &req)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", r)

	// Call List All Role
	rer := role.RoleListRequest{}
	res4r, err := role_client.RoleList(ctx, &rer)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Notes result: <%+v>\n\n", res4r)

	// Delete
	reqd := role.DeleteRoleRequest{
		Id: rest.Id,
	}
	rd, err := role_client.DeleteRole(ctx, &reqd)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", rd)


	//Create Stock
	createStockRequest := stock.CreateStockRequest{
		Stock: &stock.Stock{
			ItemsCode: "oro-112",
			ItemsName: "coffees",
			NoItems: 3400,
			SupplierId: "4685585e-3ff8-4913-9132-0a9ac070990a",
			Country: "Ethiopia",
			City: "Adama",
			State: "oromia",
			Zipcode: 1222,
			Street: "algeria street",
			Latitude: 11,
			Longitude: 33.66,
		}}
	rs, err := stock_client.CreateStock(ctx, &createStockRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", rs)

	id = rs.Id

	rss := stock.StockByIDRequest{Id: id}
	st_id, err := stock_client.StockByID(ctx, &rss)
	if err != nil {
		log.Fatalf("Get order failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", st_id)

	// Update
	req_s_id := stock.UpdateStockRequest{
		Stock: &stock.Stock{
			ItemsCode: "oro-112",
			ItemsName: "coffees",
			NoItems: 3400,
			SupplierId: "4685585e-3ff8-4913-9132-0a9ac070990a",
			Country: "Ethiopia",
			City: "Adama",
			State: "oromia",
			Zipcode: 1222,
			Street: "algeria street",
			Latitude: 11,
			Longitude: 33.66,
		},
	}
	rst, err := stock_client.UpdateStock(ctx, &req_s_id)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", rst)

	// Call List All Stock
	stocks := stock.StockListRequest{}
	stock4, err := stock_client.StockList(ctx, &stocks)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Stock result: <%+v>\n\n", stock4)

	// Delete
	std := stock.DeleteStockRequest{
		Id: rest.Id,
	}
	rds, err := stock_client.DeleteStock(ctx, &std)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", rds)





	//Create Supplier
	createSupplierRequest := supplier.CreateSupplierRequest{
		Supplier: &supplier.Supplier{
			SupplierName: "Oromia Coffee Farmers Cooperative",
			Country: "Ethiopia",
			City: "Adama",
			State: "oromia",
			Zipcode: 1222,
			Street: "algeria street",
			Latitude: 11,
			Longitude: 33.66,
			Fax: "234",
			Pobox: "390",
			Email: "oromiacoffeeunion@gmail.com",
			Website: "https://www.oromiacoffeeunion.org/",
			Phone: "+251 11 391 9070",
		}}
	sp, err := supplier_client.CreateSupplier(ctx, &createSupplierRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", sp)

	id = sp.Id

	spps := supplier.SupplierByIDRequest{Id: id}
	sup_id, err := supplier_client.SupplierByID(ctx, &spps)
	if err != nil {
		log.Fatalf("Get order failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", sup_id)

	// Update
	sup_req := supplier.UpdateSupplierRequest{
		Supplier: &supplier.Supplier{
			SupplierName: "Oromia Coffee Farmers Cooperative",
			Country: "Ethiopia",
			City: "Adama",
			State: "oromia",
			Zipcode: 1222,
			Street: "algeria street",
			Latitude: 11,
			Longitude: 33.66,
			Fax: "234",
			Pobox: "390",
			Email: "oromiacoffeeunion@gmail.com",
			Website: "https://www.oromiacoffeeunion.org/",
			Phone: "+251 11 391 9070",
		},
	}
	supps, err := supplier_client.UpdateSupplier(ctx, &sup_req)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", supps)

	// Call List All supplier
	supplr := supplier.SupplierListRequest{}
	suppl, err := supplier_client.SupplierList(ctx, &supplr)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Stock result: <%+v>\n\n", suppl)

	// Delete
	suppd := supplier.DeleteSupplierRequest{
		Id: rest.Id,
	}
	sppd, err := supplier_client.DeleteSupplier(ctx, &suppd)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", sppd)








	//Create User
	createUserRequest := user.CreateUserRequest{
		User: &user.User{
			Username: "john32",
			Password: "90901234typo",
			FirstName: "Yohannes",
			LastName: "Ketema",
			Gender: "Male",
			Email: "yohannesketem@gmail.com",
			Country: "Ethiopia",
			City: "Addis Ababa",
			State: "Addis Ababa",
			Zipcode: 1000,
			Street: "algeria street",
			Latitude: 12.564,
			Longitude: 3.677,
			Phone: "+251 91 093 7962",
			RoleId: "064a3d9f-5032-48d2-9645-b488b282ab96",
		}}
	usr, err := user_client.CreateUser(ctx, &createUserRequest)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", sp)

	id = usr.Id

	users := user.GetUserByIDRequest{Id: id}
	user_id, err := user_client.GetUserByID(ctx, &users)
	if err != nil {
		log.Fatalf("Get order failed: %v", err)
	}
	log.Printf("Get result: <%+v>\n\n", user_id)

	// Update
	user_upd := user.UpdateUserRequest{
		User: &user.User{
			Username: "john32",
			Password: "90901234typo",
			FirstName: "Yohannes",
			LastName: "Ketema",
			Gender: "Male",
			Email: "yohannesketem@gmail.com",
			Country: "Ethiopia",
			City: "Addis Ababa",
			State: "Addis Ababa",
			Zipcode: 1000,
			Street: "algeria street",
			Latitude: 12.564,
			Longitude: 3.677,
			Phone: "+251 91 093 7962",
			RoleId: "064a3d9f-5032-48d2-9645-b488b282ab96",
		},
	}
	usrsr, err := user_client.UpdateUser(ctx, &user_upd)
	if err != nil {
		log.Fatalf("Update failed: %v", err)
	}
	log.Printf("Update result: <%+v>\n\n", usrsr)

	// Call List All User
	uusr := user.GetUserListRequest{}
	usssr, err := user_client.GetUserList(ctx, &uusr)
	if err != nil {
		log.Fatalf("List All Notes failed: %v", err)
	}
	log.Printf("List All Stock result: <%+v>\n\n", usssr)

	// Delete
	userd := user.DeleteUserRequest{
		Id: users.Id,
	}
	usdd, err := user_client.DeleteUser(ctx, &userd)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}
	log.Printf("Delete result: <%+v>\n\n", usdd)
}

