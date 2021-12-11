package handler

import (
	"context"
	"log"

	"users_service/dtos"
	"users_service/interactors"
	core_grpc_api "users_service/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Handler struct {
	CreateUserInteractor  interactors.CreateUserInteractorInterface
	GetAllUserInteractor  interactors.GetAllUserInteractorInterface
	GetUserByIDInteractor interactors.GetUserByIDInteractorInterface
}

type HandlerInterface interface {
	Create(context.Context, *core_grpc_api.CreateUserRequest) (*core_grpc_api.CreateUserReply, error)
	GetAll(*empty.Empty, core_grpc_api.UsersService_GetAllServer) error
	GetById(context.Context, *core_grpc_api.GetUserById) (*core_grpc_api.User, error)
}

func (handler *Handler) Create(ctx context.Context, in *core_grpc_api.CreateUserRequest) (*core_grpc_api.CreateUserReply, error) {
	model := &dtos.CreateUserDto{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
	}

	err := handler.CreateUserInteractor.Create(model)
	if err != nil {
		log.Println(err.Error())
		grpcError := status.Error(codes.Code(err.Code.ConvertToGrpc()), err.Error())
		return nil, grpcError
	}

	response := &core_grpc_api.CreateUserReply{
		WasCreated: true,
	}

	return response, nil
}

func (handler *Handler) GetAll(in *empty.Empty, stream core_grpc_api.UsersService_GetAllServer) error {
	data, err := handler.GetAllUserInteractor.GetAll()

	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, value := range data {
		if err := stream.Send(&core_grpc_api.User{
			Id:        value.ID.String(),
			FirstName: value.FirstName,
			LastName:  value.LastName,
			Username:  value.Username,
			Email:     value.Email,
			Password:  value.Password,
			CreatedAt: timestamppb.New(value.CreatedAt),
			UpdatedAt: timestamppb.New(value.UpdatedAt),
			DeletedAt: timestamppb.New(value.DeletedAt),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (handler *Handler) GetById(ctx context.Context, in *core_grpc_api.GetUserById) (*core_grpc_api.User, error) {
	data, err := handler.GetUserByIDInteractor.GetByID(in.GetId())

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	response := &core_grpc_api.User{
		Id:        data.ID.String(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: timestamppb.New(data.CreatedAt),
		UpdatedAt: timestamppb.New(data.UpdatedAt),
		DeletedAt: timestamppb.New(data.DeletedAt),
	}

	return response, nil
}
