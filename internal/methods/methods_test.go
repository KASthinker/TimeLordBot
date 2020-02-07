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

func TestLocDate(t *testing.T) {
	t.Parallel()
	date, _ := LocDate("+08")
	require.Equal(t, date, "")
}