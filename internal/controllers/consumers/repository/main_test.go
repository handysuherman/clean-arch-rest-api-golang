package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/config"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/databases/mysql"
	"github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/logger"

	redisDb "github.com/handysuherman/studi-kasus-pt-xyz-golang-developer/internal/pkg/databases/redis"
	"github.com/redis/go-redis/v9"
)

var (
	testStore       Repository
	cfg             *config.Config
	tlog            logger.Logger
	dbConnection    *sql.DB
	redisConnection redis.UniversalClient
)

func TestMain(m *testing.M) {
	_tlog := logger.NewLogger()
	// log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	tlog = _tlog

	ctx := context.Background()

	wd, err := os.Getwd()
	if err != nil {
		tlog.Warnf("os.Getwd.err: %v", err)
		return
	}

	_cfg, err := config.New(fmt.Sprintf("%s/%s", findModuleRoot(wd), "config-dev.yaml"))
	if err != nil {
		tlog.Warnf("config.New.err: %v", err)
		return
	}
	cfg = _cfg

	mysqlConn()
	redisConn(ctx)

	testStore = NewStore(tlog, cfg, dbConnection, redisConnection)
	os.Exit(m.Run())
}

func mysqlConn() {
	opt := &mysql.Config{
		Driver: cfg.Databases.MySQL.Driver,
		Source: fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true",
			cfg.Databases.MySQL.Username,
			cfg.Databases.MySQL.Password,
			cfg.Databases.MySQL.Host,
			cfg.Databases.MySQL.Port,
			cfg.Databases.MySQL.DBName,
		),
	}

	conn, err := mysql.New(opt)
	if err != nil {
		tlog.Warnf("mysql.mysql.New.err: %v", err)
		return
	}

	dbConnection = conn
}

func redisConn(ctx context.Context) {
	opt := redisDb.Config{
		Host:      cfg.Databases.Redis.Servers[0],
		Password:  cfg.Databases.Redis.Password,
		DB:        cfg.Databases.Redis.DB,
		PoolSize:  cfg.Databases.Redis.PoolSize,
		TLsEnable: cfg.Databases.Redis.EnableTLS,
	}

	conn, err := redisDb.NewUniversalRedisClient(ctx, &opt)
	if err != nil {
		tlog.Warnf("redis.redisDb.NewUniversalRedisClient.err: %v", err)
		return
	}

	redisConnection = conn
}

func findModuleRoot(dir string) string {
	for {
		_, err := os.Stat(filepath.Join(dir, "go.mod"))
		if err == nil {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return ""
}
