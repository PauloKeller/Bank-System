package interactors

import (
	"testing"
)

func TestFailToGetAll(t *testing.T) {
	interactor := &GetAllUserInteractor{
		repository: &MockUserRepository{
			ShouldFailGetAll: true,
		},
	}

	data, err := interactor.GetAll()

	if len(data) > 0 {
		t.Errorf("Test failed. Shouldn't return data.")
	}

	if err == nil {
		t.Errorf("Test failed. Should return error.")
	}
}

func TestGetAll(t *testing.T) {
	interactor := &GetAllUserInteractor{
		repository: &MockUserRepository{},
	}

	data, err := interactor.GetAll()

	if err != nil {
		t.Errorf("Test failed. Shouldn't return error.")
	}

	if len(data) <= 0 {
		t.Errorf("Test failed. Should return data.")
	}
}
