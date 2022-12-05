package repository

import (
	"github.com/kimjonguns/Nuclear-Go-Make-Unit-Testing-With-Mock/package/entity"
	"github.com/stretchr/testify/mock"
)

type CategoryRepoMock struct {
	Mock mock.Mock
}

func (crm *CategoryRepoMock) FindById(id string) *entity.Category {
	argumens := crm.Mock.Called(id)
	if argumens.Get(0) == nil {
		return nil
	}

	category := argumens.Get(0).(entity.Category)
	return &category
}
