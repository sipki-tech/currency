package currencypb

import (
	"strings"

	"github.com/sipki-tech/currency"
)

// Code returns currency code by ISO4217 standard.
func (x Code) Code() currency.Code {
	c, _ := currency.Parse(strings.TrimLeft(x.String(), "CODE_"))
	return c
}

// Display returns currency code by string ISO4217 standard.
func (x Code) Display() string {
	return strings.TrimLeft(x.String(), "CODE_")
}
