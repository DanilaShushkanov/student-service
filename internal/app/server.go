package app

import (
	"context"
	"fmt"
	"github.com/danilashushkanov/student-service/internal/app/studentservice"
	"github.com/danilashushkanov/student-service/internal/app/teacherservice"
	"github.com/danilashushkanov/student-service/internal/bootstrap"
	"github.com/danilashushkanov/student-service/internal/closer"
	"github.com/danilashushkanov/student-service/internal/config"
	"github.com/danilashushkanov/student-service/internal/repository"
	"github.com/danilashushkanov/student-service/pkg/logging"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, cfg *config.Config, logger logging.Logger) error {
	s := grpc.NewServer()

	_, cancel := context.WithCancel(ctx)

	logger.Infof("listen GRPC server to %s", cfg.GRPCAddr)
	l, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		cancel()
		logger.Fatalf("failed to listen tcp %s, %v\n", cfg.GRPCAddr, err)
	}

	initServices(s, cfg, logger)

	go func() {
		if err = s.Serve(l); err != nil {
			logger.Fatal("ERROR: ", err.Error())
		}
	}()

	gracefulShutdown(s, cancel)
	return nil
}

func initServices(s *grpc.Server, cfg *config.Config, logger logging.Logger) {
	logger.Info("start connect to DB")
	conn, err := bootstrap.InitDB(cfg)
	if err != nil {
		logger.Fatalf("not connect to db :%v", err)
	}

	logger.Info("register services")
	api.RegisterStudentServiceServer(s, studentservice.NewService(
		repository.NewStudentRepository(
			conn,
			repository.NewTeacherRepository(conn),
		),
		logger,
	),
	)
	api.RegisterTeacherServiceServer(s, teacherservice.NewService(
		repository.NewTeacherRepository(conn),
		logger,
	),
	)
}

func gracefulShutdown(s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	sig := <-ch
	errorMessage := fmt.Sprintf("%s %v - %s", "Received shutdown signal:", sig, "Graceful shutdown done")
	fmt.Println(errorMessage)
	s.GracefulStop()
	cancel()
	closer.CloseAll()
}
