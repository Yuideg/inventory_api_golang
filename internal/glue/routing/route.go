package routing

import (
	"github.com/Yideg/inventory_app/internal/handler/rest"
	"github.com/Yideg/inventory_app/platform/routers"
	"net/http"
)
//AllRouting generates all possible routes int the application
func AllRouting(ch  rest.UserHandler, oh rest.OrderHandler, ph rest.ProductHandler, r rest.RoleHandler, stoh rest.StockHandler, suph rest.SupplierHandler) []routers.Router {
	return []routers.Router{
		{
			Method:      http.MethodGet,
			Path:        "/users",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:    ch.UsersHandler,

		},
		{
			Method:      http.MethodPost,
			Path:        "/login",
			Authenticated: nil,
			Authorized: nil,
			Handler:    ch.LoginUserHandler,

		},
		{
			Method:      http.MethodGet,
			Path:        "/users/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ch.GetUserByIDHandler,

		},

		{
			Method:      http.MethodPost,
			Path:        "/users",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ch.CreateUserHandler,

		},
		{
			Method:      http.MethodPut,
			Path:        "/users/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ch.UpdateUserHandler,
		},

		{
			Method:      http.MethodDelete,
			Path:        "/users/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ch.DeleteUserHandler,
		},
		{
			Method:      http.MethodGet,
			Path:        "/orders",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     oh.Orders,

		},
		{
			Method:      http.MethodGet,
			Path:        "/orders/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     oh.OrderById,

		},

		{
			Method:      http.MethodPost,
			Path:        "/orders",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     oh.CreateOrder,

		},
		{
			Method:      http.MethodPut,
			Path:        "/orders/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     oh.UpdateOrder,

		},

		{
			Method:      http.MethodDelete,
			Path:        "/orders/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     oh.DeleteOrder,

		},
		{
			Method:      http.MethodGet,
			Path:        "/products",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ph.Products,

		},
		{
			Method:      http.MethodGet,
			Path:        "/products/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ph.ProductById,

		},

		{
			Method:      http.MethodPost,
			Path:        "/products",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ph.CreateProduct,

		},
		{
			Method:      http.MethodPut,
			Path:        "/products/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ph.UpdateProduct,

		},

		{
			Method:      http.MethodDelete,
			Path:        "/products/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     ph.DeleteProduct,

		},
		{
			Method:      http.MethodGet,
			Path:        "/role",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     r.Roles,

		},
		{
			Method:      http.MethodGet,
			Path:        "/role/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     r.RoleById,

		},

		{
			Method:      http.MethodPost,
			Path:        "/role",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     r.CreateRole,

		},
		{
			Method:      http.MethodPut,
			Path:        "/role/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     r.UpdateRole,

		},

		{
			Method:      http.MethodDelete,
			Path:        "/role/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     r.DeleteRole,

		},
		{
			Method:      http.MethodGet,
			Path:        "/stocks",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     stoh.Stocks,

		},
		{
			Method:      http.MethodGet,
			Path:        "/stocks/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     stoh.StockById,

		},

		{
			Method:      http.MethodPost,
			Path:        "/stocks",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     stoh.CreateStock,

		},
		{
			Method:      http.MethodPut,
			Path:        "/stocks/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     stoh.UpdateStock,

		},

		{
			Method:      http.MethodDelete,
			Path:        "/stocks/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     stoh.DeleteStock,

		},
		{
			Method:      http.MethodGet,
			Path:        "/supplier",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     suph.Suppliers,

		},
		{
			Method:      http.MethodGet,
			Path:        "/supplier/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     suph.SupplierById,

		},

		{
			Method:      http.MethodPost,
			Path:        "/supplier",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     suph.CreateSupplier,

		},
		{
			Method:      http.MethodPut,
			Path:        "/supplier/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     suph.UpdateSupplier,

		},

		{
			Method:      http.MethodDelete,
			Path:        "/supplier/:id",
			Authenticated: ch.Authenticated,
			Authorized: ch.Authorized,
			Handler:     suph.DeleteSupplier,

		},
	}
}

