package teacherservice

import (
	"github.com/danilashushkanov/student-service/internal/repository"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var positionTypeValidationRule = validation.In(
	api.PositionType_POSTGRADUATE,
	api.PositionType_ASSISTANT,
	api.PositionType_DEAN,
)

type Service struct {
	api.UnimplementedTeacherServiceServer
	teacherRepository repository.TeacherRepository
}

func NewService(teacherRepository repository.TeacherRepository) *Service {
	return &Service{
		teacherRepository: teacherRepository,
	}
}
