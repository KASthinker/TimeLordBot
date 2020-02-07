package configs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfigs(t *testing.T) {
	t.Parallel()
	conf := Configs()
	require.NotEqual(t, conf.User, "")
	require.NotEqual(t, conf.Password, "")
	require.NotEqual(t, conf.Host, "")
	require.NotEqual(t, conf.DBname, "")
}

func TestGetToken(t *testing.T) {
	t.Parallel()
	token := GetToken()
	require.NotEqual(t, token, "")
}
