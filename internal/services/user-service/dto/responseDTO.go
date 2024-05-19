package userdto

type UserInternalDTO struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Msisdn   *string `json:"msisdn"`
	ImageUrl *string `json:"image_url"`
	Gender   *string `json:"gender"`
	Password *string
}

type UserDTO struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Msisdn   *string `json:"msisdn"`
	ImageUrl *string `json:"image_url"`
	Gender   *string `json:"gender"`
}

type CompanyInternalDTO struct {
	ID       string  `json:"id"`
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Msisdn   *string `json:"msisdn"`
	ImageUrl *string `json:"image_url"`
	Gender   *string `json:"gender"`
	Password *string
}
