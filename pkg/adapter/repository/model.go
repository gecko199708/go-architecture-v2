package repository

type Model[T any] interface {
	Entity() T
	Migrate(entity T) error
}
