package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"test_svc/internal/constants"
	"test_svc/internal/logger"
	"test_svc/internal/model"
	"test_svc/internal/transport"
	"time"
)

type Server struct {
	cfg *model.Config
	log *logrus.Logger
	h   transport.Handler
	app *echo.Echo
}

func InitServer(cfg *model.Config, log *logrus.Logger, h transport.Handler) *Server {
	return &Server{
		cfg: cfg,
		log: log,
		h:   h,
	}
}

func (s *Server) ServerStart(ctx context.Context) {
	s.app = s.echoEngine()
	s.setupRoutes()
	go func() {
		if err := s.app.Start(fmt.Sprintf("%s:%s", s.cfg.Server.Host, s.cfg.Server.Port)); err != nil && err != http.ErrServerClosed {
			logger.FillLoggerFields(s.log, "", constants.LoggerClassServer, "ServerStart").Fatal(err)

		}
	}()
	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := s.app.Shutdown(ctxShutDown); err != nil {
		logger.FillLoggerFields(s.log, "", constants.LoggerClassServer, "Shutdown").Fatal(err)
	}
}

func (s *Server) echoEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
		}))
	return e
}
