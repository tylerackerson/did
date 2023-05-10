package did

import (
	"database/sql/driver"
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
func New(prefix string, opts ...string) (Did, error) {
	sep := DefaultSeparator
	if len(opts) != 0 {
		sep = opts[0]
	}

	if err := validatePrefix(prefix); err != nil {
		return Did{}, err
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return Did{}, errors.Wrap(err, "failed to create did")
	}

	did, err := FromUuid(u, prefix, sep)
	if err != nil {
		return Did{}, errors.Wrap(err, "failed to create did")
	}

	return did, nil
}

// FromUuid creates a did from a UUID, the provided prefix, and default separator.
// Prefix strings must be 2-3 upper or lower case alpha characters.
func FromUuid(uuid uuid.UUID, prefix string, opts ...string) (Did, error) {
	sep := DefaultSeparator
	if len(opts) != 0 {
		sep = opts[0]
	}

	if err := validatePrefix(prefix); err != nil {
		return Did{}, err
	}

	hex := strings.ReplaceAll(uuid.String(), "-", "")
	return Did{
		prefix:    prefix,
		separator: sep,
		hex:       hex,
	}, nil
}

// FromString creates a did from a string.
// Basic validation is performed to ensure the did is correctly formatted.
func FromString(s string, opts ...string) (Did, error) {
	// TODO: determine separator from the string
	sep := DefaultSeparator
	if len(opts) != 0 {
		sep = opts[0]
	}

	parts := strings.Split(s, sep)
	if len(parts) != 2 {
		return Did{}, fmt.Errorf("invalid did string '%s'", s)
	}

	if err := validatePrefix(parts[0]); err != nil {
		return Did{}, errors.Wrapf(err, "invalid did string '%s'", s)
	}

	if len(parts[1]) != DefaultHexLength {
		return Did{}, fmt.Errorf("invalid did string '%s'", s)
	}

	return Did{
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

// Scan is very similar to UUID's Scan.
// It implements sql.Scanner so dids can be read from databases transparently.
// Database types mapping to string and []byte are supported.
func (did *Did) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		// if an empty did comes from a table, we return a null did
		if src == "" {
			return nil
		}

		d, err := FromString(src)
		if err != nil {
			return errors.Wrapf(err, "failed to Scan value into did: %s", src)
		}
		*did = d

	case []byte:
		// if an empty did comes from a table, we return a nil did
		if len(src) == 0 {
			return nil
		}

		d, err := FromString(string(src))
		if err != nil {
			return errors.Wrapf(err, "failed to Scan value into did: %s", src)
		}
		*did = d

	default:
		return fmt.Errorf("unable to scan type %T into did", src)
	}

	return nil
}

// Value is very similar to UUID's Value.
// It implements sql.Valuer so that dids can be written to databases transparently.
func (did Did) Value() (driver.Value, error) {
	return did.String(), nil
}

// Must returns a did if err is nil and panics otherwise.
func Must(did Did, err error) Did {
	if err != nil {
		panic(err)
	}
	return did
}

// MustNew creates a randomly-generated did or panics.
func MustNew(prefix string, opts ...string) Did {
	return Must(New(prefix, opts...))
}

// MustFromUuid creates a did from a UUID or panics.
func MustFromUuid(uuid uuid.UUID, prefix string, opts ...string) Did {
	return Must(FromUuid(uuid, prefix, opts...))
}

// MustFromString creates a did from a string or panics.
func MustFromString(s string, opts ...string) Did {
	return Must(FromString(s, opts...))
}
