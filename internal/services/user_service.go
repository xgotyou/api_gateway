package services

import (
	"context"
	"fmt"

	"github.com/xgotyou/api_gateway/internal/dtos"
	"github.com/xgotyou/api_gateway/internal/http"
	xworkpb "github.com/xgotyou/api_gateway/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userService struct {
	addr string
}

func NewUserService(addr string) http.UserService {
	return &userService{addr}
}

func (us *userService) GetUser(id int) (*dtos.User, error) {
	cc, err := grpc.Dial(us.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("can't dial UserService via gRPC: %w", err)
	}
	defer cc.Close()

	c := xworkpb.NewUserServiceClient(cc)
	user, err := c.GetUser(context.Background(), &xworkpb.UserId{Id: int32(id)})
	if err != nil {
		return nil, fmt.Errorf("error while fetching user from UserService: %w", err)
	}

	userDTO := &dtos.User{
		Id:        int(user.GetId()),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		// TODO: BirthDate: user.BirthDate,
		// TODO: Role: user.Role,
	}

	return userDTO, nil
}

func (us *userService) CreateUser(params http.CreateUserParams) (*dtos.User, error) {
	cc, err := grpc.Dial(us.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("can't dial UserService via gRPC: %w", err)
	}
	defer cc.Close()

	c := xworkpb.NewUserServiceClient(cc)
	user, err := c.CreateUser(context.Background(), &xworkpb.User{FirstName: params.FirstName, LastName: params.LastName, Role: params.Role})
	if err != nil {
		return nil, fmt.Errorf("error while fetching user from UserService: %w", err)
	}

	userDTO := &dtos.User{
		Id:        int(user.GetId()),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		// TODO: BirthDate: &time.Time{},
		Role: dtos.Role(user.Role),
	}

	return userDTO, nil
}
