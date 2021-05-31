package user

import (
	"github.com/Yideg/inventory_app/internal/module/usecase"
	"golang.org/x/net/context"
)

type userServiceServerIMP struct {
	user * UserServiceServer
	us   *usecase.UserUsecaseIMP
}

func (u userServiceServerIMP) CreateUser(ctx context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	panic("implement me")
}

func (u userServiceServerIMP) GetUserByID(ctx context.Context, request *GetUserByIDRequest) (*GetUserByIDResponse, error) {
	panic("implement me")
}

func (u userServiceServerIMP) GetUserList(ctx context.Context, request *GetUserListRequest) (*GetUserListResponse, error) {
	panic("implement me")
}

func (u userServiceServerIMP) UpdateUser(ctx context.Context, request *UpdateUserRequest) (*UpdateUserResponse, error) {
	panic("implement me")
}

func (u userServiceServerIMP) DeleteUser(ctx context.Context, request *DeleteUserRequest) (*DeleteUserResponse, error) {
	panic("implement me")
}

func (u userServiceServerIMP) mustEmbedUnimplementedUserServiceServer() {
	panic("implement me")
}

func NewUserServiceServer(us *UserServiceServer,userserv   *usecase.UserUsecaseIMP) UserServiceServer {
	return &userServiceServerIMP{user: us,us: userserv}
}



