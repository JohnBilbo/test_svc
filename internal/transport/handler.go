package transport

import (
	"github.com/sirupsen/logrus"
	"test_svc/internal/model"
	"test_svc/internal/service"
)

type handler struct {
	cfg *model.Config
	log *logrus.Logger
	svc service.Service
}

func InitHandler(cfg *model.Config, log *logrus.Logger, svc service.Service) Handler {
	return &handler{
		cfg: cfg,
		log: log,
		svc: svc,
	}
}
