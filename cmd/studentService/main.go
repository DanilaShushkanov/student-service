package main

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/danilashushkanov/student-service/internal/app"
	"github.com/danilashushkanov/student-service/internal/config"
	"github.com/danilashushkanov/student-service/pkg/logging"
	"github.com/joho/godotenv"
)

func main() {
	logger := logging.NewLogger()
	logger.Info("start application")

	logger.Info("start parsing environments")
	err := godotenv.Load("../../.env")
	if err != nil {
		logger.Fatal(err)
	}

	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		logger.Fatalf("Failed to retrieve env variables, %v", err)
		return
	}

	logger.Info("run application")
	if err := app.Run(context.Background(), cfg, logger); err != nil {
		logger.Fatal("error running grpc server ", err)
	}
}
