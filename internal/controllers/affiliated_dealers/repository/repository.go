package repository

import (
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
	log = log.WithPrefix(fmt.Sprintf("%s-%s", "affiliated-dealers", constants.Repository))

	return &Store{
		log:                 log,
		cfg:                 cfg,
		db:                  dbConnection,
		Queries:             New(dbConnection),
		RedisRepositoryImpl: NewRedisRepositoryImpl(log, cfg, redisConnection),
	}
}
