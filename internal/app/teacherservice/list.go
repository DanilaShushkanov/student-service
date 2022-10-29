package teacherservice

import (
	"context"
	"github.com/danilashushkanov/student-service/internal/app/teacherservice/adapters"
	"github.com/danilashushkanov/student-service/pkg/logging"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListTeachers(ctx context.Context, req *api.ListTeacherRequest) (*api.ListTeacherResponse, error) {
	if err := validateListTeacherRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	logging.GetLogger(ctx).Info("list teacher")
	teacherList, err := s.teacherRepository.List(ctx, adapters.ListFilterTeacherFromPb(req))
	if err != nil {
		return nil, status.Error(codes.Internal, "error list teacher")
	}

	return &api.ListTeacherResponse{
		Teachers: adapters.TeachersToPb(teacherList),
	}, nil
}

func validateListTeacherRequest(req *api.ListTeacherRequest) error {
	err := validation.Errors{
		"teacher_ids": validation.Validate(req.GetTeacherIds(), validation.Required),
	}.Filter()

	if err != nil {
		return err
	}
	return nil
}
