//go:build develop

package routing

import (
	"app/pkg/app/dependencies"
	"app/pkg/infrastructure/web"
)

func SetRoutings(server *web.Server, deps *dependencies.Dependencies) {
	setCommonRoutings(server, deps)
}
