package studentservice

import (
	"context"
	"github.com/danilashushkanov/student/internal/app/studentservice/adapters"
	"github.com/danilashushkanov/student/internal/app/teacherservice"
	api "github.com/danilashushkanov/student/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Service) PatchStudent(ctx context.Context, req *api.UpdateStudentRequest) (*api.Student, error) {
	if err := validateUpdateStudentRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	student, err := s.studentRepository.Update(ctx, adapters.UpdateStudentFromPb(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "error update student")
	}

	return adapters.StudentToPb(student), nil
}

func validateUpdateStudentRequest(req *api.UpdateStudentRequest) error {
	err := validation.Errors{
		"id":        validation.Validate(req.GetId(), validation.Required),
		"full_name": validation.Validate(strings.TrimSpace(req.GetFullName()), validation.Required, validation.Length(1, 0)),
		"age":       validation.Validate(req.GetAge(), validation.Required),
		"salary":    validation.Validate(req.GetSalary(), validation.Required),
	}.Filter()
	if err != nil {
		return err
	}

	for _, teacher := range req.GetTeachers() {
		err = teacherservice.ValidateUpdateTeacherRequest(teacher)
		if err != nil {
			return err
		}
	}
	return nil
}
