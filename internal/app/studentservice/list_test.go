package studentservice

import (
	"errors"
	"github.com/danilashushkanov/student/internal/app/studentservice/adapters"
	"github.com/danilashushkanov/student/internal/model"
	api "github.com/danilashushkanov/student/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestListStudent(t *testing.T) {
	t.Run("validate error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListStudentRequest{
			StudentIds: []int64{},
		}

		students, err := te.studentService.ListStudents(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedStudents *api.ListStudentResponse
		assert.Equal(t, expectedStudents, students)
	})

	t.Run("repository Error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListStudentRequest{
			StudentIds: []int64{148},
		}

		expectedMockStudentIds := adapters.ListFilterStudentFromPb(req)
		te.studentRepository.EXPECT().List(te.ctx, expectedMockStudentIds).Return(nil, errors.New("any catalog error"))

		students, err := te.studentService.ListStudents(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedStudents *api.ListStudentResponse
		assert.Equal(t, expectedStudents, students)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.ListStudentRequest{
			StudentIds: []int64{1},
		}

		expectedMockStudentIds := adapters.ListFilterStudentFromPb(req)
		modelStudents := []*model.Student{
			{
				ID:       1,
				FullName: "Павел Жирнов",
				Age:      19,
				Salary:   12345,
			},
		}

		te.studentRepository.EXPECT().List(te.ctx, expectedMockStudentIds).Return(modelStudents, nil)

		resource, err := te.studentService.ListStudents(te.ctx, req)
		assert.NoError(t, err)
		expectedResponse := &api.ListStudentResponse{
			Students: adapters.StudentsToPb(modelStudents),
		}
		assert.Equal(t, expectedResponse, resource)
	})
}
