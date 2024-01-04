package grpcHandlers

import (
	"github.com/Egor123qwe/GraphQL-project/internal/models"
	"github.com/Egor123qwe/GraphQL-project/internal/scenarios"
	"github.com/Egor123qwe/GraphQL-project/proto/api/generate/desc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"context"
)

type Handlers struct {
	scenarios scenarios.User
	desc.UnimplementedUserServiceServer
}

func New(scenarios scenarios.User) *Handlers {
	return &Handlers{
		scenarios: scenarios,
	}
}

func (h *Handlers) CreateUser(ctx context.Context, usr *desc.UserData) (*desc.UserAccessInfo, error) {
	newUser, err := h.scenarios.CreateUser(
		ctx,
		&models.User{
			Email:      usr.GetEmail(),
			Name:       usr.GetName(),
			Age:        usr.GetAge(),
			University: usr.GetUniversity(),
			Course:     usr.GetCourse(),
			Hobbies:    usr.GetHobbies(),
		},
	)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &desc.UserAccessInfo{
		Id: newUser.Id,
	}, nil
}

func (h *Handlers) DeleteUser(ctx context.Context, in *desc.UserRequest) (*emptypb.Empty, error) {

	//if err := h.scenarios.DeleteUser(ctx, user.Id); err != nil {
	//	return nil, status.Error(codes.InvalidArgument, err.Error())
	//}
	return &emptypb.Empty{}, nil
}
