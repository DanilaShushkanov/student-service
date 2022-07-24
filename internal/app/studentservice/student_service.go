package studentservice

import (
	"github.com/danilashushkanov/student/internal/repository"
	api "github.com/danilashushkanov/student/pkg/studentServiceApi"
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
