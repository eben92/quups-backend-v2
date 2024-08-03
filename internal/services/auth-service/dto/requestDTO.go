package authdto

type SignInRequestDTO struct {
	Email    string `json:"email"`
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}

type AccountSigninDTO struct {
	ID string `json:"id"`
}

type SignUpRequestDTO struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Msisdn   string `json:"msisdn"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}
