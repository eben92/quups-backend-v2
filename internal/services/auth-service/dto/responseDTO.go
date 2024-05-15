package authdto

type ResponseUserDTO struct {
	ID          string  `json:"id"`
	Email       string  `json:"email"`
	Name        *string `json:"name"`
	Msisdn      *string `json:"msisdn"`
	ImageUrl    *string `json:"image_url"`
	Gender      *string `json:"gender"`
	AccessToken *string `json:"access_token"`
}
