package dependencies

import "app/pkg/adapter/web"

type Dependencies struct {
	Server web.Server
}

func New() *Dependencies {
	deps := new(Dependencies)
	return deps
}
