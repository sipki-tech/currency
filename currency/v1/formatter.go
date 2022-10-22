package currencypb

import (
	"strings"
)

// Code returns currency code by ISO4217 standard.
func (x Code) Code() string {
	return strings.TrimLeft(x.String(), "CODE_")
}
