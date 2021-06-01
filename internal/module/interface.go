package module

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

//Role repository
type RoleRepository interface {
	GetRoles() ([]model.Role, error)
	GetRoleByID(id uuid.UUID) (*model.Role, error)
	CreateRole(info model.Role) (pgconn.CommandTag, error)
	UpdateRole(info *model.Role) (pgconn.CommandTag, error)
	DeleteRole(id uuid.UUID) error
}
//order repository
type OrderRepository interface {
	GetOrders() ([]model.Order, error)
	GetOrderByID(id uuid.UUID) (*model.Order, error)
	GetOrderByUserID(user_id uuid.UUID) ([]model.Order, error)
	CreateOrder(info model.Order) (pgconn.CommandTag, error)
	UpdateOrder(info *model.Order) (pgconn.CommandTag, error)
	DeleteOrder(id uuid.UUID) error
}
//Stock repository
type StockRepository interface {
	GetStocks() ([]model.Stock, error)
	GetStockByID(id uuid.UUID) (*model.Stock, error)
	CreateStock(info model.Stock) (pgconn.CommandTag, error)
	UpdateStock(info *model.Stock) (pgconn.CommandTag, error)
	DeleteStock(id uuid.UUID) error
}
//Product repository
type ProductRepository interface {
	GetProducts() ([]model.Product, error)
	GetProductsByID(id uuid.UUID) (*model.Product, error)
	CreateProduct(info model.Product) (pgconn.CommandTag, error)
	UpdateProduct(info *model.Product) (pgconn.CommandTag, error)
	DeleteProduct(id uuid.UUID) error
}
//Supplier repository
type SupplierRepository interface {
	GetSuppliers() ([]model.Supplier, error)
	GetSupplierByID(id uuid.UUID) (*model.Supplier, error)
	GetSupplierBySupplierID(supplierId uuid.UUID) ([]model.Supplier, error)
	CreateSupplier(info model.Supplier) (pgconn.CommandTag, error)
	UpdateSupplier(info *model.Supplier) (pgconn.CommandTag, error)
	DeleteSupplier(id uuid.UUID) error
}
//User repository
type UserRepository interface {
	GetUsers() ([]model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
	CreateUser(info model.User) (pgconn.CommandTag, error)
	UpdateUser(info *model.User) (pgconn.CommandTag, error)
	DeleteUser(id uuid.UUID) error
}





//role services services
type RoleServices interface {
	GetRoles() ([]model.Role, error)
	GetRoleByID(id uuid.UUID) (*model.Role, error)
	CreateRole(role model.Role) (pgconn.CommandTag, error)
	UpdateRole(role *model.Role) (pgconn.CommandTag, error)
	DeleteRole(id uuid.UUID) error
}

//order usecase
type OrderUsecase interface {
	GetOrders() ([]model.Order, error)
	GetOrderByID(id uuid.UUID) (*model.Order, error)
	GetOrderByUserID(user_id uuid.UUID) ([]model.Order, error)
	CreateOrder(info model.Order) (pgconn.CommandTag, error)
	UpdateOrder(info *model.Order) (pgconn.CommandTag, error)
	DeleteOrder(id uuid.UUID) error
}
//Stock ueecase
type StockUsecae interface {
	GetStocks() ([]model.Stock, error)
	GetStockByID(id uuid.UUID) (*model.Stock, error)
	CreateStock(info model.Stock) (pgconn.CommandTag, error)
	UpdateStock(info *model.Stock) (pgconn.CommandTag, error)
	DeleteStock(id uuid.UUID) error
}
//Product usecase
type ProductUsecase interface {
	GetProducts() ([]model.Product, error)
	GetProductsByID(id uuid.UUID) (*model.Product, error)
	CreateProduct(info model.Product) (pgconn.CommandTag, error)
	UpdateProduct(info *model.Product) (pgconn.CommandTag, error)
	DeleteProduct(id uuid.UUID) error
}
//Supplier usecase
type SupplierUsecase interface {
	GetSuppliers() ([]model.Supplier, error)
	GetSupplierByID(id uuid.UUID) (*model.Supplier, error)
	GetSupplierBySupplierID(supplierId uuid.UUID) ([]model.Supplier, error)
	CreateSupplier(info model.Supplier) (pgconn.CommandTag, error)
	UpdateSupplier(info *model.Supplier) (pgconn.CommandTag, error)
	DeleteSupplier(id uuid.UUID) error
}
//User usecase
type UserUsecase interface {
	GetUsers() ([]model.User, error)
	GetUserByID(id uuid.UUID) (*model.User, error)
	CreateUser(info model.User) (pgconn.CommandTag, error)
	UpdateUser(info *model.User) (pgconn.CommandTag, error)
	DeleteUser(id uuid.UUID) error
}