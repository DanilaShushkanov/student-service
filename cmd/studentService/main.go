package main

import (
	"context"
	"github.com/caarlos0/env"
	"github.com/danilashushkanov/student-service/internal/app"
	"github.com/danilashushkanov/student-service/internal/config"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal(err)
	}

	cfg := &config.Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Failed to retrieve env variables, %v", err)
		return
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatal("error running grpc server ", err)
	}
}
