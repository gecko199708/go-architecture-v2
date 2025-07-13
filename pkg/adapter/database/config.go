package database

type Config interface {
	Driver() string
}
