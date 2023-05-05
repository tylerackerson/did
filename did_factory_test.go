package did

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDidFactory(t *testing.T) {
	for _, prefixData := range prefixes {
		for _, sepData := range separators {
			df, err := NewDidFactory(prefixData.val, sepData.val)

			if prefixData.err != nil || sepData.err != nil {
				require.Error(t, err, prefixData.desc, sepData.desc)
			} else {
				require.NoError(t, err, prefixData.desc, sepData.desc)

				df, err := df.NewDid()
				require.NoError(t, err)
				require.NotEmpty(t, df)
			}
		}
	}
}
