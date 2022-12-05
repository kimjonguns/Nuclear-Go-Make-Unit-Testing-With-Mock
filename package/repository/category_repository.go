package repository

import "github.com/kimjonguns/Nuclear-Go-Make-Unit-Testing-With-Mock/package/entity"

type CategoryRepo interface {
	FindById(Id string) *entity.Category
}
