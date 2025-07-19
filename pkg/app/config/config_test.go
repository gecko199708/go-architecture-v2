package config_test

import (
	"app/pkg/app/config"
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Setenv("DB_DRIVER", "mydb")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("MAINTENANCE", "1")

	cfg := config.Load()
	fmt.Printf("database.Driver: %s\n", cfg.Database.Driver)
	fmt.Printf("database.Host: %s\n", cfg.Database.Host)
	fmt.Printf("database.Port: %d\n", cfg.Database.Port)
	fmt.Printf("Maintenance: %v\n", cfg.Maintenance)
	fmt.Printf("server.Resource: %v\n", cfg.Server.Resource)
}
