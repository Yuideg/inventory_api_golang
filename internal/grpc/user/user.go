package user

import (
	"github.com/Yideg/inventory_app/internal/constant/model"
	"github.com/Yideg/inventory_app/internal/module"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"time"
)
type userServiceServer struct {
	userServ  module.UserUsecase
}
func NewUserServer(usr  module.UserUsecase) UserServiceServer {
	return &userServiceServer{userServ: usr}
}
func (u userServiceServer) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	_, err := u.userServ.CreateUser(*ConvertProtoUserToUser(request.User))
	if err!=nil {
		return nil, err
	}
	return &CreateUserResponse{Id:request.User.Id }, nil
}

func (u userServiceServer) GetUserByID(ctx context.Context, request *GetUserByIDRequest) (*GetUserByIDResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	usr, errs := u. userServ.GetUserByID(id)
	if errs!=nil {
		return nil, errs
	}
	return &GetUserByIDResponse{User: ConvertUserToProto(usr)}, nil
}

func (u userServiceServer) GetUserList(ctx context.Context, request *GetUserListRequest) (*GetUserListResponse, error) {
	users, errs := u.userServ.GetUsers()
	if errs!=nil{
		return nil, errs
	}
	return &GetUserListResponse{Users: ConvertToProtoUser(users),}, nil
}

func (u userServiceServer) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	usr := ConvertProtoUserToUser(request.User)
	_, errs := u.userServ.UpdateUser(usr)
	if errs!=nil {
		return nil, errs
	}
	return &UpdateUserResponse{}, nil
}

func (u userServiceServer) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	id, _ :=uuid.FromString(request.Id)
	_ = u.userServ.DeleteUser(id)
	return &DeleteUserResponse{}, nil
}

func (u userServiceServer) mustEmbedUnimplementedUserServiceServer() {
	panic("implement me")
}


func ConvertToProtoUser(users []model.User) []*User {
	var user_proto []*User
	for i := 0; i < len(users); i++ {
		usr := User{
			Id:users[i].ID.String(),
			Username: users[i].Username,
			Password:  users[i].Password,
			FirstName: users[i].FirstName,
			LastName:  users[i].LastName,
			Gender:    users[i].Gender,
			Email:     users[i].Email,
			Country:   users[i].Address.Country,
			City:      users[i].Address.City,
			State:     users[i].Address.State,
			Zipcode:   int32(users[i].Address.Zip),
			Street:    users[i].Address.Street,
			Latitude:  float32(users[i].Address.Latitude),
			Longitude: float32(users[i].Address.Longitude),
			Phone:     users[i].Phone,
			RoleId:    users[i].RoleID.String(),
			CreatedOn: users[i].CreatedOn.String(),
			UpdatedOn: users[i].UpdatedOn.String(),
		}
		user_proto = append(user_proto, &usr)
	}
	return user_proto
}



func ConvertProtoUserToUser(user *User) *model.User {
	id, _ :=uuid.FromString(user.Id)
	role_id, _ :=uuid.FromString(user.RoleId)
	ut, _ := time.Parse(time.RFC3339,user.UpdatedOn)
	ct, _ :=time.Parse(time.RFC3339,user.CreatedOn)
	return &model.User{
		ID: id,
		Username: user.Username,
		Password:  user.Country,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Email:     user.Email,
		Address: model.SubAddress{
			Country: user.Country,
			City:  user.City,
			State: user.State,
			Zip:  int(user.Zipcode),
			Street: user.Street,
			Latitude: float64(user.Latitude),
			Longitude: float64(user.Longitude),
		},
		Phone: user.Phone,
		RoleID: role_id,
		CreatedOn: ct,
		UpdatedOn: ut,
	}
}

func ConvertUserToProto(user *model.User) *User {
	return &User{
		Id:user.ID.String(),
		Username: user.Username,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Email:     user.Email,
		Country:   user.Address.Country,
		City:      user.Address.City,
		State:     user.Address.State,
		Zipcode:   int32(user.Address.Zip),
		Street:    user.Address.Street,
		Latitude:  float32(user.Address.Latitude),
		Longitude: float32(user.Address.Longitude),
		Phone:     user.Phone,
		RoleId:    user.RoleID.String(),
		CreatedOn: user.CreatedOn.String(),
		UpdatedOn: user.UpdatedOn.String(),
	}
}