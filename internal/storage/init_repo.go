package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"test_svc/internal/model"
)

type storage struct {
	cfg *model.Config
	log *logrus.Logger
	db  *pgxpool.Pool
}

func InitStorage(cfg *model.Config, log *logrus.Logger, db *pgxpool.Pool) Storage {
	return &storage{
		cfg: cfg,
		log: log,
		db:  db,
	}
}
