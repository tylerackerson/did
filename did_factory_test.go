package did

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var validPrefixes = map[string]string{
	"ab":  "two alpha chars",
	"abc": "three alpha chars",
	"AZ":  "uppercase",
	"AbC": "upper and lowercase combined",
}

var invalidPrefixes = map[string]string{
	"abcd": "too many chars",
	"a":    "not enough chars",
	"":     "empty",
	".$":   "has special chars",
	"k9":   "has numbers",
}

var invalidStrs = map[string]string{
	"ab-526cac35b-e74429beb4f2ecca5-6c57":  "more than 1 separator",
	"a9-526cac35b7e74429beb4f2ecca56c571":  "prefix invalid",
	"ab=526cac35b7e74429beb4f2ecca56c571":  "separator invalid",
	"ab-526cac357e74429beb4f2ecca56c571":   "hex has not enough chars",
	"ab-526cac35b7e74429beb4f2ecca56c5711": "hex has too many chars",
}

var validSeparators = []string{"_", "-", "+", ""}
var invalidSeparators = []string{"%", "--", "  ", "^"}

func TestNewValidPrefixDefault(t *testing.T) {
	for pr, desc := range validPrefixes {
		_, err := New(pr)
		require.NoError(t, err, desc)
	}
}

func TestNewInvalidPrefixDefault(t *testing.T) {
	for pr, desc := range invalidPrefixes {
		_, err := New(pr)
		require.Error(t, err, desc)
	}
}

func TestDidFromValidStringDefault(t *testing.T) {
	for pr, desc := range validPrefixes {
		d, _ := New(pr)
		d, err := FromString(d.String())
		require.NoError(t, err, desc)
	}
}

func TestDidFromInvalidStringDefaultg(t *testing.T) {
	for s, desc := range invalidStrs {
		_, err := FromString(s)
		require.Error(t, err, desc)
	}
}

func TestDidFromUuidDefault(t *testing.T) {
	for pr, desc := range validPrefixes {
		u := uuid.New()
		d, err := FromUuid(u, pr)
		require.NoError(t, err, desc)
		require.NotEmpty(t, d)
	}
}

func TestNewDidFactoryValid(t *testing.T) {
	for _, s := range validSeparators {
		for pr := range validPrefixes {
			_, err := NewDidFactory(pr, s)
			require.NoError(t, err)
		}
	}
}

func TestNewDidFactoryInvalidSeparator(t *testing.T) {
	for _, s := range invalidSeparators {
		for pr := range validPrefixes {
			_, err := NewDidFactory(pr, s)
			require.Error(t, err)
		}
		for pr := range invalidPrefixes {
			_, err := NewDidFactory(pr, s)
			require.Error(t, err)
		}
	}
}

func TestNewDidFactoryInvalidPrefix(t *testing.T) {
	for _, s := range invalidSeparators {
		for pr := range validPrefixes {
			_, err := NewDidFactory(pr, s)
			require.Error(t, err)
		}
		for pr := range invalidPrefixes {
			_, err := NewDidFactory(pr, s)
			require.Error(t, err)
		}
	}
}

func TestFactoryNewDid(t *testing.T) {
	df, _ := NewDidFactory("us", "_")
	d, err := df.NewDid()
	require.NoError(t, err)
	require.NotEmpty(t, d)
}

func TestFactoryDidFromValidString(t *testing.T) {
	// for now, DidFactory does not validate  for prefix or separator matches
	df, _ := NewDidFactory("us", "_")
	d, _ := df.NewDid()
	d, err := df.DidFromString(d.String())
	require.NoError(t, err, d)
}

func TestFactoryDidFromInvalidString(t *testing.T) {
	df, _ := NewDidFactory("us", "_")
	for str, desc := range invalidStrs {
		_, err := df.DidFromString(str)
		require.Error(t, err, desc)
	}
}

func TestFactoryDidFromUuid(t *testing.T) {
	df, _ := NewDidFactory("us", "_")
	u := uuid.New()
	d, err := df.DidFromUuid(u)
	require.NoError(t, err)
	require.NotEmpty(t, d)
}
