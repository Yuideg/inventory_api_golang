package order

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)
type orderServiceServer struct {
	ordserv  module.OrderUsecase
}
func NewOrderServer(ord  module.OrderUsecase) OrderServiceServer {
	return &orderServiceServer{ordserv: ord}
}
func (o orderServiceServer) CreateOrder(ctx context.Context, request *CreateOrderRequest) (*CreateOrderResponse, error) {
	_, err := o.ordserv.CreateOrder(*ConvertProtoOrderToOrder(request.Order))
	if err!=nil {
		return nil, err
	}
	return &CreateOrderResponse{Id:request.Order.Id }, nil
}

func (o orderServiceServer) OrderByID(ctx context.Context, request *OrderByIDRequest) (*OrderByIDResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	order, errs := o.ordserv.GetOrderByID(id)
	if errs!=nil {
		return nil, errs
	}
	return &OrderByIDResponse{Order: ConvertOrderToProto(order)}, nil
}

func (o orderServiceServer) GetOrderByUserID(ctx context.Context, request *OrderByUSerIDRequest) (*OrderByUserIDResponse, error) {
	id, _ :=uuid.FromString(request.UserId)
	order, errs := o.ordserv.GetOrderByUserID(id)
	if errs!=nil {
		return nil, errs
	}
	return &OrderByUserIDResponse{Order: ConvertToProtoOrders(order)}, nil
}

func (o orderServiceServer) OrderList(ctx context.Context, request *OrderListRequest) (*OrderListResponse, error) {
	orders, errs := o.ordserv.GetOrders()
	if errs!=nil{
		return nil, errs
	}
	return &OrderListResponse{Order: ConvertToProtoOrders(orders),}, nil
}

func (o orderServiceServer) UpdateOrder(ctx context.Context, request *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	ord := ConvertProtoOrderToOrder(request.Order)
	_, errs := o.ordserv.UpdateOrder(ord)
	if errs!=nil {
		return nil, errs
	}
	return &UpdateOrderResponse{}, nil
}

func (o orderServiceServer) DeleteOrder(ctx context.Context, request *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	_ = o.ordserv.DeleteOrder(id)

	return &DeleteOrderResponse{}, nil
}

func (o orderServiceServer) mustEmbedUnimplementedOrderServiceServer() {
	panic("implement me")
}









func ConvertToProtoOrders(orders []model.Order) []*Order {
	var order_proto []*Order
	for i := 0; i < len(orders); i++ {
		order := Order{
			Id:orders[i].ID.String(),
			Quantity:  orders[i].Quantity,
			Unit: orders[i].Unit,
			ConfirmBy:  orders[i].CertifiedBy.String(),
			UserId: orders[i].Customer_ID.String(),
			ProductId:  orders[i].Product_ID.String(),
			Status: orders[i].Status,
			ExpiredOn:  orders[i].ExpiredOn.String(),
			CreatedOn: orders[i].CreatedOn.String(),
			UpdatedOn: orders[i].UpdatedOn.String(),
		}
		order_proto = append(order_proto, &order)
	}
	return order_proto
}

func ConvertProtoOrderToOrder(order *Order) *model.Order {
	id, _ :=uuid.FromString(order.Id)
	cert, _ :=uuid.FromString(order.ConfirmBy)
	user, _ :=uuid.FromString(order.UserId)
	proid, _ :=uuid.FromString(order.ProductId)
	ct, _ := time.Parse(time.RFC3339,order.CreatedOn)
	ut, _ := time.Parse(time.RFC3339,order.UpdatedOn)
	et, _ := time.Parse(time.RFC3339,order.ExpiredOn)
	return &model.Order{
		ID:id,
		Quantity:  order.Quantity,
		Unit: order.Unit,
		CertifiedBy:  cert,
		Customer_ID: user,
		Product_ID:  proid,
		Status: order.Status,
		ExpiredOn: et,
		CreatedOn: ct,
		UpdatedOn: ut,
	}
}







func ConvertOrderToProto(order *model.Order) *Order {
	return &Order{
		Id:order.ID.String(),
		Quantity:  order.Quantity,
		Unit: order.Unit,
		ConfirmBy:  order.CertifiedBy.String(),
		UserId: order.Customer_ID.String(),
		ProductId:  order.Product_ID.String(),
		Status: order.Status,
		ExpiredOn:  order.ExpiredOn.String(),
		CreatedOn: order.CreatedOn.String(),
		UpdatedOn: order.UpdatedOn.String(),
	}
}