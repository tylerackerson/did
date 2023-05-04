package did

import (
	"fmt"
)

type Did struct {
	prefix    string
	separator string
	hex       string
}

// String returns the string representation of a did.
func (d Did) String() string {
	return fmt.Sprintf("%s%s%s", d.prefix, d.separator, d.hex)
}

// Length returns the integer length of a did, including the prefix, separator, and hex.
func (d Did) Length() int {
	return len(d.String())
}
