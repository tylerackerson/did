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

func TestDidFromString(t *testing.T) {
	d, _ := New("ab")
	_, err := DidFromString(d.String())
	require.NoError(t, err)
}

func TestDidFromUuid(t *testing.T) {
	u := uuid.New()
	d, err := DidFromUuid(u, "us")
	require.NoError(t, err)
	require.NotEmpty(t, d)
}
