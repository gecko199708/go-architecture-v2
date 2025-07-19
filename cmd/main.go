package main

import (
	"app/pkg/app"
	"errors"
	"fmt"
	"os"
)

const (
	ExitCodeUnknown = 1

	// failure starting server
	ExitCodeFailureServerStart = 2

	// failure closing server
	ExitCodeFailureServerClose = 3
)

func isVersionCommand() bool {
	return len(os.Args) > 1 && os.Args[1] == "version"
}

func main() {
	if isVersionCommand() {
		fmt.Println(app.BuildVersion())
		return
	}

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal: %v\n", err)

		var exitCode int
		switch {
		case errors.Is(err, app.ErrFailureServerStart):
			exitCode = ExitCodeFailureServerStart
		case errors.Is(err, app.ErrFailureCloseDependencies):
			exitCode = ExitCodeFailureServerClose
		default:
			exitCode = ExitCodeUnknown
		}
		os.Exit(exitCode)
	}
}
