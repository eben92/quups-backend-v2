package userservice

import (
	"errors"
	"fmt"
	"log"
	model "quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
	"quups-backend/internal/utils"
	local_jwt "quups-backend/internal/utils/jwt"
	"slices"
	"strings"
)

func (s *Service) createCompanyParams(body *userdto.CreateCompanyParams) (*model.CreateCompanyParams, error) {

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

	log.Printf("setting up params to create a new company with name: [%s], email: [%s], msisdn: [%s], by: [%s]", body.Name, body.Email, body.Msisdn, auth_user.Sub)

	p := &model.CreateCompanyParams{
		ID:           utils.GenerateIntID(6),
		Email:        body.Email,
		Name:         *cname,
		Msisdn:       msisdn,
		Slug:         strings.ToLower(strings.ReplaceAll(*cname, " ", "-")),
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

		p.ID = utils.GenerateIntID(6)

		log.Printf("new company id generated [%s]", p.ID)
	}

	if c, _ := s.repo.GetCompanyByName(s.ctx, body.Name); c.ID != "" {
		log.Printf("user name already exist  [%s]", body.Name)

		return nil, errors.New("username already in use. Please choose another one")
	}

	if body.BannerUrl != "" && utils.IsVaildEmail(body.BannerUrl) {
		// check if the string is a url
		p.BannerUrl.String = body.BannerUrl
		p.BannerUrl.Valid = true
	}

	if body.ImageUrl != "" && utils.IsVaildEmail(body.ImageUrl) {
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

	nc, err := s.repo.CreateCompany(s.ctx, *params)

	if err != nil {
		log.Printf("error creating company. [%s]", err.Error())

		return nil, err
	}

	c := mapToCompanyInternalDTO(nc)

	return c, nil

}

func (s *Service) GetAllCompanies() ([]*userdto.CompanyInternalDTO, error) {

	c, err := s.repo.GetAllCompanies(s.ctx)

	if err != nil {
		log.Printf("error fetching all companies  [%s]", err.Error())
		return nil, err
	}

	var comp = []*userdto.CompanyInternalDTO{}

	for _, cu := range c {
		c := mapToCompanyInternalDTO(cu)

		comp = append(comp, c)
	}

	return comp, nil
}

func (s *Service) GetCompanyByName(name string) (*userdto.CompanyInternalDTO, error) {

	res, err := s.repo.GetCompanyByName(s.ctx, name)

	if err != nil {
		log.Printf("error fetching company with name: [%s]", err.Error())
		return nil, fmt.Errorf("company with name: [%s] not found", name)
	}

	c := mapToCompanyInternalDTO(res)

	return c, nil
}

func (s *Service) GetCompanyByID(id string) (*userdto.CompanyInternalDTO, error) {

	res, err := s.repo.GetCompanyByID(s.ctx, id)

	if err != nil {
		log.Printf("error fetching company with id [%s]", err.Error())
		return nil, fmt.Errorf("company with id [%s] not found", id)
	}

	c := mapToCompanyInternalDTO(res)

	return c, nil
}

func mapToCompanyInternalDTO(c model.Company) *userdto.CompanyInternalDTO {

	dto := &userdto.CompanyInternalDTO{
		ID:     c.ID,
		Name:   c.Name,
		Email:  c.Email,
		Msisdn: c.Msisdn,
		// About:          &c.About.String,
		// ImageUrl:       &c.ImageUrl.String,
		// BannerUrl:      &c.BannerUrl.String,
		// Tin:            &c.Tin.String,
		BrandType:    c.BrandType,
		OwnerID:      c.OwnerID,
		CurrencyCode: c.CurrencyCode,
		// InvitationCode: &c.InvitationCode.String,
		Slug:       c.Slug,
		TotalSales: c.TotalSales,
		IsActive:   c.IsActive,
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
	}

	if c.About.Valid {
		dto.About = &c.About.String
	}

	if c.ImageUrl.Valid {
		dto.ImageUrl = &c.ImageUrl.String
	}

	if c.BannerUrl.Valid {
		dto.BannerUrl = &c.BannerUrl.String
	}

	if c.Tin.Valid {
		dto.Tin = &c.Tin.String
	}

	return dto

}
