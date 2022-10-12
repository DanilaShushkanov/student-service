package adapters

import (
	"github.com/danilashushkanov/student-service/internal/model"
	"github.com/danilashushkanov/student-service/internal/repository"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"strings"
)

func CreateTeacherFromPb(teacher *api.CreateTeacherRequest) *model.Teacher {
	return &model.Teacher{
		FullName:     strings.TrimSpace(teacher.GetFullName()),
		PositionType: int64(teacher.GetPositionType()),
		StudentID:    teacher.GetStudentId(),
	}
}

func UpdateTeacherFromPb(teacher *api.UpdateTeacherRequest) *model.Teacher {
	return &model.Teacher{
		ID:           teacher.GetId(),
		FullName:     strings.TrimSpace(teacher.GetFullName()),
		PositionType: int64(teacher.GetPositionType()),
	}
}

func ListFilterTeacherFromPb(filter *api.ListTeacherRequest) *repository.TeacherListFilter {
	return &repository.TeacherListFilter{
		IDList:    filter.GetTeacherIds(),
		FieldName: "id",
	}
}
