package userdto

type CreateUserParams struct {
	Email    string `json:"email"    yaml:"email"`
	Name     string `json:"name"     yaml:"name"`
	Msisdn   string `json:"msisdn"   yaml:"msisdn"`
	Gender   string `json:"gender"   yaml:"gender"`
	Password string `json:"password" yaml:"password"`
}

type CreateCompanyParams struct {
	Name           string `json:"name,omitempty"`
	Email          string `json:"email,omitempty"`
	Msisdn         string `json:"msisdn,omitempty"`
	About          string `json:"about,omitempty"`
	ImageUrl       string `json:"image_url"`
	BannerUrl      string `json:"banner_url"`
	Tin            string `json:"tin,omitempty"`
	BrandType      string `json:"brand_type"`
	OwnerID        string
	CurrencyCode   string `json:"currency_code,omitempty"`
	InvitationCode string `json:"invitation_code"`
	Slug           string
}
