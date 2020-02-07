package database

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Parallel()
	db, err = Connect()
	require.Equal(t, err, nil)
	db.Close()
}

func TestIfUserExists(t *testing.T) {
	t.Parallel()
	var test int64 = 1111
	require.Equal(t, IfUserExists(test), false)
	db.Close()
}

func TestGetUsersID(t *testing.T) {
	t.Parallel()
	users, _ := GetUsers()
	require.Equal(t, users, 0)
}