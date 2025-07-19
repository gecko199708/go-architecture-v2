package app

import (
	"app/pkg/adapter/web"
	"app/pkg/app/dependencies"
	"app/pkg/app/routing"
	infrastructure "app/pkg/infrastructure/dependencies"
	logic "app/pkg/logic/dependencies"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Version = "builtin"
	Commit  = "unknown"
	Date    = "unknown"
)

func BuildVersion() string {
	var buf bytes.Buffer
	fmt.Fprint(
		&buf,
		fmt.Sprintf("Version: %s\n", Version),
		fmt.Sprintf("Commit: %s\n", Commit),
		fmt.Sprintf("ReleaseAt: %s\n", Date),
	)
	return buf.String()
}

func Run() error {
	server, deps := setup()
	errCh := startServerAsync(server)

	shutdownErr := waitForShutdown(server, errCh)
	cleanupErr := gracefulShutdown(deps)

	return errors.Join(shutdownErr, cleanupErr)
}

// setup initializes dependencies and routing
func setup() (*web.Server, *dependencies.Dependencies) {
	deps := dependencies.New()
	infrastructure.Initialize(deps)
	logic.Initialize(deps)

	server := web.NewServer()
	routing.SetRoutings(server, deps)

	return server, deps
}

// startServerAsync launches server.Run() in a goroutine and returns an error channel
func startServerAsync(server *web.Server) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		err := server.Run()
		if err != nil {
			errCh <- errors.Join(ErrFailureServerStart, err)
		}
		close(errCh)
	}()
	return errCh
}

// waitForShutdown handles OS signals or server errors
func waitForShutdown(server *web.Server, errCh <-chan error) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case sig := <-sigCh:
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Close(ctx); err != nil {
			return errors.Join(ErrFailureGracefulShutdown, err)
		}
		_ = sig // ログに使うならここで出力可能
		return nil
	case err := <-errCh:
		return err
	}
}

// gracefulShutdown closes remaining dependencies
func gracefulShutdown(deps *dependencies.Dependencies) error {
	ctx := context.Background()
	if err := deps.Close(ctx); err != nil {
		return errors.Join(ErrFailureCloseDependencies, err)
	}
	return nil
}
