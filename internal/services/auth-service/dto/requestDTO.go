package authdto

type SignInRequestDTO struct {
	Email    string
	Msisdn   string
	Password string
}

type SignUpRequestDTO struct {
	Email    string
	Name     string
	Msisdn   string
	Gender   string
	Password string
}
