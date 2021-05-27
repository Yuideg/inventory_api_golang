package customer

import "golang.org/x/net/context"
import (
	"github.com/Yideg/inventory_app/internal/module/usecase"
)

type customerServiceServerIMP struct {
	customer *CustomerServiceServer
	cust   *usecase.CustomerUsecaseIMP
}
func NewCustomerServiceServer(cust *CustomerServiceServer,custmerusecase   *usecase.CustomerUsecaseIMP) CustomerServiceServer {
	return &customerServiceServerIMP{customer: cust,cust: custmerusecase}
}
func (c *customerServiceServerIMP) CreateCustomer(ctx context.Context, request *CreateCustomerRequest) (*CreateCustomerResponse, error) {
	//cust,err:=c.cust.Create(request.Customer)

	//return cust, err
	panic("")
}

func (c *customerServiceServerIMP) CustomerByID(ctx context.Context, request *CustomerByIDRequest) (*CustomerByIDResponse, error) {
	panic("implement me")
}

func (c *customerServiceServerIMP) CustomerList(ctx context.Context, request *CustomerListRequest) (*CustomerListResponse, error) {
	panic("implement me")
}

func (c *customerServiceServerIMP) UpdateCustomer(ctx context.Context, request *UpdateCustomerRequest) (*UpdateCustomerResponse, error) {
	panic("implement me")
}

func (c *customerServiceServerIMP) DeleteCustomer(ctx context.Context, request *DeleteCustomerRequest) (*DeleteCustomerResponse, error) {
	panic("implement me")
}

func (c *customerServiceServerIMP) mustEmbedUnimplementedCustomerServiceServer() {
	panic("implement me")
}


