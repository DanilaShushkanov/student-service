package teacherservice

import (
	"context"
	"github.com/danilashushkanov/student-service/internal/app/teacherservice/adapters"
	"github.com/danilashushkanov/student-service/internal/model"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

func (s *Service) CreateTeacher(ctx context.Context, req *api.CreateTeacherRequest) (*api.Teacher, error) {
	if err := ValidateCreateTeacherRequest(req, true); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	modelTeacher := adapters.CreateTeacherFromPb(req)
	teachers, err := s.teacherRepository.Create(ctx, []*model.Teacher{modelTeacher})
	if err != nil {
		return nil, status.Error(codes.Internal, "error create teacher")
	}

	return adapters.TeacherToPb(teachers[len(teachers)-1:][0]), nil
}

func ValidateCreateTeacherRequest(req *api.CreateTeacherRequest, checkStudentID bool) error {
	checkMap := validation.Errors{
		"full_name":     validation.Validate(strings.TrimSpace(req.GetFullName()), validation.Required, validation.Length(1, 0)),
		"position_type": validation.Validate(req.GetPositionType(), positionTypeValidationRule),
	}

	if checkStudentID {
		checkMap["student_id"] = validation.Validate(req.GetStudentId(), validation.Required)
	}

	if err := checkMap.Filter(); err != nil {
		return err
	}
	return nil
}
