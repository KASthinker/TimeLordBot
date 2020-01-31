package methods

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPath(t *testing.T) {
	t.Parallel()
	path := GetPath("/configs/helpconf.toml")
	require.Equal(t, path, "/media/data/Projects/GO/TimeLordBot/configs/helpconf.toml")
}

