package studentservice

import (
	"context"
	"errors"
	"github.com/danilashushkanov/student-service/internal/app/studentservice/adapters"
	"github.com/danilashushkanov/student-service/internal/repository"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetStudent(ctx context.Context, req *api.GetStudentRequest) (*api.Student, error) {
	if err := validateStudentIDRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	s.logger.Info("get student")
	student, err := s.studentRepository.Get(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, repository.ErrEntityNotFound) {
			return nil, status.Error(codes.NotFound, "student not found")
		}
		return nil, status.Error(codes.Internal, "error get student")
	}

	return adapters.StudentToPb(student), nil
}

func validateStudentIDRequest(req *api.GetStudentRequest) error {
	err := validation.Errors{
		"student_id": validation.Validate(req.GetId(), validation.Required),
	}.Filter()
	if err != nil {
		return err
	}
	return nil
}
