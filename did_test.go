package did

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewValidPrefix(t *testing.T) {
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

func TestNewInvalidPrefix(t *testing.T) {
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

func TestString(t *testing.T) {
	prefix := "ab"
	d, _ := New(prefix)
	s := d.String()
	require.True(t, strings.HasPrefix(s, prefix))
	require.Equal(t, 1, strings.Count(s, DefaultSeparator))
}

func TestDidFromValidString(t *testing.T) {
	d, _ := New("ab")
	d, err := DidFromString(d.String())
	require.NoError(t, err, d)
}

func TestDidFromInvalidString(t *testing.T) {
	prefixes := map[string]string{
		"ab-526cac35b-e74429beb4f2ecca5-6c57":  "more than 1 separator",
		"a9-526cac35b7e74429beb4f2ecca56c571":  "prefix invalid",
		"ab_526cac35b7e74429beb4f2ecca56c571":  "separator invalid",
		"ab-526cac357e74429beb4f2ecca56c571":   "hex has not enough chars",
		"ab-526cac35b7e74429beb4f2ecca56c5711": "hex has too many chars",
	}

	for pr, desc := range prefixes {
		_, err := DidFromString(pr)
		require.Error(t, err, desc)
	}
}

func TestDidFromUuid(t *testing.T) {
	u := uuid.New()
	d, err := DidFromUuid(u, "us")
	require.NoError(t, err)
	require.NotEmpty(t, d)
}
