package models

type BankType string

const (
	MOBILE_MONEY BankType = "mobile_money"
	BANK         BankType = "ghipss"
	CHARGE       int      = 8
)

type ThirdPartyWallet struct {
}
