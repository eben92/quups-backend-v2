package userservice

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	model "quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
	"quups-backend/internal/utils"
)

func (s *Service) TestCreate(body *userdto.CreateUserParams) (*model.CreateUserParams, error) {
	if body.Name == "" {
		return nil, fmt.Errorf("user name is required")
	}

	u := &model.CreateUserParams{
		Email: body.Email,
		Name: sql.NullString{
			String: body.Name,
		},
	}

	if body.Gender != "" {
		u.Gender.String = body.Gender
	}

	if body.Msisdn != "" {
		u.Msisdn.String = body.Msisdn
	}

	if body.Password != "" {
		// todo: hash password here
		u.Password.String = body.Password
	}

	return u, nil
}

func (s *Service) createUserParams(
	body *userdto.CreateUserParams,
) (*model.CreateUserParams, error) {
	if body.Email == "" || body.Msisdn == "" {
		return nil, fmt.Errorf("email and phone number is required")
	}

	log.Printf(
		"setting up params to create user with name, email and msisdn: [%s] [%s] [%s]",
		body.Name,
		body.Email,
		body.Msisdn,
	)

	if len(strings.TrimSpace(body.Name)) < 3 {
		return nil, fmt.Errorf("full name must be at least 5 characters.")
	}

	if !utils.IsVaildEmail(body.Email) {
		return nil, invalidEmailErr
	}

	msisdn, isValidMsisdn := utils.IsValidMsisdn(body.Msisdn)

	if !isValidMsisdn {
		return nil, invalidMsisdnErr
	}

	p := &model.CreateUserParams{
		Email: body.Email,
		Name: sql.NullString{
			String: body.Name,
			Valid:  true,
		},
		Msisdn: sql.NullString{
			String: msisdn,
			Valid:  true,
		},
	}

	u, _ := s.repo.GetUserByEmail(s.ctx, p.Email)

	if u.ID != "" {
		log.Printf("User with email  [%s] already exist", body.Email)
		return nil, fmt.Errorf("User with email [%s] already exist", body.Email)
	}

	if body.Gender != "" {
		p.Gender.String = body.Gender
		p.Gender.Valid = true
	}

	u, _ = s.repo.GetUserByMsisdn(s.ctx, sql.NullString{
		String: body.Msisdn,
		Valid:  true,
	})

	if u.ID != "" {
		log.Printf("User with msisdn [%s] already exist", body.Msisdn)
		return nil, fmt.Errorf("Phone number [%s] already in use", body.Msisdn)
	}

	if body.Password != "" {

		if len(body.Password) < 4 {
			log.Printf("user entered an invalid password < 4 msisdn: [%s]", body.Msisdn)
			return nil, fmt.Errorf("password should be atleast 6 characters")
		}

		hashpass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
		if err != nil {
			return nil, fmt.Errorf("Something went wrong. Please try again. #1")
		}

		p.Password.String = string(hashpass)
		p.Password.Valid = true
	}

	return p, nil
}

func (s *Service) Create(body *userdto.CreateUserParams) (*userdto.UserInternalDTO, error) {
	var user *userdto.UserInternalDTO
	params, err := s.createUserParams(body)
	if err != nil {
		log.Printf("failed to create user error: [%s]", err.Error())

		return nil, err
	}

	log.Printf("about to create new user wih email [%s]", params.Email)

	u, err := s.repo.CreateUser(s.ctx, *params)
	if err != nil {
		log.Printf("error fetching user with email error:[%s]", err.Error())

		return nil, err
	}

	if u.ID == "" {
		log.Printf("failed to save data in db")

		return nil, fmt.Errorf("failed to create user. Please try again later")
	}

	user = mapToUserInternalDTO(u)

	log.Printf("new user created successfully -- email: [%s]", params.Email)

	return user, nil
}

/*
this returns full user dto includinng password
NOTE: response of ths should not be sent to the frontend/client
*/
func (s *Service) FindByEmail(e string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with email [%s]", e)
	var user *userdto.UserInternalDTO

	u, err := s.repo.GetUserByEmail(s.ctx, e)
	if err != nil {
		log.Printf("error fetching user with email [%s] error: [%s]", e, err.Error())

		return nil, fmt.Errorf("no user found")
	}

	if u.ID == "" {
		log.Printf("user with email [%s] does not exist", e)

		return nil, fmt.Errorf("no user found")
	}

	user = mapToUserInternalDTO(u)

	return user, nil
}

