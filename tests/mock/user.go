package mock_test

import (
	"context"
	"quups-backend/internal/database"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
)

type MockUserMgt interface {
	CreateUser() (userdto.UserInternalDTO, error)
}

type MockUserSvc struct {
	db database.Service
}

func NewMockSvc(db database.Service) MockUserMgt {
	return &MockUserSvc{db: db}
}

func (m *MockUserSvc) CreateUser() (userdto.UserInternalDTO, error) {
	usersvc := userservice.NewUserService(context.Background(), m.db)

	u, err := usersvc.Create(userdto.CreateUserParams{
		Email:    "test@user.com",
		Name:     "Test User",
		Msisdn:   "0200000000",
		Gender:   "male",
		Password: "123456",
	})

	return u, err
}
