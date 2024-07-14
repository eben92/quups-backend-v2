package userdto

import "time"

type UserInternalDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name,omitempty"`
	Msisdn   string `json:"msisdn,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Password string
}

type UserDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Msisdn   string `json:"msisdn"`
	ImageUrl string `json:"image_url"`
	Gender   string `json:"gender"`
}

type CompanyInternalDTO struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	About          string    `json:"about"`
	Msisdn         string    `json:"msisdn"`
	Email          string    `json:"email"`
	Tin            string    `json:"tin,omitempty"`
	ImageUrl       string    `json:"image_url"`
	BannerUrl      string    `json:"banner_url"`
	BrandType      string    `json:"brand_type"`
	OwnerID        string    `json:"owner_id"`
	TotalSales     int32     `json:"total_sales"`
	IsActive       bool      `json:"is_active"`
	CurrencyCode   string    `json:"currency_code,omitempty"`
	InvitationCode string    `json:"invitation_code,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TeamCompanyDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Msisdn    string `json:"msisdn"`
	BannerUrl string `json:"banner_url"`
	ImageUrl  string `json:"image_url"`
	IsActive  bool   `json:"is_active"`
	Slug      string `json:"slug"`
}

type UserTeamDTO struct {
	Company     TeamCompanyDTO `json:"company"`
	ID          string         `json:"id"`
	Role        string         `json:"role"`
	Status      string         `json:"status"`
	Msisdn      string         `json:"msisdn"`
	Email       string         `json:"email"`
	CompanyID   string         `json:"company_id"`
	AccessToken string         `json:"-"`
}
