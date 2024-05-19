package userdto

type CreateUserParams struct {
	Email    string
	Name     string
	Msisdn   string
	Gender   string
	Password string
}

type CreateCompanyParams struct {
	Name           string
	Email          string
	Msisdn         string
	About          string
	ImageUrl       string
	BannerUrl      string
	Tin            string
	BrandType      string
	OwnerID        string
	CurrencyCode   string
	InvitationCode string
	Slug           string
}
