package userservice

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"strings"

	"golang.org/x/crypto/bcrypt"

	model "quups-backend/internal/database/repository"
	userdto "quups-backend/internal/services/user-service/dto"
	"quups-backend/internal/utils"
	local_jwt "quups-backend/internal/utils/jwt"
)

func (s *service) TestCreate(body *userdto.CreateUserParams) (*model.CreateUserParams, error) {
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

func ValidateCreateUserQ(body userdto.CreateUserParams) error {
	if body.Email == "" || body.Msisdn == "" {
		return fmt.Errorf("email and phone number is required")
	}

	if len(strings.TrimSpace(body.Name)) < 3 {
		return fmt.Errorf("full name must be at least 5 characters")
	}

	if !utils.IsVaildEmail(body.Email) {
		return errInvalidEmail
	}

	_, isValidMsisdn := utils.ParseMsisdn(body.Msisdn)

	if !isValidMsisdn {
		return errInvalidMsisdn
	}

	if len(body.Password) < 4 {
		log.Printf("user entered an invalid password < 4 msisdn: [%s]", body.Msisdn)
		return fmt.Errorf("password should be atleast 6 characters")
	}

	return nil

}

func (s *service) prepareUserParams(body userdto.CreateUserParams) (model.CreateUserParams, error) {

	slog.Info("setting up params to create user with name,  msisdn:", body.Name, body.Msisdn)

	p := model.CreateUserParams{
		Email: body.Email,
		Name: sql.NullString{
			String: body.Name,
			Valid:  true,
		},
		Msisdn: sql.NullString{
			String: body.Msisdn,
			Valid:  true,
		},
	}

	u, _ := s.FindByEmail(p.Email)

	if u.ID != "" {
		slog.Error("User with email  already exist", "Error", body.Email)
		return p, fmt.Errorf("user with email [%s] already exist", body.Email)
	}

	if body.Gender != "" {
		p.Gender.String = body.Gender
		p.Gender.Valid = true
	}

	msidsn, ok := utils.ParseMsisdn(body.Msisdn)

	if !ok {
		slog.Error("createUserParams - ParseMsisdn", "Error", body.Msisdn)

		return p, fmt.Errorf("invalid phone number")
	}

	u, _ = s.FindByMsisdn(msidsn)

	if u.ID != "" {
		slog.Error("User with msisdn [%s] already exist", "Error", body.Msisdn)
		return p, fmt.Errorf("phone number [%s] already in use", body.Msisdn)
	}

	hashpass, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return p, fmt.Errorf("something went wrong. Please try again. #1")
	}

	p.Password.String = string(hashpass)
	p.Password.Valid = true

	return p, nil
}

func (s *service) Create(body userdto.CreateUserParams) (userdto.UserInternalDTO, error) {
	slog.Info("about to create user", body.Name, body.Msisdn)

	result := userdto.UserInternalDTO{}
	params, err := s.prepareUserParams(body)

	if err != nil {
		slog.Error("failed to create user", "Error", err)

		return result, fmt.Errorf("failed to create user. Please try again later")
	}

	repo := s.db.NewRepository()
	u, err := repo.CreateUser(s.ctx, params)
	if err != nil {
		slog.Error("error creating user", "Error", err)

		return result, fmt.Errorf("failed to create user. Please try again later")
	}

	if u.ID == "" {
		slog.Error("failed to save data in db")

		return result, fmt.Errorf("failed to create user. Please try again later")
	}

	result = mapToUserInternalDTO(u)

	slog.Info("new user created successfully", "-- name", params.Name)

	return result, nil
}

