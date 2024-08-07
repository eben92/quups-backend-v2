package mock_test

import (
	"context"
	"quups-backend/internal/database"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
)

type MockUserMgt interface {
	CreateUser(data userdto.CreateUserParams) (userdto.UserInternalDTO, error)
}

type MockUserSvc struct {
	db  database.Service
	ctx context.Context
}

func NewMockSvc(ctx context.Context, db database.Service) MockUserMgt {
	return &MockUserSvc{ctx: ctx, db: db}
}

func GetSampleUser() userdto.CreateUserParams {
	return userdto.CreateUserParams{
		Email:    "test@user.com",
		Name:     "Test User",
		Msisdn:   "0200000000",
		Gender:   "male",
		Password: "123456",
	}
}

func (m *MockUserSvc) CreateUser(data userdto.CreateUserParams) (userdto.UserInternalDTO, error) {
	usersvc := userservice.NewUserService(m.ctx, m.db)

	u, err := usersvc.Create(data)

	return u, err
}
