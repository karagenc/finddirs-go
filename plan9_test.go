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
		SubdirState: "state",
		SubdirCache: "cache",
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
		SubdirState: "state",
		SubdirCache: "cache",
	}
	d, err := RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/lib/foo/bar", d.ConfigDir)
	require.Equal(t, home+"/lib/foo/bar/state", d.StateDir)
	require.Equal(t, home+"/lib/cache/foo/bar", d.CacheDir)
}

func TestPlan9AppDirsSubdirPlan9(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		SubdirPlan9: "zoo/zar",
		SubdirState: "state",
		SubdirCache: "cache",
	}
	d, err := RetrieveAppDirs(true, config)
	require.NoError(t, err)

	require.Equal(t, "/lib/zoo/zar", d.ConfigDir)
	require.Equal(t, "/lib/zoo/zar/state", d.StateDir)
	require.Equal(t, "/lib/cache/zoo/zar", d.CacheDir)

	d, err = RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/lib/zoo/zar", d.ConfigDir)
	require.Equal(t, home+"/lib/zoo/zar/state", d.StateDir)
	require.Equal(t, home+"/lib/cache/zoo/zar", d.CacheDir)
}

func TestPlan9UserDirs(t *testing.T) {
	_, err := RetrieveUserDirs()
	require.Error(t, err)
}
