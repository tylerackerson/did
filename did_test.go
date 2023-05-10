package did

import (
	"fmt"
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

func TestScan(t *testing.T) {
	// happy path, scanning valid strings
	for _, data := range prefixes {
		if data.err != nil {
			continue
		}

		d := MustNew(data.val)

		// strings
		var did Did
		err := (&did).Scan(d.String())
		require.NoError(t, err)

		// bytes
		bytes := []byte(d.String())
		err = (&did).Scan(bytes)
		require.NoError(t, err)
	}

	// bad strings, error expected
	for i, data := range prefixes {
		if data.err != nil {
			continue
		}

		d := MustNew(data.val)
		var invalidStr string
		if i%2 == 0 {
			invalidStr = d.String()[:d.Length()-2] // remove chars
		} else {
			invalidStr = d.String() + "a9d" // add extra chars
		}

		var did Did
		err := (&did).Scan(invalidStr)
		require.Error(t, err)
		require.Contains(t, err.Error(), fmt.Sprintf("failed to Scan value into did: %s", invalidStr))
	}

	// bad types, error expected
	var did Did
	err := (&did).Scan(90)
	require.Error(t, err)
	require.Contains(t, err.Error(), "unable to scan type int")

	err = (&did).Scan(Did{})
	require.Error(t, err)
	require.Contains(t, err.Error(), "unable to scan type did.Did")

	// empty did, should not fail
	empty := Did{}
	err = (&did).Scan(empty.String())
	require.NoError(t, err)
	require.Empty(t, did)

	var emptyBytes []byte
	err = (&did).Scan(emptyBytes)
	require.NoError(t, err)
	require.Empty(t, did)
}

func TestValue(t *testing.T) {
	for _, data := range prefixes {
		if data.err != nil {
			continue
		}

		d := MustNew(data.val)
		val, _ := d.Value()
		require.Equal(t, d.String(), val)
	}
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
