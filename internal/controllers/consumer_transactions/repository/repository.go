package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/constants"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"
	"github.com/redis/go-redis/v9"
)

type Repository interface {
	Querier
	RedisRepository

	// Create Consumer Transaction with Database Transaction for better consistency
	CreateTx(ctx context.Context, arg *CreateTxParams) (CreateTxResult, error)

	// Update Consumer Transaction with Database Transaction
	UpdateTx(ctx context.Context, arg *UpdateTxParams) (UpdateTxResult, error)
}

type Store struct {
	log logger.Logger
	cfg *config.Config
	db  *sql.DB
	*Queries
	*RedisRepositoryImpl
}

func NewStore(
	log logger.Logger,
	cfg *config.Config,
	dbConnection *sql.DB,
	redisConnection redis.UniversalClient,
) Repository {
	log = log.WithPrefix(fmt.Sprintf("%s-%s", "consumer-transactions", constants.Repository))

	return &Store{
		log:                 log,
		cfg:                 cfg,
		db:                  dbConnection,
		Queries:             New(dbConnection),
		RedisRepositoryImpl: NewRedisRepositoryImpl(log, cfg, redisConnection),
	}
}

func (r *Store) execTx(ctx context.Context, fnCallback func(*Queries) error) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelReadCommitted,
	})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fnCallback(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %w, rb err: %w", err, rbErr)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
