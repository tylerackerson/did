package did

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	p := "ab"
	s := "+"
	h := "aaabbbcccddd"
	d := Did{prefix: p, separator: s, hex: h}
	str := d.String()
	require.True(t, strings.HasPrefix(str, p))
	require.Equal(t, 1, strings.Count(s, s))
	require.True(t, strings.HasSuffix(str, h))
}

func TestLength(t *testing.T) {
	d := Did{prefix: "ab", separator: "-", hex: "aaabbbcccddd"}
	expected := len(d.prefix) + len(d.separator) + len(d.hex)
	require.Equal(t, expected, d.Length())
}
