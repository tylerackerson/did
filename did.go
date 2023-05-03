package did

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const DefaultSeparator = "-"
const DefaultHexLength = 32

var prefixRegex = regexp.MustCompile("^[a-zA-Z]{2,3}$")

type Did struct {
	prefix    string
	separator string
	hex       string
}

// New creates a randomly-generated did with the provided prefix and default separator.
// Prefix strings must be 2-3 upper or lower case alpha characters.
func New(prefix string) (*Did, error) {
	if err := validatePrefix(prefix); err != nil {
		return nil, err
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
// Prefix strings must be 2-3 upper or lower case alpha characters.
func DidFromUuid(uuid uuid.UUID, prefix string) (*Did, error) {
	if err := validatePrefix(prefix); err != nil {
		return nil, err
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

	if err := validatePrefix(parts[0]); err != nil {
		return nil, errors.Wrapf(err, "invalid did string '%s'", s)
	}

	if len(parts[1]) != DefaultHexLength {
		return nil, fmt.Errorf("invalid did string '%s'", s)
	}

	return &Did{
		prefix:    parts[0],
		separator: DefaultSeparator,
		hex:       parts[1],
	}, nil
}

func validatePrefix(p string) error {
	if match := prefixRegex.MatchString(p); !match {
		return fmt.Errorf("invalid prefix '%s'", p)
	}
	return nil
}

// String returns the string representation of a did.
func (d Did) String() string {
	return fmt.Sprintf("%s%s%s", d.prefix, d.separator, d.hex)
}

// Length returns the integer length of a did, including the prefix, separator, and hex.
func (d Did) Length() int {
	return len(d.String())
}
