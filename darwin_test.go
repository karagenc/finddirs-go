//go:build darwin || ios

package finddirs

import (
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestDarwinAppDirsSystem(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		UseRoaming:  false,
		StateSubdir: "state",
		CacheSubdir: "cache",
	}
	d, err := RetrieveAppDirs(true, config)
	require.NoError(t, err)

	require.Equal(t, "/Library/Application Support/foo/bar", d.ConfigDir)
	require.Equal(t, "/Library/Application Support/foo/bar/state", d.StateDir)
	require.Equal(t, "/Library/Caches/foo/bar", d.CacheDir)
}

func TestDarwinAppDirsLocal(t *testing.T) {
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

	require.Equal(t, home+"/Library/Application Support/foo/bar", d.ConfigDir)
	require.Equal(t, home+"/Library/Application Support/foo/bar/state", d.StateDir)
	require.Equal(t, home+"/Library/Caches/foo/bar", d.CacheDir)
}

func TestDarwinUserDirs(t *testing.T) {
	d, err := RetrieveUserDirs()
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/Desktop", d.Desktop)
	require.Equal(t, home+"/Downloads", d.Downloads)
	require.Equal(t, home+"/Documents", d.Documents)
	require.Equal(t, home+"/Pictures", d.Pictures)
	require.Equal(t, home+"/Movies", d.Videos)
	require.Equal(t, home+"/Music", d.Music)
	require.Equal(t, home+"/Templates", d.Templates)
	require.Equal(t, home+"/Public", d.PublicShare)

	require.Equal(t,
		[]string{
			home + "/Library/Fonts",
			"/Library/Fonts",
			"/System/Library/Fonts",
			"/Network/Library/Fonts",
		},
		d.Fonts,
	)
}
