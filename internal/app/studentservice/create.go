package studentservice

import (
	"context"
	"github.com/danilashushkanov/student-service/internal/app/studentservice/adapters"
	"github.com/danilashushkanov/student-service/internal/app/teacherservice"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Service) CreateStudent(ctx context.Context, req *api.CreateStudentRequest) (*api.Student, error) {
	if err := validateCreateStudentRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	student, err := s.studentRepository.Create(ctx, adapters.CreateStudentFromPb(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "error create student")
	}

	return adapters.StudentToPb(student), nil
}

func validateCreateStudentRequest(req *api.CreateStudentRequest) error {
	err := validation.Errors{
		"full_name": validation.Validate(strings.TrimSpace(req.GetFullName()), validation.Required, validation.Length(1, 0)),
		"age":       validation.Validate(req.GetAge(), validation.Required),
		"salary":    validation.Validate(req.GetSalary(), validation.Required),
	}.Filter()
	if err != nil {
		return err
	}

	for _, teacher := range req.GetTeachers() {
		err = teacherservice.ValidateCreateTeacherRequest(teacher, false)
		if err != nil {
			return err
		}
	}

	return nil
}