// FindByEmail retrieves a user from the database based on the provided email.
// It returns a UserInternalDTO and an error if any.
func (s *service) FindByEmail(email string) (userdto.UserInternalDTO, error) {
	slog.Info("fetching user with", " email: ", email)
	result := userdto.UserInternalDTO{}

	repo := s.db.NewRepository()
	u, err := repo.GetUserByEmail(s.ctx, email)

	if err != nil {
		slog.Error("error fetching user", "Error", err)

		return result, fmt.Errorf("no user found")
	}

	if u.ID == "" {
		slog.Error("error fetching user", "Error", err)

		return result, fmt.Errorf("no user found")
	}

	result = mapToUserInternalDTO(u)

	slog.Info("user retrieved successfully")

	return result, nil
}

// FindByID retrieves a user by their ID. It uses the user's ID from the JWT token.
// It returns the user's internal DTO (Data Transfer Object) and an error, if any.
func (s *service) FindByID() (userdto.UserInternalDTO, error) {

	result := userdto.UserInternalDTO{}
	authuser, err := local_jwt.GetAuthContext(s.ctx, local_jwt.AUTH_CTX_KEY)

	if err != nil {
		slog.Error("FindByID", "Error", err)

		return result, fmt.Errorf("no user found")

	}

	slog.Info("fetching user with", " ID: ", authuser.Sub)

	repo := s.db.NewRepository()
	u, err := repo.GetUserByID(s.ctx, authuser.Sub)
	if err != nil {
		slog.Error("FindUserByID", "Error", err)

		return result, fmt.Errorf("no user found")
	}

	if u.ID == "" {
		slog.Warn("user with id: does not exist", " ID: ", authuser.Sub)

		return result, fmt.Errorf("no user found")
	}

	result = mapToUserInternalDTO(u)

	slog.Info("user retrieved successfully")

	return result, nil
}

// FindByMsisdn fetches a user by their MSISDN (Mobile Station International Subscriber Directory Number).
// It takes an MSISDN as input and returns a UserInternalDTO and an error.
// If the user is found, the UserInternalDTO is populated with the user's information.
// If the user is not found, an error is returned.
func (s *service) FindByMsisdn(msisdn utils.Msisdn) (userdto.UserInternalDTO, error) {
	slog.Info("fetching user with  ", "msisdn:", msisdn)

	result := userdto.UserInternalDTO{}

	repo := s.db.NewRepository()

	u, err := repo.GetUserByMsisdn(s.ctx, sql.NullString{
		String: string(msisdn),
		Valid:  true,
	})

	if err != nil {
		slog.Error("error fetching user with msisdn error:", "Error", err)

		return result, fmt.Errorf("no user found")
	}

	if u.ID == "" {
		slog.Error("user does not exist:", "Error", "no user found")

		return result, fmt.Errorf("no user found")
	}

	result = mapToUserInternalDTO(u)
	slog.Info("user retrieved successfully")

	return result, nil
}

// GetUserTeams retrieves the teams associated with the user.
// It returns a slice of userdto.UserTeamDTO and an error if any.
func (s *service) GetUserTeams() ([]userdto.TeamMemberDTO, error) {

	authuser, err := local_jwt.GetAuthContext(s.ctx, local_jwt.AUTH_CTX_KEY)
	slog.Info("getting user teams", "user:", authuser.Sub)

	if err != nil {
		slog.Error("GetUserTeams", "Error", err)

		return nil, errors.New("no data found")

	}

	repo := s.db.NewRepository()

	results := []userdto.TeamMemberDTO{}
	t, err := repo.GetUserTeams(s.ctx, sql.NullString{
		String: authuser.Sub,
		Valid:  true,
	})

	if err != nil {
		slog.Error("error fetching user teams err: ", "Error", err)

		return nil, errors.New("could not find user teams")
	}

	for _, tm := range t {
		ut := mapToUserTeamsInternalDTO(tm)

		results = append(results, ut)
	}

	slog.Info("user teams retrieved successfully")

	return results, nil
}

