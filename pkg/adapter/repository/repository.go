package repository

type Tx interface {
	Commit() error
	Rollback() error
}

type Repository interface {
	Begin() (Tx, error)

	Users() interface {
	}
}
