package teacherservice

import (
	"errors"
	"github.com/danilashushkanov/student-service/internal/app/teacherservice/adapters"
	"github.com/danilashushkanov/student-service/internal/model"
	api "github.com/danilashushkanov/student-service/pkg/studentServiceApi"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestUpdateTeacher(t *testing.T) {
	t.Run("validate error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateTeacherRequest{
			Id:           0,
			FullName:     "",
			PositionType: 0,
		}

		teacher, err := te.teacherService.PatchTeacher(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedTeacher *api.Teacher
		assert.Equal(t, expectedTeacher, teacher)
	})

	t.Run("repository error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateTeacherRequest{
			Id:           1,
			FullName:     "Паша Жирнов",
			PositionType: 1,
		}

		expectedMockStudent := adapters.UpdateTeacherFromPb(req)
		te.teacherRepository.EXPECT().Update(te.ctx, expectedMockStudent).Return(nil, errors.New("any catalog error"))

		teacher, err := te.teacherService.PatchTeacher(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedTeacher *api.Teacher
		assert.Equal(t, expectedTeacher, teacher)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.UpdateTeacherRequest{
			Id:           1,
			FullName:     "Паша Жирнов",
			PositionType: 1,
		}

		expectedMockTeacher := adapters.UpdateTeacherFromPb(req)
		modelTeacher := &model.Teacher{
			ID:           17,
			FullName:     "name",
			PositionType: 1,
			StudentID:    1,
		}

		te.teacherRepository.EXPECT().Update(te.ctx, expectedMockTeacher).Return(modelTeacher, nil)

		teacher, err := te.teacherService.PatchTeacher(te.ctx, req)
		assert.NoError(t, err)
		expectedTeacher := adapters.TeacherToPb(modelTeacher)
		assert.Equal(t, expectedTeacher, teacher)
	})
}
