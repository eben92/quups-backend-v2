package userdto

type CreateUserParams struct {
	Email    string `json:"email"    yaml:"email"`
	Name     string `json:"name"     yaml:"name"`
	Msisdn   string `json:"msisdn"   yaml:"msisdn"`
	Gender   string `json:"gender"   yaml:"gender"`
	Password string `json:"password" yaml:"password"`
}

type CreateCompanyParams struct {
	Name           string
	Email          string
	Msisdn         string
	About          string
	ImageUrl       string `json:"image_url"`
	BannerUrl      string `json:"banner_url"`
	Tin            string
	BrandType      string `json:"brand_type"`
	OwnerID        string
	CurrencyCode   string `json:"currency_code"`
	InvitationCode string `json:"invitation_code"`
	Slug           string
}
