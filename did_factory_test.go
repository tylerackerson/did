package did

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewValidPrefixDefault(t *testing.T) {
	prefixes := map[string]string{
		"ab":  "two alpha chars",
		"abc": "three alpha chars",
		"AZ":  "uppercase",
		"AbC": "upper and lowercase combined",
	}

	for pr, desc := range prefixes {
		_, err := New(pr)
		require.NoError(t, err, desc)
	}
}

func TestNewInvalidPrefixDefault(t *testing.T) {
	prefixes := map[string]string{
		"abcd": "too many chars",
		"a":    "not enough chars",
		"":     "empty",
		".$":   "has special chars",
		"k9":   "has numbers",
	}

	for pr, desc := range prefixes {
		_, err := New(pr)
		require.Error(t, err, desc)
	}
}

func TestDidFromValidStringDefault(t *testing.T) {
	d, _ := New("ab")
	d, err := FromString(d.String())
	require.NoError(t, err, d)
}

func TestDidFromInvalidStringDefaultg(t *testing.T) {
	strs := map[string]string{
		"ab-526cac35b-e74429beb4f2ecca5-6c57":  "more than 1 separator",
		"a9-526cac35b7e74429beb4f2ecca56c571":  "prefix invalid",
		"ab_526cac35b7e74429beb4f2ecca56c571":  "separator invalid",
		"ab-526cac357e74429beb4f2ecca56c571":   "hex has not enough chars",
		"ab-526cac35b7e74429beb4f2ecca56c5711": "hex has too many chars",
	}

	for s, desc := range strs {
		_, err := FromString(s)
		require.Error(t, err, desc)
	}
}

func TestDidFromUuidDefault(t *testing.T) {
	u := uuid.New()
	d, err := FromUuid(u, "us")
	require.NoError(t, err)
	require.NotEmpty(t, d)
}

func TestNewDidFactoryValid(t *testing.T) {
	prefixes := []string{"_", "-", "+", ""}
	for _, p := range prefixes {
		_, err := NewDidFactory("us", p)
		require.NoError(t, err)
	}
}

func TestNewDidFactoryInvalid(t *testing.T) {
	prefixes := []string{"%", "--", "  ", "^"}
	for _, p := range prefixes {
		_, err := NewDidFactory("us", p)
		require.Error(t, err)
	}
}

func TestFactoryNewDid(t *testing.T) {
	df, _ := NewDidFactory("us", "_")
	d, err := df.NewDid()
	require.NoError(t, err)
	require.NotEmpty(t, d)
}

func TestFactoryDidFromValidString(t *testing.T) {
	df, _ := NewDidFactory("us", "_")
	d, _ := df.NewDid()
	d, err := df.DidFromString(d.String())
	require.NoError(t, err, d)
}

func TestFactoryDidFromInvalidString(t *testing.T) {
	// TODO
}

func TestFactoryDidFromUuid(t *testing.T) {
	df, _ := NewDidFactory("us", "_")
	u := uuid.New()
	d, err := df.DidFromUuid(u)
	require.NoError(t, err)
	require.NotEmpty(t, d)
}
