package adapters

import (
	"github.com/danilashushkanov/student/internal/model"
	api "github.com/danilashushkanov/student/pkg/studentServiceApi"
)

func TeacherToPb(teacher *model.Teacher) *api.Teacher {
	return &api.Teacher{
		Id:           teacher.ID,
		FullName:     teacher.FullName,
		PositionType: api.PositionType(teacher.PositionType),
		StudentId:    teacher.StudentID,
	}
}

func TeachersToPb(teachers []*model.Teacher) []*api.Teacher {
	apiTeachers := make([]*api.Teacher, 0, len(teachers))
	for _, teacher := range teachers {
		apiTeacher := TeacherToPb(teacher)
		apiTeachers = append(apiTeachers, apiTeacher)
	}
	return apiTeachers
}
