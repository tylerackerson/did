package did

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

var u uuid.UUID
var d Did

func TestDidFromUuid(t *testing.T) {
	u = uuid.New()
	did, err := DidFromUuid(u, "us")
	require.NoError(t, err)
	require.NotEmpty(t, did)

	d = *did
}

func TestString(t *testing.T) {
	uStr := strings.ReplaceAll(u.String(), "-", "")
	expected := "us" + DefaultSeparator + uStr
	didStr := d.String()
	require.Equal(t, expected, didStr)
}
