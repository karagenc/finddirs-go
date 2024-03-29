//go:build plan9

package finddirs

import (
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestPlan9AppDirsSystem(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		UseRoaming:  false,
		StateSubdir: "state",
		CacheSubdir: "cache",
	}
	d, err := RetrieveAppDirs(true, config)
	require.NoError(t, err)

	require.Equal(t, "/lib/foo/bar", d.ConfigDir)
	require.Equal(t, "/lib/foo/bar/state", d.StateDir)
	require.Equal(t, "/lib/cache/foo/bar", d.CacheDir)
}

func TestPlan9AppDirsLocal(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		UseRoaming:  false,
		StateSubdir: "state",
		CacheSubdir: "cache",
	}
	d, err := RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/lib/foo/bar", d.ConfigDir)
	require.Equal(t, home+"/lib/foo/bar/state", d.StateDir)
	require.Equal(t, home+"/lib/cache/foo/bar", d.CacheDir)
}

func TestPlan9UserDirs(t *testing.T) {
	_, err := RetrieveUserDirs()
	require.Error(t, err)
}
