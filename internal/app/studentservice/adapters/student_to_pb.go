package adapters

import (
	"github.com/danilashushkanov/student/internal/app/teacherservice/adapters"
	"github.com/danilashushkanov/student/internal/model"
	api "github.com/danilashushkanov/student/pkg/studentServiceApi"
)

func StudentToPb(student *model.Student) *api.Student {
	return &api.Student{
		Id:       student.ID,
		FullName: student.FullName,
		Age:      student.Age,
		Salary:   student.Salary,
		Teachers: adapters.TeachersToPb(student.Teachers),
	}
}

func StudentsToPb(students []*model.Student) []*api.Student {
	apiStudents := make([]*api.Student, 0, len(students))

	for _, item := range students {
		apiStudents = append(apiStudents, StudentToPb(item))
	}
	return apiStudents
}
