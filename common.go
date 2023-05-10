package did

import (
	"fmt"
	"regexp"
)

const DefaultSeparator = "-"
const DefaultHexLength = 32

var prefixRegex = regexp.MustCompile("^[a-zA-Z]{2,3}$")
var separatorRegex = regexp.MustCompile(`[_+-]`)

func validatePrefix(p string) error {
	if match := prefixRegex.MatchString(p); !match {
		return fmt.Errorf("invalid prefix '%s'", p)
	}
	return nil
}

func validateSeparator(s string) error {
	if match := separatorRegex.MatchString(s); !match {
		return fmt.Errorf("invalid prefix '%s'", s)
	}
	return nil
}
