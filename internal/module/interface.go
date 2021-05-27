package module

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/jackc/pgconn"
	uuid "github.com/satori/go.uuid"
)

//Role repository
type RoleRepository interface {
	Get() ([]model.Role, error)
	GetById(id uuid.UUID) (*model.Role, error)
	Create(info model.Role) (pgconn.CommandTag, error)
	Update(info *model.Role) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//order repository
type OrderRepository interface {
	Get() ([]model.Order, error)
	GetById(id uuid.UUID) (*model.Order, error)
	Create(info model.Order) (pgconn.CommandTag, error)
	Update(info *model.Order) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Stock repository
type StockRepository interface {
	Get() ([]model.Stock, error)
	GetById(id uuid.UUID) (*model.Stock, error)
	Create(info model.Stock) (pgconn.CommandTag, error)
	Update(info *model.Stock) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Staff repository
type StaffRepository interface {
	Get() ([]model.Staff, error)
	GetById(id uuid.UUID) (*model.Staff, error)
	Create(info model.Staff) (pgconn.CommandTag, error)
	Update(info *model.Staff) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Product repository
type ProductRepository interface {
	Get() ([]model.Product, error)
	GetById(id string) (*model.Product, error)
	Create(info model.Product) (pgconn.CommandTag, error)
	Update(info *model.Product) (pgconn.CommandTag, error)
	Delete(id string) error
}
//Supplier repository
type SupplierRepository interface {
	Get() ([]model.Supplier, error)
	GetById(id uuid.UUID) (*model.Supplier, error)
	Create(info model.Supplier) (pgconn.CommandTag, error)
	Update(info *model.Supplier) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Customer repository
type CustomerRepository interface {
	Get() ([]model.Customer, error)
	GetById(id uuid.UUID) (*model.Customer, error)
	Create(info model.Customer) (pgconn.CommandTag, error)
	Update(info *model.Customer) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}





//role services services
type RoleServices interface {
	Get() ([]model.Role, error)
	GetById(id uuid.UUID) (*model.Role, error)
	Create(info model.Role) (pgconn.CommandTag, error)
	Update(info *model.Role) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}

//order usecase
type OrderUsecase interface {
	Get() ([]model.Order, error)
	GetById(id uuid.UUID) (*model.Order, error)
	Create(info model.Order) (pgconn.CommandTag, error)
	Update(info *model.Order) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Stock ueecase
type StockUsecae interface {
	Get() ([]model.Stock, error)
	GetById(id uuid.UUID) (*model.Stock, error)
	Create(info model.Stock) (pgconn.CommandTag, error)
	Update(info *model.Stock) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Staff usecase
type StaffUsecase interface {
	Get() ([]model.Staff, error)
	GetById(id uuid.UUID) (*model.Staff, error)
	Create(info model.Staff) (pgconn.CommandTag, error)
	Update(info *model.Staff) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Product usecase
type ProductUsecase interface {
	Get() ([]model.Product, error)
	GetById(id string) (*model.Product, error)
	Create(info model.Product) (pgconn.CommandTag, error)
	Update(info *model.Product) (pgconn.CommandTag, error)
	Delete(id string) error
}
//Supplier usecase
type SupplierUsecase interface {
	Get() ([]model.Supplier, error)
	GetById(id uuid.UUID) (*model.Supplier, error)
	Create(info model.Supplier) (pgconn.CommandTag, error)
	Update(info *model.Supplier) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}
//Customer usecase
type CustomerUsecase interface {
	Get() ([]model.Customer, error)
	GetById(id uuid.UUID) (*model.Customer, error)
	Create(info model.Customer) (pgconn.CommandTag, error)
	Update(info *model.Customer) (pgconn.CommandTag, error)
	Delete(id uuid.UUID) error
}