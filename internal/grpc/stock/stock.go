package stock

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)
type stockServiceServer struct {
	stockserv  module.StockUsecae
}
func NewStockServer(stock  module.StockUsecae) StockServiceServer {
	return &stockServiceServer{stockserv: stock}
}
func (s stockServiceServer) CreateStock(ctx context.Context, request *CreateStockRequest) (*CreateStockResponse, error) {
	_, err := s.stockserv.CreateStock(*ConvertProtoStockToStock(request.Stock))
	if err!=nil {
		return nil, err
	}
	return &CreateStockResponse{Id:request.Stock.Id }, nil
}

func (s stockServiceServer) StockByID(ctx context.Context, request *StockByIDRequest) (*StockByIDResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	stock, errs := s.stockserv.GetStockByID(id)
	if errs!=nil {
		return nil, errs
	}
	return &StockByIDResponse{Stock: ConvertStockToProto(stock)}, nil
}

func (s stockServiceServer) StockList(ctx context.Context, request *StockListRequest) (*StockListResponse, error) {
	stocks, errs := s.stockserv.GetStocks()
	if errs!=nil{
		return nil, errs
	}
	return &StockListResponse{Stock: ConvertToProtoStock(stocks),}, nil

}

func (s stockServiceServer) UpdateStock(ctx context.Context, request *UpdateStockRequest) (*UpdateStockResponse, error) {
	stoc := ConvertProtoStockToStock(request.Stock)
	_, errs := s.stockserv.UpdateStock(stoc)
	if errs!=nil {
		return nil, errs
	}
	return &UpdateStockResponse{}, nil
}

func (s stockServiceServer) DeleteStock(ctx context.Context, request *DeleteStockRequest) (*DeleteStockResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	_ = s.stockserv.DeleteStock(id)
	return &DeleteStockResponse{}, nil
}
func (s stockServiceServer) mustEmbedUnimplementedStockServiceServer() {
	panic("implement me")
}
func ConvertToProtoStock(stocks []model.Stock) []*Stock {
	var stock_proto []*Stock
	for i := 0; i < len(stocks); i++ {
		stock := Stock{
			Id:stocks[i].ID.String(),
			ItemsCode: stocks[i].ProductsCode,
			ItemsName:  stocks[i].ProductsName,
			NoItems: int32(stocks[i].No_products),
			SupplierId:  stocks[i].SupplierID.String(),
			Country: stocks[i].Address.Country,
			City:  stocks[i].Address.City,
			State: stocks[i].Address.State,
			Zipcode:  int32(stocks[i].Address.Zip),
			Street: stocks[i].Address.Street,
			Latitude: float32(stocks[i].Address.Latitude),
			Longitude: float32(stocks[i].Address.Longitude),
			CreatedOn: stocks[i].CreatedOn.String(),
			UpdatedOn: stocks[i].UpdatedOn.String(),
		}
		stock_proto = append(stock_proto, &stock)
	}
	return stock_proto
}
func ConvertProtoStockToStock(stock *Stock) *model.Stock {
	id, _ :=uuid.FromString(stock.SupplierId)
	id1, _ :=uuid.FromString(stock.Id)
	ut, _ := time.Parse(time.RFC3339,stock.UpdatedOn)
	ct, _ :=time.Parse(time.RFC3339,stock.CreatedOn)
	return &model.Stock{
		ID:id1,
		ProductsCode: stock.ItemsCode,
		ProductsName:  stock.ItemsName,
		No_products: int(stock.NoItems),
		SupplierID:  id,
		Address: model.SubAddress{
			Country: stock.Country,
			City:  stock.City,
			State: stock.State,
			Zip:  int(stock.Zipcode),
			Street: stock.Street,
			Latitude: float64(stock.Latitude),
			Longitude: float64(stock.Longitude),
		},
		CreatedOn: ct,
		UpdatedOn: ut,
	}
}

func ConvertStockToProto(stock *model.Stock) *Stock {
	return &Stock{
		Id:stock.ID.String(),
		ItemsCode: stock.ProductsCode,
		ItemsName:  stock.ProductsName,
		NoItems: int32(stock.No_products),
		SupplierId:  stock.SupplierID.String(),
		Country: stock.Address.Country,
		City:  stock.Address.City,
		State: stock.Address.State,
		Zipcode:  int32(stock.Address.Zip),
		Street: stock.Address.Street,
		Latitude: float32(stock.Address.Latitude),
		Longitude: float32(stock.Address.Longitude),
		CreatedOn: stock.CreatedOn.String(),
		UpdatedOn: stock.UpdatedOn.String(),
	}
}