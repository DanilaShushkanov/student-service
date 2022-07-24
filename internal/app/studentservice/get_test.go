package studentservice

import (
	"errors"
	"github.com/danilashushkanov/student-service/internal/app/studentservice/adapters"
	"github.com/danilashushkanov/student-service/internal/model"
	"github.com/danilashushkanov/student-service/internal/repository"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestGetStudent(t *testing.T) {
	t.Run("validation Error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetStudentRequest{
			Id: 0,
		}

		resource, err := te.studentService.GetStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedResponse *api.Student
		assert.Equal(t, expectedResponse, resource)
	})

	t.Run("repository Error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetStudentRequest{
			Id: 1,
		}

		te.studentRepository.EXPECT().Get(te.ctx, req.GetId()).Return(nil, errors.New("any catalog error"))

		resource, err := te.studentService.GetStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedResponse *api.Student
		assert.Equal(t, expectedResponse, resource)
	})

	t.Run("student not found", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetStudentRequest{
			Id: 94,
		}

		te.studentRepository.EXPECT().Get(te.ctx, req.GetId()).Return(nil, repository.ErrEntityNotFound)

		student, err := te.studentService.GetStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.NotFound.String(), status.Code(err).String())
		var expectedResponse *api.Student
		assert.Equal(t, expectedResponse, student)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.GetStudentRequest{
			Id: 1,
		}

		modelStudent := &model.Student{
			ID:       2,
			FullName: "Павел Жирнов",
			Age:      18,
			Salary:   123425,
		}

		te.studentRepository.EXPECT().Get(te.ctx, req.GetId()).Return(modelStudent, nil)

		student, err := te.studentService.GetStudent(te.ctx, req)
		assert.NoError(t, err)
		expectedStudent := adapters.StudentToPb(modelStudent)
		assert.Equal(t, expectedStudent, student)
	})
}
