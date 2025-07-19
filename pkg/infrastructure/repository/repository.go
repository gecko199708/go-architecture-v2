package repository

import "app/pkg/infrastructure/database"

type Repository struct {
	db *database.Database
}
