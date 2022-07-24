package repository

import (
	"context"
	"github.com/danilashushkanov/student/internal/model"
)

type StudentRepository interface {
	Create(context.Context, *model.Student) (*model.Student, error)
	List(context.Context, *StudentListFilter) ([]*model.Student, error)
	Get(context.Context, int64) (*model.Student, error)
	Update(context.Context, *model.Student) (*model.Student, error)
	Delete(context.Context, int64) error
}

type TeacherRepository interface {
	Get(context.Context, int64) (*model.Teacher, error)
	Create(context.Context, []*model.Teacher) ([]*model.Teacher, error)
	Update(context.Context, *model.Teacher) (*model.Teacher, error)
	UpdateTeachers(context.Context, int64, []*model.Teacher) ([]*model.Teacher, error)
	DeleteByStudentID(context.Context, int64) error
	List(context.Context, *TeacherListFilter) ([]*model.Teacher, error)
}
