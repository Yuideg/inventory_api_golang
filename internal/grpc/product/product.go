package product

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)
type productServiceServer struct {
	prodserv  module.ProductUsecase
}
func NewProductServer(prod  module.ProductUsecase) ProductServiceServer {
	return &productServiceServer{prodserv: prod}
}
func (p productServiceServer) CreateProduct(ctx context.Context, request *CreateProductRequest) (*CreateProductResponse, error) {
	_, err := p.prodserv.CreateProduct(*ConvertProtoProductToProduct(request.Product))
	if err!=nil {
		return nil, err
	}
	return &CreateProductResponse{Id:request.Product.Id }, nil
}

func (p productServiceServer) ProductByID(ctx context.Context, request *ProductByIDRequest) (*ProductByIDResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	product, errs := p.prodserv.GetProductsByID(id)
	if errs!=nil {
		return nil, errs
	}
	return &ProductByIDResponse{Product: ConvertProductToProto(product)}, nil
}

func (p productServiceServer) ProductList(ctx context.Context, request *ProductListRequest) (*ProductListResponse, error) {
	products, errs := p.prodserv.GetProducts()
	if errs!=nil{
		return nil, errs
	}
	return &ProductListResponse{Products: ConvertToProtoProduct(products),}, nil
}

func (p productServiceServer) UpdateProduct(ctx context.Context, request *UpdateProductRequest) (*UpdateProductResponse, error) {
	prod := ConvertProtoProductToProduct(request.Product)
	_, errs := p.prodserv.UpdateProduct(prod)
	if errs!=nil {
		return nil, errs
	}
	return &UpdateProductResponse{}, nil
}

func (p productServiceServer) DeleteProduct(ctx context.Context, request *DeleteProductRequest) (*DeleteProductResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	_ = p.prodserv.DeleteProduct(id)
	return &DeleteProductResponse{}, nil
}

func (p productServiceServer) mustEmbedUnimplementedProductServiceServer() {
	panic("implement me")
}


func ConvertToProtoProduct(products []model.Product) []*Product {
	var product_proto []*Product
	for i := 0; i < len(products); i++ {
		product := Product{
			Id:products[i].ID.String(),
			Barcode: products[i].Barcode,
			ProductName:  products[i].ProductName,
			StockId: products[i].StockId.String(),
			Brand:  products[i].Brand,
			Image: products[i].Image,
			Cost:  products[i].Cost,
			Price: products[i].Price,
			Unit:  products[i].Unit,
			Tax: products[i].Tax,
			Discount: products[i].Discount,
			MfgDate: products[i].MfgDate.String(),
			ExpiredAt: products[i].ExpiredOn.String(),
			CreatedOn: products[i].CreatedOn.String(),
			UpdatedOn: products[i].UpdatedOn.String(),
		}
		product_proto = append(product_proto, &product)
	}
	return product_proto
}

func ConvertProtoProductToProduct(product *Product) *model.Product {
	id, _ :=uuid.FromString(product.Id)
	stockid, _ :=uuid.FromString(product.StockId)
	mfg, _ := time.Parse(time.RFC3339,product.MfgDate)
	ut, _ := time.Parse(time.RFC3339,product.UpdatedOn)
	et, _ := time.Parse(time.RFC3339,product.ExpiredAt)
	ct, _ :=time.Parse(time.RFC3339,product.CreatedOn)
	return &model.Product{
		ID:id,
		Barcode: product.Barcode,
		ProductName:  product.ProductName,
		StockId: stockid,
		Brand:  product.Brand,
		Image: product.Image,
		Cost:  product.Cost,
		Price: product.Price,
		Unit:  product.Unit,
		Tax: product.Tax,
		Discount: product.Discount,
		MfgDate: mfg,
		ExpiredOn: et,
		CreatedOn: ct,
		UpdatedOn: ut,
	}
}


func ConvertProductToProto(product *model.Product) *Product {
	return &Product{
		Id:product.ID.String(),
		Barcode: product.Barcode,
		ProductName:  product.ProductName,
		StockId: product.StockId.String(),
		Brand:  product.Brand,
		Image: product.Image,
		Cost:  product.Cost,
		Price: product.Price,
		Unit:  product.Unit,
		Tax: product.Tax,
		Discount: product.Discount,
		MfgDate: product.MfgDate.String(),
		ExpiredAt: product.ExpiredOn.String(),
		CreatedOn: product.CreatedOn.String(),
		UpdatedOn: product.UpdatedOn.String(),
	}
}