package controller

import (
	"net/http"
	"testing"

	"github.com/pradeep/golang-micro/model"
	repositories "github.com/pradeep/golang-micro/repository"
	"github.com/stretchr/testify/mock"
)

type MockAuthStore struct {
	mock.Mock
}

func (m *MockAuthStore) GetAllEmployees() ([]model.Users, error) {
	args := m.Called()
	return args.Get(0).([]model.Users), args.Error(1)
}

func (m *MockAuthStore) GetUserbyID(id string) (*model.Users, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Users), args.Error(1)
}
func TestAuthController_GetAllEmployee(t *testing.T) {
	type fields struct {
		auth repositories.AuthStore
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "",
			fields: fields{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &AuthController{
				auth: tt.fields.auth,
			}
			s.GetAllEmployee(tt.args.w, tt.args.r)
		})
	}
}