func (s *Service) FindByID(id string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with ID [%s] ", id)
	var user *userdto.UserInternalDTO

	u, err := s.repo.GetUserByID(s.ctx, id)
	if err != nil {
		log.Printf("error fetching user with ID [%s] error: [%s]", id, err.Error())

		return nil, fmt.Errorf("no user found")
	}

	if u.ID == "" {
		log.Printf("user with id [%s] does not exist", id)

		return nil, fmt.Errorf("no user found")
	}

	user = mapToUserInternalDTO(u)

	return user, nil
}

func (s *Service) FindByMsisdn(msisdn string) (*userdto.UserInternalDTO, error) {
	log.Printf("fetching user with msisdn [%s]", msisdn)

	if msisdn == "" {
		log.Printf("no msisdn provided")

		return nil, fmt.Errorf("Phone number is required")
	}

	var user *userdto.UserInternalDTO
	msisdn, isValidMsisdn := utils.IsValidMsisdn(msisdn)

	if !isValidMsisdn {
		log.Printf("error fetching user -- invalid msisdn [%s]", msisdn)

		return nil, fmt.Errorf("Invalid phone number")
	}

	u, err := s.repo.GetUserByMsisdn(s.ctx, sql.NullString{
		String: msisdn,
		Valid:  true,
	})
	if err != nil {
		log.Printf("error fetching user with msisdn [%s] error: [%s]", msisdn, err.Error())

		return nil, fmt.Errorf("no user found")
	}

	if u.ID == "" {
		log.Printf("user with msisdn [%s] does not exist", msisdn)

		return nil, fmt.Errorf("no user found.")
	}

	user = mapToUserInternalDTO(u)
	log.Printf("user with msisdn [%s] found", msisdn)

	return user, nil
}

func (s *Service) GetUserTeams(userId string) ([]*userdto.UserTeamDTO, error) {
	log.Printf("getting user teams user: [%s]", userId)

	teams := []*userdto.UserTeamDTO{}
	t, err := s.repo.GetUserTeams(s.ctx, sql.NullString{
		String: userId,
		Valid:  true,
	})
	if err != nil {
		log.Printf("error fetching user teams err: [%s]", err.Error())

		return nil, errors.New("could not find user teams")
	}

	for _, tm := range t {
		ut := mapToUserTeamInternalDTO(tm)

		teams = append(teams, ut)
	}

	return teams, nil
}

func (s *Service) Update(id string) {
	// todo:
}

func (s *Service) Delete(id string) {
	// todo:
}

func mapToUserInternalDTO(user model.User) *userdto.UserInternalDTO {
	dto := &userdto.UserInternalDTO{
		ID:    user.ID,
		Email: user.Email,
	}

	if user.Name.Valid {
		dto.Name = &user.Name.String
	}

	if user.Msisdn.Valid {
		dto.Msisdn = &user.Msisdn.String
	}

	if user.ImageUrl.Valid {
		dto.ImageUrl = &user.ImageUrl.String
	}

	if user.Gender.Valid {
		dto.Gender = &user.Gender.String
	}

	if user.Password.Valid {
		dto.Password = &user.Password.String
	}

	return dto
}

func mapToUserTeamInternalDTO(t model.GetUserTeamsRow) *userdto.UserTeamDTO {
	tm := &userdto.UserTeamDTO{
		ID:        t.ID,
		CompanyID: t.CompanyID,
		Msisdn:    t.Msisdn,
		Status:    t.Status,
		Company: &userdto.TeamCompanyDTO{
			ID:    t.CompanyID,
			Name:  t.CompanyName,
			Email: t.CompanyEmail,
			Slug:  t.CompanySlug,
		},
	}

	if t.Email.Valid {
		tm.Email = &t.Email.String
	}

	if t.CompanyBannerUrl.Valid {
		tm.Company.BannerUrl = &t.CompanyBannerUrl.String
	}

	if t.CompanyImageUrl.Valid {
		tm.Company.ImageUrl = &t.CompanyImageUrl.String
	}

	return tm
}
