package did

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Did struct {
	prefix    string
	separator string
	hex       string
}

// New creates a randomly-generated did with the provided prefix and default separator.
// Prefix strings must be 2-3 upper or lower case alpha characters.
func New(prefix string, opts ...string) (*Did, error) {
	sep := DefaultSeparator
	if len(opts) != 0 {
		sep = opts[0]
	}

	if err := validatePrefix(prefix); err != nil {
		return nil, err
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "failed to create did")
	}

	did, err := FromUuid(u, prefix, sep)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create did")
	}

	return did, nil
}

// DidFromUuid creates a did from a UUID, the provided prefix, and default separator.
// Prefix strings must be 2-3 upper or lower case alpha characters.
func FromUuid(uuid uuid.UUID, prefix string, opts ...string) (*Did, error) {
	sep := DefaultSeparator
	if len(opts) != 0 {
		sep = opts[0]
	}

	if err := validatePrefix(prefix); err != nil {
		return nil, err
	}

	hex := strings.ReplaceAll(uuid.String(), "-", "")
	return &Did{
		prefix:    prefix,
		separator: sep,
		hex:       hex,
	}, nil
}

// DidFromString creates a did from a string.
// Basic validation is performed to ensure the did is correctly formatted.
func FromString(s string, opts ...string) (*Did, error) {
	sep := DefaultSeparator
	if len(opts) != 0 {
		sep = opts[0]
	}

	parts := strings.Split(s, sep)
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
		separator: sep,
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
