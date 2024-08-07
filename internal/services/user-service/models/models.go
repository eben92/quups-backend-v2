package models

type Role string

const (
	OWNER_ROLE      Role = "OWNER"
	DISPATCHER_ROLE Role = "DISPATCHER"
	WAITER_ROLE     Role = "WAITER"
	AGENT_ROLE      Role = "AGENT"
)

type Status string

const (
	ACTIVE_STATUS   Status = "ACTIVE"
	INACTIVE_STATUS Status = "INACTIVE"
)

type Address struct {
	CompanyID        string  `json:"company_id"`
	UserID           string  `json:"user_id"`
	Msisdn           string  `json:"msisdn"`
	IsDefault        bool    `json:"is_default"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Description      string  `json:"description"`
	FormattedAddress string  `json:"formatted_address"`
	CountryCode      string  `json:"country_code"`
	Region           string  `json:"region"`
	Street           string  `json:"street"`
	City             string  `json:"city"`
	Country          string  `json:"country"`
	PostalCode       string  `json:"postal_code"`
}
