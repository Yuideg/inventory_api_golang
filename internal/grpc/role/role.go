package role

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
)
type roleServiceServer struct {
	rolserv  module.RoleServices
}
func NewRoleServer(role  module.RoleServices) RoleServiceServer {
	return &roleServiceServer{rolserv: role}
}
func (r roleServiceServer) CreateRole(ctx context.Context, request *CreateRoleRequest) (*CreateRoleResponse, error) {
	_, err := r.rolserv.CreateRole(*ConvertProtoRoleToRole(request.Role))
	if err!=nil {
		return nil, err
	}
	return &CreateRoleResponse{Id:request.Role.Id }, nil
}

func (r roleServiceServer) RoleByID(ctx context.Context, request *RoleByIDRequest) (*RoleByIDResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	role, errs := r.rolserv.GetRoleByID(id)
	if errs!=nil {
		return nil, errs
	}
	return &RoleByIDResponse{Role: ConvertRoleToProto(role)}, nil
}

func (r roleServiceServer) RoleList(ctx context.Context, request *RoleListRequest) (*RoleListResponse, error) {
	roles, errs := r.rolserv.GetRoles()
	if errs!=nil{
		return nil, errs
	}
	return &RoleListResponse{Roles: ConvertToProtoRoles(roles)}, nil
}

func (r roleServiceServer) UpdateRole(ctx context.Context, request *UpdateRoleRequest) (*UpdateRoleResponse, error) {
	rol := ConvertProtoRoleToRole(request.Role)
	_, errs := r.rolserv.UpdateRole(rol)
	if errs!=nil {
		return nil, errs
	}
	return &UpdateRoleResponse{}, nil
}

func (r roleServiceServer) DeleteRole(ctx context.Context, request *DeleteRoleRequest) (*DeleteRoleResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	_ = r.rolserv.DeleteRole(id)
	return &DeleteRoleResponse{}, nil
}

func (r roleServiceServer) mustEmbedUnimplementedRoleServiceServer() {
	panic("implement me")
}



func ConvertToProtoRoles(roles []model.Role) []*Role {
	var role_proto []*Role
	for i := 0; i < len(roles); i++ {
		rol := Role{
			Id:roles[i].ID.String(),
			Name: roles[i].Name,
		}
		role_proto = append(role_proto, &rol)
	}
	return role_proto
}

func ConvertProtoRoleToRole(rol *Role) *model.Role {
	id, _ :=uuid.FromString(rol.Id)
	return &model.Role{
		ID:id,
		Name: rol.Name,
	}
}

func ConvertRoleToProto(rol *model.Role) *Role {
	return &Role{
		Id:rol.ID.String(),
		Name: rol.Name,

	}
}