package mock_test

import (
	"quups-backend/internal/database"
	userdto "quups-backend/internal/services/user-service/dto"
)

type MockCompanyMgt interface {
	CreateCompany(body userdto.CreateCompanyParams) (userdto.CompanyInternalDTO, error)
}

type MockCompanySvc struct {
	db database.Service
}

func NewMockCompanySvc(db database.Service) MockCompanyMgt {
	return &MockCompanySvc{db: db}
}

func (m *MockCompanySvc) CreateCompany(body userdto.CreateCompanyParams) (userdto.CompanyInternalDTO, error) {

	return userdto.CompanyInternalDTO{}, nil
}
