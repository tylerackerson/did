package did

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	for _, data := range prefixes {
		_, err := New(data.val)

		if data.err != nil {
			require.Error(t, err, data.desc)
		} else {
			require.NoError(t, err, data.desc)
		}
	}
}

func TestDidFromString(t *testing.T) {
	for _, data := range prefixes {
		if data.err != nil {
			continue
		}
		d, _ := New(data.val)
		ds, err := FromString(d.String())
		require.NoError(t, err, data.desc)
		require.Equal(t, d, ds)
	}

	for str, desc := range invalidStrs {
		_, err := New(str)
		require.Error(t, err, desc)
	}
}

func TestDidFromUuid(t *testing.T) {
	for _, data := range prefixes {
		u := uuid.New()
		_, err := FromUuid(u, data.val)

		if data.err != nil {
			require.Error(t, err, data.desc)
		} else {
			require.NoError(t, err, data.desc)
		}
	}
}

func TestString(t *testing.T) {
	p := "ab"
	s := "+"
	h := "aaabbbcccddd"
	d := Did{prefix: p, separator: s, hex: h}
	str := d.String()
	require.True(t, strings.HasPrefix(str, p))
	require.Equal(t, 1, strings.Count(str, s))
	require.True(t, strings.HasSuffix(str, h))
}

func TestLength(t *testing.T) {
	d := Did{prefix: "ab", separator: "-", hex: "aaabbbcccddd"}
	expected := len(d.prefix) + len(d.separator) + len(d.hex)
	require.Equal(t, expected, d.Length())
}

func TestMust(t *testing.T) {
	for _, data := range prefixes {
		if data.err != nil {
			require.Panics(t, func() {
				Must(New(data.val))
			})
		} else {
			require.NotPanics(t, func() {
				Must(New(data.val))
			})
		}
	}
}

func TestMustNew(t *testing.T) {
	for _, data := range prefixes {
		if data.err != nil {
			require.Panics(t, func() {
				MustNew(data.val)
			})
		} else {
			require.NotPanics(t, func() {
				MustNew(data.val)
			})
		}
	}
}

func TestMustFromUuid(t *testing.T) {
	for _, data := range prefixes {
		u := uuid.New()

		if data.err != nil {
			require.Panics(t, func() {
				MustFromUuid(u, data.val)
			})
		} else {
			require.NotPanics(t, func() {
				MustFromUuid(u, data.val)
			})
		}
	}
}

func TestMustFromString(t *testing.T) {
	for _, data := range prefixes {
		if data.err != nil {
			continue
		}

		d := MustNew(data.val)
		require.NotPanics(t, func() {
			MustFromString(d.String())
		})
		require.Panics(t, func() {
			MustFromString(d.String() + "random#string$invalid")
		})
	}
}
