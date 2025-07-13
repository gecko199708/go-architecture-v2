package sales

import (
	"app/pkg/app/dependencies"
	"app/pkg/common/helper/repository"
	"app/pkg/entity"
)

type Repository interface {
	Users() interface {
		repository.Selector[entity.User]
	}
}

type Logic struct {
}

func NewLogic(deps *dependencies.Dependencies) *Logic {
	return &Logic{}
}
