package did

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const DefaultSeparator = "-"

var prefixRegex = regexp.MustCompile("^[a-zA-Z]{2,3}$")

type Did struct {
	prefix    string
	separator string
	hex       string
}

// New creates a randomly-generated did with the provided prefix and default separator.
// Prefix strings must be two characters.
func New(prefix string) (*Did, error) {
	if match := prefixRegex.MatchString(prefix); !match {
		return nil, fmt.Errorf("invalid prefix '%s'", prefix)
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create did")
	}

	did, err := DidFromUuid(u, prefix)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create did")
	}

	return did, nil
}

// DidFromUuid creates a did from a UUID, the provided prefix, and default separator.
// Prefix strings must be two characters.
func DidFromUuid(uuid uuid.UUID, prefix string) (*Did, error) {
	if match := prefixRegex.MatchString(prefix); !match {
		return nil, fmt.Errorf("invalid prefix '%s'", prefix)
	}

	hex := strings.ReplaceAll(uuid.String(), "-", "")
	return &Did{
		prefix:    prefix,
		separator: DefaultSeparator,
		hex:       hex,
	}, nil
}

// DidFromString creates a did from a string.
// Basic validation is performed to ensure the did is correctly formatted.
func DidFromString(s string) (*Did, error) {
	parts := strings.Split(s, DefaultSeparator)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid did string '%s'", s)
	}

	return &Did{
		prefix:    parts[0],
		separator: DefaultSeparator,
		hex:       parts[1],
	}, nil
}

// String returns the string representation of a did.
func (d Did) String() string {
	return fmt.Sprintf("%s%s%s", d.prefix, d.separator, d.hex)
}

// Length returns the integer length of a did, including the prefix, separator, and hex.
func (d Did) Length() int {
	return len(d.String())
}
