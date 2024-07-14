package userservice

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"slices"
	"strings"

	model "quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
	"quups-backend/internal/services/user-service/models"
	"quups-backend/internal/utils"
	local_jwt "quups-backend/internal/utils/jwt"
)

func ValidateCreateCompanyQParams(body userdto.CreateCompanyParams) error {
	if body.Email == "" || body.Msisdn == "" || body.Name == "" {
		return errors.New("company email, name & phone number is required")
	}

	if !utils.IsVaildEmail(body.Email) {
		return invalidEmailErr
	}

	_, isvalid := utils.IsValidCompanyName(body.Name)

	if !isvalid {
		return invalidNameErr
	}

	_, validmsisdn := utils.ParseMsisdn(body.Msisdn)
	if !validmsisdn {
		return invalidMsisdnErr
	}

	return nil
}

func (s *service) createCompanyParams(body userdto.CreateCompanyParams) (model.CreateCompanyParams, error) {
	auth_user := local_jwt.GetAuthContext(s.ctx)

	cname, _ := utils.IsValidCompanyName(body.Name)
	msisdn, _ := utils.ParseMsisdn(body.Msisdn)

	slog.Info(
		"setting up params to create a new company with",
		"name",
		body.Name,
		"email",
		body.Email,
		"msisdn",
		body.Msisdn,
		"by",
		auth_user.Sub,
	)

	p := model.CreateCompanyParams{
		ID:           utils.GenerateIntID(6),
		Email:        body.Email,
		Name:         cname,
		Msisdn:       string(msisdn),
		Slug:         strings.ToLower(strings.ReplaceAll(cname, " ", "-")),
		CurrencyCode: "GHS",
		IsActive:     false,
		OwnerID:      auth_user.Sub,
	}

	if body.BrandType != "" && slices.Contains(BRAND_TYPES, body.BrandType) {
		p.BrandType = body.BrandType
	} else {
		return p, invalidBrandTypeErr
	}

	// TODO: check invitationCode

	if c, _ := s.GetCompanyByID(p.ID); c.ID != "" {
		slog.Warn("company id already exist. used", "BY", c.Name)
		slog.Info("generating new company id")

		p.ID = utils.GenerateIntID(6)

		slog.Info("new company id generated", "id", p.ID)
	}

	if c, _ := s.GetCompanyByName(body.Name); c.ID != "" {
		slog.Warn("company name already exist ", "Warn", body.Name)

		return p, errors.New("company already in use. Please choose another one")
	}

	if utils.ParseURL(body.BannerUrl) == nil {
		// check if the string is a url
		p.BannerUrl.String = body.BannerUrl
		p.BannerUrl.Valid = true
	}

	if utils.ParseURL(body.ImageUrl) == nil {
		// check if the string is a url
		p.ImageUrl.String = body.ImageUrl
		p.ImageUrl.Valid = true
	}

	if body.About != "" {
		p.About.String = body.About
		p.About.Valid = true
	}

	return p, nil
}

func (s *service) CreateCompany(body userdto.CreateCompanyParams) (userdto.CompanyInternalDTO, error) {
	slog.Info("about to create new company")

	repo := s.db.NewRepository()
	result := userdto.CompanyInternalDTO{}

	params, err := s.createCompanyParams(body)
	if err != nil {
		log.Printf("failed to create company error: [%s]", err.Error())

		return result, err
	}

	user, err := s.FindByID()

	if err != nil {
		slog.Error("error fetching user by id", "Error", err.Error())
		return result, err
	}

	tx, err := s.db.NewRawDB().Begin()

	if err != nil {
		return result, err
	}

	defer func() {
		_ = tx.Rollback()
	}()

	qtx := repo.WithTx(tx)

	nc, err := qtx.CreateCompany(s.ctx, params)

	if err != nil {
		slog.Error("error creating company.", "Error", err)

		return result, errors.New("an error occured while creating company. please try again")

	}

	_, err = qtx.AddMember(s.ctx, model.AddMemberParams{
		CompanyID: nc.ID,
		Name:      user.Name,
		Msisdn:    user.Msisdn,
		Role:      string(models.ADMIN_ROLE),
		Status:    string(models.ACTIVE_STATUS),
		UserID: sql.NullString{
			String: user.ID,
			Valid:  true,
		},
		Email: sql.NullString{
			String: user.Email,
			Valid:  true,
		},
	})

	if err != nil {
		slog.Error("error creating company.", "Error", err)

		return result, errors.New("an error occured while creating company. please try again")
	}

	err = tx.Commit()

	if err != nil {
		slog.Error("error creating company.", "Error", err)
		return result, errors.New("an error occured while creating company. please try again")
	}

	c := mapToCompanyInternalDTO(nc)

	slog.Info("company created successully")

	return c, nil
}

