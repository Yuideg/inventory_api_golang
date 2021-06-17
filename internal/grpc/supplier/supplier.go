package supplier

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)
type supplierServiceServer struct {
	supplserv  module.SupplierUsecase
}
func NewSupplierServer(suppl  module.SupplierUsecase) SupplierServiceServer {
	return &supplierServiceServer{supplserv: suppl}
}
func (s supplierServiceServer) CreateSupplier(ctx context.Context, request *CreateSupplierRequest) (*CreateSupplierResponse, error) {
	_, err := s.supplserv.CreateSupplier(*ConvertProtoSupplierToSupplier(request.Supplier))
	if err!=nil {
		return nil, err
	}
	return &CreateSupplierResponse{Id:request.Supplier.Id }, nil
}

func (s supplierServiceServer) SupplierByID(ctx context.Context, request *SupplierByIDRequest) (*SupplierByIDResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	suppl, errs := s.supplserv.GetSupplierByID(id)
	if errs!=nil {
		return nil, errs
	}
	return &SupplierByIDResponse{Supplier: ConvertSupplierToProto(suppl)}, nil
}

func (s supplierServiceServer) GetSupplierBySupplierID(ctx context.Context, request *SupplierBySupplierIDRequest) (*SupplierBySupplierIDResponse, error) {
	panic("implement me")
}

func (s supplierServiceServer) SupplierList(ctx context.Context, request *SupplierListRequest) (*SupplierListResponse, error) {
	sup, errs := s.supplserv.GetSuppliers()
	if errs!=nil{
		return nil, errs
	}
	return &SupplierListResponse{Supplier: ConvertToProtoSupplier(sup),}, nil

}

func (s supplierServiceServer) UpdateSupplier(ctx context.Context, request *UpdateSupplierRequest) (*UpdateSupplierResponse, error) {
	sup := ConvertProtoSupplierToSupplier(request.Supplier)
	_, errs := s.supplserv.UpdateSupplier(sup)
	if errs!=nil {
		return nil, errs
	}
	return &UpdateSupplierResponse{}, nil
}

func (s supplierServiceServer) DeleteSupplier(ctx context.Context, request *DeleteSupplierRequest) (*DeleteSupplierResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	_ = s.supplserv.DeleteSupplier(id)
	return &DeleteSupplierResponse{}, nil
}

func (s supplierServiceServer) mustEmbedUnimplementedSupplierServiceServer() {
	panic("implement me")
}


func ConvertToProtoSupplier(suppliers []model.Supplier) []*Supplier {
	var supplier_proto []*Supplier
	for i := 0; i < len(suppliers); i++ {
		sup := Supplier{
			Id:suppliers[i].ID.String(),
			SupplierName: suppliers[i].Name,
			Country: suppliers[i].Address.Country,
			City:  suppliers[i].Address.City,
			State: suppliers[i].Address.State,
			Zipcode:  int32(suppliers[i].Address.Zip),
			Street: suppliers[i].Address.Street,
			Latitude: float32(suppliers[i].Address.Latitude),
			Longitude: float32(suppliers[i].Address.Longitude),
			Fax: suppliers[i].Fax,
			Pobox: suppliers[i].PoBox,
			Email: suppliers[i].Email,
			Website: suppliers[i].WebSite,
			Phone: suppliers[i].Phone,
			CreatedOn: suppliers[i].CreatedOn.String(),
			UpdatedOn: suppliers[i].UpdatedOn.String(),
		}
		supplier_proto = append(supplier_proto, &sup)
	}
	return supplier_proto
}



func ConvertProtoSupplierToSupplier(suppl *Supplier) *model.Supplier {
	id, _ :=uuid.FromString(suppl.Id)
	ut, _ := time.Parse(time.RFC3339,suppl.UpdatedOn)
	ct, _ :=time.Parse(time.RFC3339,suppl.CreatedOn)
	return &model.Supplier{
		ID: id,
		Name: suppl.SupplierName,
		Address: model.SubAddress{
			Country: suppl.Country,
			City:  suppl.City,
			State: suppl.State,
			Zip:  int(suppl.Zipcode),
			Street: suppl.Street,
			Latitude: float64(suppl.Latitude),
			Longitude: float64(suppl.Longitude),
		},
		Fax: suppl.Fax,
		PoBox: suppl.Pobox,
		Email: suppl.Email,
		WebSite: suppl.Website,
		Phone: suppl.Phone,
		CreatedOn: ct,
		UpdatedOn: ut,
	}
}

func ConvertSupplierToProto(suppl *model.Supplier) *Supplier {
	return &Supplier{
		Id:suppl.ID.String(),
		SupplierName: suppl.Name,
		Country: suppl.Address.Country,
		City:  suppl.Address.City,
		State: suppl.Address.State,
		Zipcode:  int32(suppl.Address.Zip),
		Street: suppl.Address.Street,
		Latitude: float32(suppl.Address.Latitude),
		Longitude: float32(suppl.Address.Longitude),
		Fax: suppl.Fax,
		Pobox: suppl.PoBox,
		Email: suppl.Email,
		Website: suppl.WebSite,
		Phone: suppl.Phone,
		CreatedOn: suppl.CreatedOn.String(),
		UpdatedOn: suppl.UpdatedOn.String(),
	}
}