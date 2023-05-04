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
var separatorRegex = regexp.MustCompile(`^[_+-]?$`)

type DidFactory struct {
	prefix    string
	separator string
}

func NewDidFactory(prefix string, separator string) (*DidFactory, error) {
	if err := validatePrefix(prefix); err != nil {
		return nil, err
	}
	if err := validateSeparator(separator); err != nil {
		return nil, err
	}

	return &DidFactory{prefix: prefix, separator: separator}, nil
}

func (df *DidFactory) NewDid() (*Did, error) {
	return New(df.prefix, df.separator)
}

func (df *DidFactory) DidFromUuid(uuid uuid.UUID) (*Did, error) {
	return FromUuid(uuid, df.prefix, df.separator)
}

func (df *DidFactory) DidFromString(s string) (*Did, error) {
	return FromString(s, df.separator)
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

func validateSeparator(s string) error {
	if match := separatorRegex.MatchString(s); !match {
		return fmt.Errorf("invalid prefix '%s'", s)
	}
	return nil
}
