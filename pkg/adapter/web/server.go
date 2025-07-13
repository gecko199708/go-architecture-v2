package web

type Server interface {
	Run() error
	Close() error
}
