package repository

import (
	"chat-app-backend-service/pkg/log"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
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

type Transaction interface {
	WithTransaction(fn func(tx *sqlx.Tx) error) error
	Exec(query string, args ...interface{}) (int64, error)
}

func NewTransation(db *Repository) Transaction {
	return db
}

func (r *Repository) WithTransaction(fn func(tx *sqlx.Tx) error) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("启动事务失败: %w", err)
	}

	err = fn(tx)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			r.logger.Error("事务回滚失败", zap.Error(rollbackErr))
		}
		return fmt.Errorf("事务执行失败: %w", err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return fmt.Errorf("事务提交失败: %w", err)
	}

	return nil
}

func (r *Repository) Exec(query string, args ...interface{}) (int64, error) {
	result, err := r.db.Exec(query, args...)
	if err != nil {
		r.logger.Error("Exec函数执行异常", zap.Error(err))
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error("Exec函数执行异常", zap.Error(err))
		return 0, err
	}

	return rowsAffected, nil
}
