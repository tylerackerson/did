package did

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNewValidPrefix(t *testing.T) {
	prefixes := []string{"ab", "XX", "19", "%$"} // special characters ok for now
	for _, pr := range prefixes {
		_, err := New(pr)
		require.NoError(t, err)
	}
}

func TestNewInvalidPrefix(t *testing.T) {
	prefixes := []string{"abc", "a", "", "_"} // we only validate length for now
	for _, pr := range prefixes {
		_, err := New(pr)
		require.Error(t, err)
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
