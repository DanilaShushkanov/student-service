package main

import (
	"context"
	"github.com/danilashushkanov/student-service/internal/app"
	"github.com/danilashushkanov/student-service/internal/config"
	"github.com/danilashushkanov/student-service/pkg/logging"
)

func main() {
	ctx := context.Background()
	logger := logging.GetLogger(ctx)
	logger.Info("start application")

	logger.Info("init config")
	cfg, err := config.InitConfig()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("run application")
	if err := app.Run(ctx, cfg, logger); err != nil {
		logger.Fatal("error running grpc server ", err)
	}
}
