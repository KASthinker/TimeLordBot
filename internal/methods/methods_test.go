package methods

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetPath(t *testing.T) {
	t.Parallel()
	path := GetPath("/configs/helpconf.toml")
	require.Equal(t, path, "/media/data/Projects/GO/TimeLordBot/configs/helpconf.toml")
}