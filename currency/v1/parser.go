package currencypb

import (
	"errors"

	"github.com/sipki-tech/currency"
)

// ErrUnknownCode for invalid ISO4217 format.
var ErrUnknownCode = errors.New("unknown code")

// Parse code from string.
func Parse(txt string) (Code, error) {
	for key := range Code_name {
		if key == 0 { // Ignore CODE_UNSPECIFIED.
			continue
		}

		c := Code(key)
		if txt == c.Code() {
			return c, nil
		}
	}

	return 0, ErrUnknownCode
}

// FromCode returns code from enum.
func FromCode(code currency.Code) (Code, error) {
	return Parse(code.String())
}
