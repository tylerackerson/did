package did

import (
	"github.com/google/uuid"
)

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

func (df *DidFactory) NewDid() (Did, error) {
	return New(df.prefix, df.separator)
}

func (df *DidFactory) DidFromUuid(uuid uuid.UUID) (Did, error) {
	return FromUuid(uuid, df.prefix, df.separator)
}