func (s *service) GetUserTeam(companyid string) (userdto.TeamMemberDTO, error) {
	results := userdto.TeamMemberDTO{}
	authuser, err := local_jwt.GetAuthContext(s.ctx, local_jwt.AUTH_CTX_KEY)
	slog.Info("getting user company", "user", authuser.Name)

	if err != nil {
		slog.Error("GetUserTeam", "Error", err)

		return results, errors.New("internal server error")

	}

	repo := s.db.NewRepository()

	t, err := repo.GetUserTeam(s.ctx, model.GetUserTeamParams{
		CompanyID: companyid,
		UserID: sql.NullString{
			String: authuser.Sub,
			Valid:  true,
		},
	})

	if err != nil {
		slog.Error("error fetching user team err: ", "Error", err)

		return results, errors.New("could not find user company")
	}

	results = mapToUserTeamInternalDTO(t)

	slog.Info("user company retrieved successfully")

	return results, nil
}

func (s *service) CreateUserTeam(companyId string) (model.Member, error) {
	result := model.Member{}

	slog.Info("about to create user team for ", "company:", companyId)

	u, err := s.FindByID()
	if err != nil {

		slog.Error("CreateUserTeam - FindByID", "Error", err)
		return result, err

	}

	payload := model.AddMemberParams{
		CompanyID: companyId,
		UserID: sql.NullString{
			String: u.ID,
			Valid:  true,
		},
		Name: u.Name,
		Email: sql.NullString{
			String: u.Email,
			Valid:  true,
		},
		Msisdn: u.Msisdn,
		Role:   "OWNER",
		Status: "APPROVED",
	}

	repo := s.db.NewRepository()

	result, err = repo.AddMember(s.ctx, payload)

	if err != nil {

		slog.Error("CreateUserTeam - AddMember", "Error", err)
		return result, err
	}

	slog.Info("user team created successfully")

	return result, nil
}

func (s *service) Update(id string) {
	// todo:
}

func (s *service) Delete(id string) {
	// todo:
}

func mapToUserInternalDTO(user model.User) userdto.UserInternalDTO {
	dto := userdto.UserInternalDTO{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name.String,
		Msisdn:   user.Msisdn.String,
		ImageUrl: user.ImageUrl.String,
		Gender:   user.Gender.String,
		Password: user.Password.String,
	}

	return dto
}

func mapToUserTeamInternalDTO(t model.GetUserTeamRow) userdto.TeamMemberDTO {
	tm := userdto.TeamMemberDTO{
		ID:        t.ID,
		CompanyID: t.CompanyID,
		Msisdn:    t.Msisdn,
		Status:    t.Status,
		Role:      t.Role,
		Email:     t.Email.String,
		Company: userdto.TeamCompanyDTO{
			ID:           t.CompanyID,
			Name:         t.CompanyName,
			Email:        t.CompanyEmail,
			Slug:         t.CompanySlug,
			Msisdn:       t.CompanyMsisdn,
			HasOnboarded: t.CompanyHasOnboarded,
			IsActive:     t.CompanyIsActive,
			ImageUrl:     t.CompanyImageUrl.String,
			BannerUrl:    t.CompanyBannerUrl.String,
		},
	}

	return tm
}

func mapToUserTeamsInternalDTO(t model.GetUserTeamsRow) userdto.TeamMemberDTO {
	tm := userdto.TeamMemberDTO{
		ID:        t.ID,
		CompanyID: t.CompanyID,
		Msisdn:    t.Msisdn,
		Status:    t.Status,
		Role:      t.Role,
		Email:     t.Email.String,
		Company: userdto.TeamCompanyDTO{
			ID:           t.CompanyID,
			Name:         t.CompanyName,
			Email:        t.CompanyEmail,
			Slug:         t.CompanySlug,
			Msisdn:       t.CompanyMsisdn,
			HasOnboarded: t.CompanyHasOnboarded,
			IsActive:     t.CompanyIsActive,
			ImageUrl:     t.CompanyImageUrl.String,
			BannerUrl:    t.CompanyBannerUrl.String,
		},
	}

	return tm
}
