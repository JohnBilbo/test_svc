package service

import (
	"github.com/sirupsen/logrus"
	"test_svc/internal/model"
	"test_svc/internal/storage"
)

type service struct {
	cfg     *model.Config
	log     *logrus.Logger
	storage storage.Storage
}

func InitService(cfg *model.Config, log *logrus.Logger, storage storage.Storage) Service {
	return &service{
		cfg:     cfg,
		log:     log,
		storage: storage,
	}
}
