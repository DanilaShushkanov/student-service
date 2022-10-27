package studentservice

import (
	"github.com/danilashushkanov/student-service/internal/repository"
	"github.com/danilashushkanov/student-service/pkg/logging"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
)

type Service struct {
	api.UnimplementedStudentServiceServer
	studentRepository repository.StudentRepository
	logger            logging.Logger
}

func NewService(studentRepository repository.StudentRepository, logger logging.Logger) *Service {
	return &Service{
		studentRepository: studentRepository,
		logger:            logger,
	}
}
