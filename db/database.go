package db

import (
	"context"
	"time"

	"bicomsystems.com/network/remote-agent/config"
	"bicomsystems.com/network/remote-agent/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Pool   *pgxpool.Pool
	Logger *logger.Logger
}

func NewDatabase(cnfg config.Config, log *logger.Logger) (db *Database, err error) {
	dbCnfg, err := pgxpool.ParseConfig(cnfg.DbUrl)
	if err != nil {
		log.Error().Err(err).Str("source", "database.go").Send()
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbCnfg)
	if err != nil {
		log.Error().Err(err).Str("source", "database.go").Send()
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		log.Error().Err(err).Str("source", "database.go").Send()
		return nil, err
	}

	log.Info().Str("source", "database.go").Msg("pgxpool successfuly connected to the database connectiion pool")
	return &Database{Pool: pool, Logger: log}, nil
}

// Close closes the database connection pool.
func (d *Database) Close() {
	log.Info().Str("source", "database.go").Msg("pgxpool successfuly closed database connection pool")
	d.Pool.Close()
}
