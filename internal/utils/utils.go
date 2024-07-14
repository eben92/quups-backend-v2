package utils

import (
	"errors"
	"fmt"
	"log"
	"net/mail"
	"net/url"
	"strings"

	"github.com/aidarkhanov/nanoid"
)

const (
	MSISDN_PREFIX_P233    string = "+233"
	DEFAULT_MSISDN_PREFIX string = "233"
)

func String(s string) *string {
	return &s
}

// ReplacePrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func ReplacePrefix(s, prefix, with string) string {

	if strings.HasPrefix(s, prefix) {
		return with + strings.TrimPrefix(s, prefix)
	}

	return s

}

// eg. 0550404071 will be converted to 233550404071 in (bytes)
//
//	If msisdn doesnt start with prefix, msisdn is returned unchanged.
func ConvertToLocalMsisdn(msisdn string) ([]byte, error) {

	if len(msisdn) < 9 || len(msisdn) > 13 {
		return nil, fmt.Errorf("invalid phone number")
	}

	msisdn = ReplacePrefix(msisdn, "0", DEFAULT_MSISDN_PREFIX)
	msisdn = ReplacePrefix(msisdn, MSISDN_PREFIX_P233, DEFAULT_MSISDN_PREFIX)

	if !strings.HasPrefix(msisdn, DEFAULT_MSISDN_PREFIX) {
		return []byte(msisdn), errors.New("Invalid phone number")
	}

	return []byte(msisdn), nil
}

type Msisdn string

// ParseMsisdn removes any local prefix in msisdn(+233 or 0), returns
// the updated msisdn(233xxxx...) and a boolean value
//   - true (if msisdn is not an empty string) otherwise false.
//     If msisdn doesn't start with prefix, msisdn is returned unchanged.
func ParseMsisdn(msisdn string) (Msisdn, bool) {

	m, err := ConvertToLocalMsisdn(msisdn)

	if err != nil {
		return Msisdn(m), false
	}

	return Msisdn(m), true
}

func IsVaildEmail(e string) bool {

	_, err := mail.ParseAddress(e)

	return err == nil
}

// GenerateIntID generates random string from 0-9 based on size
func GenerateIntID(size int) string {
	dv := "0123456789"

	id, err := nanoid.Generate(dv, size)

	if err != nil {
		log.Printf("error generating company id [%s]", err.Error())
	}

	return id
}

// IsValidCompanyName returns a slice of the string s and a bool,
// with all leading and trailing white space removed, as defined by Unicode.
func IsValidCompanyName(s string) (string, bool) {
	s = strings.TrimSpace(s)

	if len(s) < 3 {
		return s, false
	}

	if includeSpecialChar(s) {
		return s, false
	}

	return s, true
}

// func includeSpecialChar(s string) bool {

// 	f := func(r rune) bool {
// 		return r < 'A' || r > 'z'
// 	}

// 	return strings.IndexFunc(s, f) != -1

// }

func includeSpecialChar(s string) bool {
	for _, char := range s {
		if !isSpaceOrLetter(char) {
			return true
		}
	}
	return false
}

// isSpaceOrLetter checks if the given rune is a space or a letter.
// It returns true if the rune is a space or a letter (uppercase or lowercase), and false otherwise.
func isSpaceOrLetter(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || r == ' '
}

// ParseURL parses the given link and checks if it has a valid scheme and host.
// It returns an error if the link is invalid.
func ParseURL(link string) error {
	u, err := url.Parse(link)

	if err != nil || u.Scheme == "" || u.Host == "" {
		return err
	}

	return nil

}
