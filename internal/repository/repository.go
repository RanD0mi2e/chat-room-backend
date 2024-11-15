package repository

import (
	"chat-app-backend-service/pkg/log"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db     *sqlx.DB
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *sqlx.DB) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}
