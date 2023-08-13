package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"test_svc/internal/config"
	"test_svc/internal/logger"
	"test_svc/internal/service"
	"test_svc/internal/storage"
	"test_svc/internal/transport"
	"test_svc/internal/transport/http"
)

func main() {
	ctx := context.Background()
	viper.Set("APP.SERVER.HOST", "localhost")
	viper.Set("APP.SERVER.PORT", "8080")
	cfg := config.ReadConfigFromEnv()
	log := logger.InitLogger()
	//db, err := postgre.InitDB(ctx, cfg)
	//if err != nil {
	//	logger.FillLoggerFields(log, "", constants.LoggerClassMain, "main").Error(err)
	//	os.Exit(-1)
	//}
	var db pgxpool.Pool
	Storage := storage.InitStorage(cfg, log, &db)
	Service := service.InitService(cfg, log, Storage)
	handler := transport.InitHandler(cfg, log, Service)
	server := http.InitServer(cfg, log, handler)
	server.ServerStart(ctx)
}
