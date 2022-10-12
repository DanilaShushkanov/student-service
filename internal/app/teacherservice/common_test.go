package teacherservice

import (
	"context"
	"github.com/danilashushkanov/student-service/internal/repository"
	"github.com/golang/mock/gomock"
	"testing"
)

type testEnv struct {
	ctx  context.Context
	ctrl *gomock.Controller

	teacherRepository *repository.MockTeacherRepository

	teacherService *Service
}

func newTestEnv(t *testing.T) *testEnv {
	tEnv := &testEnv{}
	tEnv.ctx = context.Background()
	tEnv.ctrl = gomock.NewController(t)

	tEnv.teacherRepository = repository.NewMockTeacherRepository(tEnv.ctrl)

	tEnv.teacherService = NewService(tEnv.teacherRepository)
	return tEnv
}
