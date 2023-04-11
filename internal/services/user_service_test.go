package services

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xgotyou/api_gateway/internal/dtos"
	"github.com/xgotyou/api_gateway/internal/http"
	xworkpb "github.com/xgotyou/api_gateway/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const srvAddr = "localhost:8088"

func TestGetUser_ServerUnavailable(t *testing.T) {
	_, err := NewUserService(srvAddr).GetUser(10)

	require.ErrorContains(t, err, "error while fetching user from UserService:")
}

func TestGetUser(t *testing.T) {
	s := spinUpTestServer(t)
	defer s.Stop()

	expUser := &dtos.User{
		Id:        10,
		FirstName: "Bilbo",
		LastName:  "Baggins",
	}

	user, err := NewUserService(srvAddr).GetUser(expUser.Id)
	if assert.NoError(t, err) {
		assert.Equal(t, expUser, user)
	}
}

func TestCreateUser(t *testing.T) {
	s := spinUpTestServer(t)
	defer s.Stop()

	userParams := http.CreateUserParams{
		FirstName: "Gandalf",
		LastName:  "The White",
		Role:      "Admin",
	}

	expUser := &dtos.User{
		Id:        11,
		FirstName: "Gandalf",
		LastName:  "The White",
		Role:      "Admin",
	}

	user, err := NewUserService(srvAddr).CreateUser(userParams)
	if assert.NoError(t, err) {
		assert.Equal(t, expUser, user)
	}
}

func spinUpTestServer(t *testing.T) *grpc.Server {
	lis, err := net.Listen("tcp", srvAddr)
	if err != nil {
		t.Fatalf("Can't start listen for %v to spin up test gRPC server", srvAddr)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	xworkpb.RegisterUserServiceServer(s, new(mockUserServiceSrv))

	go func() {
		defer lis.Close()
		err = s.Serve(lis)
		if err != nil {
			log.Fatalf("Can't spin up test gRPC server: %v", err)
		}
	}()

	return s
}

type mockUserServiceSrv struct {
	xworkpb.UnimplementedUserServiceServer
}

func (s *mockUserServiceSrv) GetUser(c context.Context, id *xworkpb.UserId) (*xworkpb.User, error) {
	return &xworkpb.User{Id: id.Id, FirstName: "Bilbo", LastName: "Baggins"}, nil
}

func (s *mockUserServiceSrv) CreateUser(c context.Context, u *xworkpb.User) (*xworkpb.User, error) {
	return &xworkpb.User{Id: 11, FirstName: u.FirstName, LastName: u.LastName, Role: u.Role}, nil
}
