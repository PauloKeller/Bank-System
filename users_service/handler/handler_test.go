package handler

import (
	"context"
	"testing"
	"users_service/interactors"
	core_grpc_api "users_service/proto"
)

func TestCreateUser(t *testing.T) {
	handler := &Handler{
		CreateUserInteractor: &interactors.MockCreateUserInteractor{},
	}

	reply, err := handler.Create(context.TODO(), &core_grpc_api.CreateUserRequest{})

	if err != nil {
		t.Errorf("Test failed. %v.", err)
	}

	if reply == nil {
		t.Errorf("Test failed. Should reply a user.")
	}
}

func TestFailToCreateUser(t *testing.T) {
	handler := &Handler{
		CreateUserInteractor: &interactors.MockCreateUserInteractor{},
	}

	reply, err := handler.Create(context.TODO(), &core_grpc_api.CreateUserRequest{
		Username: "fail",
	})

	if reply != nil {
		t.Errorf("Test failed. Shouldn't reply a user.")
	}

	if err == nil {
		t.Errorf("Test failed. Should return a error.")
	}
}
