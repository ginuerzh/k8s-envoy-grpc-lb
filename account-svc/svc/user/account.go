package user

import (
	"context"
	"fmt"

	"github.com/ginuerzh/k8s-envoy-grpc-lb/account-svc/api"
)

type Server struct{}

func (s *Server) Register(ctx context.Context, req *api.RegisterRequest) (*api.RegisterResponse, error) {
	user := &api.UserDetail{
		Id:     req.GetId(),
		Name:   host,
		Age:    18,
		Avatar: "http://example.com/avatar.jpg",
	}
	return user, nil
}

func (s *Server) List(ctx context.Context, req *api.UserListRequest) (*api.UserListResponse, error) {
	resp := &api.UserListResponse{}
	for i := 0; i < 3; i++ {
		resp.Users = append(resp.Users, &api.UserDetail{
			Id:     fmt.Sprintf("user%d", i),
			Name:   fmt.Sprintf("User%d", i),
			Age:    int32(10 + i),
			Avatar: fmt.Sprintf("http://www.example.com/avatar%d.jpg", i),
		})
	}
	return resp, nil
}

func (s *Server) Update(ctx context.Context, req *api.UserUpdateRequest) (*api.UserDetail, error) {
	user := req.User
	if user == nil {
		user = &api.UserDetail{}
	}
	user.Id = req.Id

	return user, nil
}
