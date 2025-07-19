package dependencies

import (
	"app/pkg/common/log"
	"app/pkg/common/repository"
	"context"
)

type Dependencies struct {
	Logger     log.Logger
	Repository repository.Repository
}

func New() *Dependencies {
	deps := new(Dependencies)
	return deps
}

func (deps *Dependencies) Close(ctx context.Context) error {
	return nil
}
