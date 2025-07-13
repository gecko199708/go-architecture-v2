//go:build develop

package routing

import (
	"app/pkg/adapter/web"
	"app/pkg/app/dependencies"
)

func SetRoutings(server web.Server, deps *dependencies.Dependencies) {
	setCommonRoutings(server, deps)
}
