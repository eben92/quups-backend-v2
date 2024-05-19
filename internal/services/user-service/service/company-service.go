package userservice

import (
	"errors"
	"log"
	model "quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
	"quups-backend/internal/utils"
	local_jwt "quups-backend/internal/utils/jwt"
	"slices"
	"strings"
)

func (s *Service) createCompanyParams(body *userdto.CreateCompanyParams) (*model.CreateCompanyParams, error) {
	cid := utils.GenerateID(6)
	auth_user := local_jwt.GetAuthContext(s.ctx)

	if body.Email == "" || body.Msisdn == "" || body.Name == "" {
		return nil, errors.New("company email, name & phone number is required")
	}

	if !utils.IsVaildEmail(body.Email) {
		return nil, invalidEmailErr
	}

	cname, isvalid := utils.IsValidCompanyName(body.Name)

	if !isvalid {
		return nil, invalidNameErr
	}

	msisdn, validmsisdn := utils.IsValidMsisdn(body.Msisdn)
	if !validmsisdn {
		return nil, invalidMsisdnErr
	}

	log.Printf("setting up params to create a new company with name: [%s], email: [%s], msisdn: [%s], by: [%s]", body.Name, body.Email, body.Msisdn, body.OwnerID)

	p := &model.CreateCompanyParams{
		ID:           cid,
		Email:        body.Email,
		Name:         *cname,
		Msisdn:       msisdn,
		Slug:         strings.ToLower(body.Name),
		CurrencyCode: "GHS",
		IsActive:     false,
		OwnerID:      auth_user.Sub,
	}

	if body.BrandType != "" && slices.Contains(BRAND_TYPES, body.BrandType) {
		p.BrandType = body.BrandType
	} else {
		return nil, invalidBrandTypeErr
	}

	// TODO: check invitationCode

	if c, _ := s.repo.GetCompanyByID(s.ctx, p.ID); c.ID != "" {
		log.Printf("company id already exist. used by [%s] - [%s]", c.Name, c.ID)
		log.Println("generating new company id")

		p.ID = utils.GenerateID(6)

		log.Printf("new company id generated [%s]", p.ID)
	}

	if body.BannerUrl != "" {
		// check if the string is a url
		p.BannerUrl.String = body.BannerUrl
		p.BannerUrl.Valid = true
	}

	if body.ImageUrl != "" {
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

func (s *Service) CreateCompany(body *userdto.CreateCompanyParams) (*userdto.CompanyInternalDTO, error) {
	params, err := s.createCompanyParams(body)

	if err != nil {
		log.Printf("failed to create company error: [%s]", err.Error())

		return nil, err
	}

	c := mapToCompanyInternalDTO(*params)

	return c, nil

}

func mapToCompanyInternalDTO(user model.CreateCompanyParams) *userdto.CompanyInternalDTO {

	dto := &userdto.CompanyInternalDTO{
		ID:    user.ID,
		Email: user.Email,
	}

	return dto

}
