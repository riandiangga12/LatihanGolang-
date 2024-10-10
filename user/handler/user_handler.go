package handler

import (
	"LatihanGolang/entity"
	pb "LatihanGolang/user/proto/user_service/v1"
	"LatihanGolang/user/service"
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, _ *emptypb.Empty) (*pb.GetUsersResponse, error) {
	users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var userProto []*pb.User
	for _, user := range users {
		userProto = append(userProto, &pb.User{
			Id:        int32(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		})
	}

	return &pb.GetUsersResponse{
		Users: userProto,
	}, nil
}

func (h *UserHandler) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := h.userService.GetUserByID(ctx, int(req.Id))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := &pb.GetUserByIDResponse{
		User: &pb.User{
			Id:        int32(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}

	return res, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.MutationResponse, error) {
	createdUser, err := h.userService.CreateUser(ctx, &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("User with ID %d has been created", createdUser.ID),
		Id:      int32(createdUser.ID),
	}, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.MutationResponse, error) {
	updatedUser, err := h.userService.UpdateUser(ctx, int(req.Id), entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("User with ID %d has been updated", updatedUser.ID),
	}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.MutationResponse, error) {
	err := h.userService.DeleteUser(ctx, int(req.Id))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.MutationResponse{
		Message: fmt.Sprintf("User with ID %d has been deleted", req.Id),
	}, nil
}
