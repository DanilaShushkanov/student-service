package studentservice

import (
	"github.com/danilashushkanov/student-service/internal/repository"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
)

type Service struct {
	api.UnimplementedStudentServiceServer
	studentRepository repository.StudentRepository
}

func NewService(studentRepository repository.StudentRepository) *Service {
	return &Service{
		studentRepository: studentRepository,
	}
}
