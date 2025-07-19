package sales

import (
	"app/pkg/app/dependencies"
	"app/pkg/logic/repository"
)

type Repository interface {
	Users() repository.UsersRepository
}

type Logic struct {
}

func NewLogic(deps *dependencies.Dependencies) *Logic {
	return &Logic{}
}