func (s *service) CreatePaymentAccount() {

}

// func (s *service) createPayoutAccount(qtx model.Queries) {

// 	qtx.CreatePayoutAccount(s.ctx, model.CreatePayoutAccountParams{})

// }

// func (s *service) createPaymentAccountDetails(qtx model.Queries) {

// 	qtx.CreatePaymentAccountDetails(s.ctx, model.CreatePaymentAccountDetailsParams{})

// }

func (s *service) GetAllCompanies() ([]userdto.CompanyInternalDTO, error) {
	repo := s.db.NewRepository()
	c, err := repo.GetAllCompanies(s.ctx)

	results := []userdto.CompanyInternalDTO{}

	if err != nil {
		log.Printf("error fetching all companies  [%s]", err.Error())
		return nil, err
	}

	for _, cu := range c {
		c := mapToCompanyInternalDTO(cu)

		results = append(results, c)
	}

	return results, nil
}

func (s *service) GetCompanyByName(name string) (userdto.CompanyInternalDTO, error) {
	slog.Info("about to fetch company ", "with name", name)

	result := userdto.CompanyInternalDTO{}

	repo := s.db.NewRepository()
	res, err := repo.GetCompanyByName(s.ctx, name)

	if err != nil {
		slog.Error("fetching company", "Error", err)
		return result, fmt.Errorf("company with name: [%s] not found", name)
	}

	result = mapToCompanyInternalDTO(res)

	slog.Info("company retrieved successfully")

	return result, nil
}

func (s *service) GetCompanyByID(id string) (userdto.CompanyInternalDTO, error) {
	repo := s.db.NewRepository()
	data, err := repo.GetCompanyByID(s.ctx, id)

	result := userdto.CompanyInternalDTO{}

	if err != nil {
		slog.Error("error fetching company with id:", "Error", err)
		return result, fmt.Errorf("company with id [%s] not found", id)
	}

	result = mapToCompanyInternalDTO(data)

	return result, nil
}

func mapToCompanyInternalDTO(c model.Company) userdto.CompanyInternalDTO {
	dto := userdto.CompanyInternalDTO{
		ID:             c.ID,
		Name:           c.Name,
		Email:          c.Email,
		Msisdn:         c.Msisdn,
		About:          c.About.String,
		ImageUrl:       c.ImageUrl.String,
		BannerUrl:      c.BannerUrl.String,
		Tin:            c.Tin.String,
		BrandType:      c.BrandType,
		OwnerID:        c.OwnerID,
		CurrencyCode:   c.CurrencyCode,
		InvitationCode: c.InvitationCode.String,
		Slug:           c.Slug,
		TotalSales:     c.TotalSales,
		IsActive:       c.IsActive,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}

	if c.About.Valid {
		dto.About = c.About.String
	}

	if c.ImageUrl.Valid {
		dto.ImageUrl = c.ImageUrl.String
	}

	if c.BannerUrl.Valid {
		dto.BannerUrl = c.BannerUrl.String
	}

	if c.Tin.Valid {
		dto.Tin = c.Tin.String
	}

	return dto
}
