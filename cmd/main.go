package main

import (
	"app/pkg/adapter/web/routing"
	"app/pkg/app/dependencies"
	"app/pkg/common/constant"
	infrastructure "app/pkg/infrastructure/dependencies"
	logic "app/pkg/logic/dependencies"
	"os"
)

var (
	Version = "builtin"
	Commit  = "unknown"
	Date    = "unknown"
)

func main() {
	deps := dependencies.New()
	infrastructure.Initialize(deps)
	logic.Initialize(deps)
	routing.SetRoutings(deps.Server, deps)

	Run(deps)
}

func Run(deps *dependencies.Dependencies) {
	server := deps.Server
	if err := server.Run(); err != nil {
		os.Exit(constant.ExitCodeFailureServerStart)
	}
	if err := server.Close(); err != nil {
		os.Exit(constant.ExitCodeFailureServerClose)
	}
}
