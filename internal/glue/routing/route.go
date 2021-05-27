package routing

import (
	"github.com/Yideg/inventory_app/internal/handler/rest"
	"github.com/Yideg/inventory_app/platform/routers"
	"net/http"
)
//AllRouting generates all possible routes int the application
func AllRouting(ch rest.CustomerHandler,oh rest.OrderHandler,ph rest.ProductHandler,r rest.RoleHandler,
	sth rest.StaffHandler,stoh rest.StockHandler,suph rest.SupplierHandler) []routers.Router {
	return []routers.Router{
		{
			Method:      http.MethodGet,
			Path:        "/customers",
			Handler:     ch.Customers,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/customers/:id",
			Handler:     ch.CustomerById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/customers",
			Handler:     ch.CreateCustomer,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/customers/:id",
			Handler:     ch.UpdateCustomer,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/customers/:id",
			Handler:     ch.DeleteCustomer,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/orders",
			Handler:     oh.Orders,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/orders/:id",
			Handler:     oh.OrderById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/orders",
			Handler:     oh.CreateOrder,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/orders/:id",
			Handler:     oh.UpdateOrder,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/orders/:id",
			Handler:     oh.DeleteOrder,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/products",
			Handler:     ph.Products,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/products/:id",
			Handler:     ph.ProductById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/products",
			Handler:     ph.CreateProduct,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/products/:id",
			Handler:     ph.UpdateProduct,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/products/:id",
			Handler:     ph.DeleteProduct,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/role",
			Handler:     r.Roles,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/role/:id",
			Handler:     r.RoleById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/role",
			Handler:     r.CreateRole,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/role/:id",
			Handler:     r.UpdateRole,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/role/:id",
			Handler:     r.DeleteRole,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/staff",
			Handler:     sth.Staffs,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/staff/:id",
			Handler:     sth.StaffById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/staff",
			Handler:     sth.CreateStaff,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/staff/:id",
			Handler:     sth.UpdateStaff,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodDelete,
			Path:        "/staff/:id",
			Handler:     sth.DeleteStaff,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/stocks",
			Handler:     stoh.Stocks,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/stocks/:id",
			Handler:     stoh.StockById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/stocks",
			Handler:     stoh.CreateStock,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/stocks/:id",
			Handler:     stoh.UpdateStock,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/stocks/:id",
			Handler:     stoh.DeleteStock,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/supplier",
			Handler:     suph.Suppliers,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodGet,
			Path:        "/supplier/:id",
			Handler:     suph.SupplierById,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodPost,
			Path:        "/supplier",
			Handler:     suph.CreateSupplier,
			MiddleWares: nil,
		},
		{
			Method:      http.MethodPut,
			Path:        "/supplier/:id",
			Handler:     suph.UpdateSupplier,
			MiddleWares: nil,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/supplier/:id",
			Handler:     suph.DeleteSupplier,
			MiddleWares: nil,
		},
	}


}

