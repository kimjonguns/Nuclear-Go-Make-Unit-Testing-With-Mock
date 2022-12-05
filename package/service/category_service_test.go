package service

import (
	"testing"

	"github.com/kimjonguns/Nuclear-Go-Make-Unit-Testing-With-Mock/package/entity"
	"github.com/kimjonguns/Nuclear-Go-Make-Unit-Testing-With-Mock/package/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepoMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "High Class",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	data, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, category.Id, data.Id)
}
