package service

import (
	"errors"

	"github.com/kimjonguns/Nuclear-Go-Make-Unit-Testing-With-Mock/package/entity"
	"github.com/kimjonguns/Nuclear-Go-Make-Unit-Testing-With-Mock/package/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepo
}

func (catSer CategoryService) Get(id string) (*entity.Category, error) {
	category := catSer.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
