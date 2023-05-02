package did

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const DefaultSeparator = "-"

type Did struct {
	prefix    string
	separator string
	hex       string
}

func DidFromUuid(uuid uuid.UUID, prefix string) (*Did, error) {
	// TODO: validate
	hex := strings.ReplaceAll(uuid.String(), "-", "")
	return &Did{
		prefix:    prefix,
		separator: DefaultSeparator,
		hex:       hex,
	}, nil
}

func (d Did) String() string {
	return fmt.Sprintf("%s%s%s", d.prefix, d.separator, d.hex)
}
