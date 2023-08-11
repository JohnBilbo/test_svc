package postgre

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"test_svc/internal/model"
	"time"
)

func InitDB(ctx context.Context, cfg *model.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s timezone=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.SSLMode,
		cfg.Database.TimeZone)
	pgConf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}
	pgConf.MaxConns = int32(cfg.Database.MaxCons)
	pgConf.MinConns = int32(cfg.Database.MinCons)
	pgConf.MaxConnIdleTime = time.Duration(int64(cfg.Database.MaxConnIdleTime)) * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, pgConf)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
