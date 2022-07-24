package studentservice

import (
	"errors"
	"fmt"
	"github.com/danilashushkanov/student/internal/app/studentservice/adapters"
	"github.com/danilashushkanov/student/internal/model"
	api "github.com/danilashushkanov/student/pkg/studentServiceApi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	t.Run("validation Error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateStudentRequest{
			FullName: "",
			Age:      0,
			Salary:   0,
			Teachers: []*api.CreateTeacherRequest{
				{
					FullName:     "",
					PositionType: 11,
					StudentId:    0,
				},
			},
		}

		student, err := te.studentService.CreateStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.InvalidArgument, status.Code(err))
		var expectedStudent *api.Student
		assert.Equal(t, expectedStudent, student)
	})

	t.Run("repository Error", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateStudentRequest{
			FullName: "name",
			Age:      12,
			Salary:   123123,
			Teachers: []*api.CreateTeacherRequest{
				{
					FullName:     "name",
					PositionType: 1,
					StudentId:    1,
				},
			},
		}

		expectedMockStudent := adapters.CreateStudentFromPb(req)
		te.studentRepository.EXPECT().Create(te.ctx, expectedMockStudent).Return(nil, errors.New("any catalog error"))

		student, err := te.studentService.CreateStudent(te.ctx, req)
		assert.Error(t, err)
		assert.Equal(t, codes.Internal.String(), status.Code(err).String())
		var expectedStudent *api.Student
		assert.Equal(t, expectedStudent, student)
	})

	t.Run("success", func(t *testing.T) {
		te := newTestEnv(t)

		req := &api.CreateStudentRequest{
			FullName: "name",
			Age:      12,
			Salary:   123123,
			Teachers: []*api.CreateTeacherRequest{
				{
					FullName:     "name",
					PositionType: 1,
					StudentId:    1,
				},
			},
		}

		expectedMockStudent := adapters.CreateStudentFromPb(req)
		modelStudent := &model.Student{
			ID:       17,
			FullName: "name",
			Age:      12,
			Salary:   123123,
			Teachers: []*model.Teacher{
				{
					ID:           1,
					FullName:     "name",
					PositionType: 1,
					StudentID:    1,
				},
			},
		}

		te.studentRepository.EXPECT().Create(te.ctx, expectedMockStudent).Return(modelStudent, nil)

		student, err := te.studentService.CreateStudent(te.ctx, req)
		assert.NoError(t, err)
		expectedStudent := adapters.StudentToPb(modelStudent)
		assert.Equal(t, expectedStudent, student)
	})
}

func TestValidateCreateStudentRequest(t *testing.T) {
	type args struct {
		req *api.CreateStudentRequest
	}

	tests := []struct {
		name       string
		args       args
		errorField string
		errorCode  string
	}{
		{
			name: "fullName no filled",
			args: args{
				&api.CreateStudentRequest{
					FullName: "",
					Age:      12,
					Salary:   123333,
					Teachers: []*api.CreateTeacherRequest{
						{
							FullName:     "name",
							PositionType: 1,
							StudentId:    1,
						},
					},
				},
			},
			errorField: "full_name",
			errorCode:  validation.ErrRequired.Code(),
		},
		{
			name: "age no filled",
			args: args{
				&api.CreateStudentRequest{
					FullName: "Шушканов Данила",
					Age:      0,
					Salary:   123333,
					Teachers: []*api.CreateTeacherRequest{
						{
							FullName:     "name",
							PositionType: 1,
							StudentId:    1,
						},
					},
				},
			},
			errorField: "age",
			errorCode:  validation.ErrRequired.Code(),
		},
		{
			name: "salary no filled",
			args: args{
				&api.CreateStudentRequest{
					FullName: "Шушканов Данила",
					Age:      12,
					Salary:   0,
					Teachers: []*api.CreateTeacherRequest{
						{
							FullName:     "name",
							PositionType: 1,
							StudentId:    1,
						},
					},
				},
			},
			errorField: "salary",
			errorCode:  validation.ErrRequired.Code(),
		},
		{
			name: "teacher full_name not filled",
			args: args{
				&api.CreateStudentRequest{
					FullName: "Шушканов Данила",
					Age:      12,
					Salary:   12,
					Teachers: []*api.CreateTeacherRequest{
						{
							FullName:     "",
							PositionType: 1,
							StudentId:    1,
						},
					},
				},
			},
			errorField: "full_name",
			errorCode:  validation.ErrRequired.Code(),
		},
		{
			name: "teacher PositionType Invalid",
			args: args{
				&api.CreateStudentRequest{
					FullName: "Шушканов Данила",
					Age:      12,
					Salary:   12,
					Teachers: []*api.CreateTeacherRequest{
						{
							FullName:     "",
							PositionType: 11,
							StudentId:    1,
						},
					},
				},
			},
			errorField: "position_type",
			errorCode:  validation.ErrInInvalid.Code(),
		},
		{
			name: "teacher student_id not filled",
			args: args{
				&api.CreateStudentRequest{
					FullName: "Шушканов Данила",
					Age:      12,
					Salary:   12,
					Teachers: []*api.CreateTeacherRequest{
						{
							FullName:     "",
							PositionType: 11,
							StudentId:    0,
						},
					},
				},
			},
			errorField: "student_id",
			errorCode:  validation.ErrRequired.Code(),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateCreateStudentRequest(test.args.req)
			if !assert.Error(t, err, "error expected") {
				assert.FailNow(t, "error expected")
			}
			validationErr, ok := err.(validation.Errors)
			if !ok {
				assert.FailNow(t, fmt.Sprintf("expected validation errors, but not received: [%+v]", err))
			}
			fieldErr, ok := validationErr[test.errorField]
			if !ok {
				assert.FailNow(t, fmt.Sprintf("expected error in field [%s] errors, but received: [%+v]", test.errorField, err))
			}

			errObject, ok := fieldErr.(validation.ErrorObject)
			if !ok {
				assert.FailNow(t, fmt.Sprintf("expected error object, but not received: [%+v]", fieldErr))
			}
			assert.Equal(t, test.errorCode, errObject.Code(), "error code must be equal to expected")
		})
	}
}
