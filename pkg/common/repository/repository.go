package repository

import (
	"app/pkg/logic/repository"
	"context"
)

type Query interface{}

type Tx interface {
	Commit() error
	Rollback() error
}

type Repository interface {
	Begin() (Tx, error)

	Users() repository.UsersRepository
}

type Selector[T any] interface {
	Select(ctx context.Context, tx Tx, query Query) (T, error)
}

type Updater[T any] interface {
	Update(ctx context.Context, tx Tx, query Query) (T, error)
}

type Inserter[T any] interface {
	Insert(ctx context.Context, tx Tx, query Query) (T, error)
}

type Deleter[T any] interface {
	Delete(ctx context.Context, tx Tx, query Query) (T, error)
}
