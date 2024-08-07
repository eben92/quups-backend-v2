package mock_test

import (
	"context"
	"quups-backend/internal/database"
	userdto "quups-backend/internal/services/user-service/dto"
	userservice "quups-backend/internal/services/user-service/service"
)

type MockCompanyMgt interface {
	CreateCompany(body userdto.CreateCompanyParams) (userdto.CompanyInternalDTO, error)
}

type MockCompanySvc struct {
	db  database.Service
	ctx context.Context
}

func NewMockCompanySvc(ctx context.Context, db database.Service) MockCompanyMgt {
	return &MockCompanySvc{db: db, ctx: ctx}
}

func GetSampleCompany() userdto.CreateCompanyParams {
	return userdto.CreateCompanyParams{
		Name:           "Test Company",
		Email:          "test@company.com",
		Msisdn:         "0200000000",
		About:          "We are a test company",
		ImageUrl:       "https://test.com/image.jpg",
		BannerUrl:      "https://test.com/banner.jpg",
		Tin:            "123456",
		BrandType:      "FOOD",
		CurrencyCode:   "GHS",
		InvitationCode: "INV123",
	}
}

func (m *MockCompanySvc) CreateCompany(body userdto.CreateCompanyParams) (userdto.CompanyInternalDTO, error) {

	usersvc := userservice.NewCompanyService(m.ctx, m.db)

	c, err := usersvc.CreateCompany(body)

	return c, err
}
