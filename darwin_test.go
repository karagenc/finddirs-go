//go:build darwin

package finddirs

import (
	"path"
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
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, path.Join(home, "/Library/Application Support/foo/bar"), d.ConfigDir)
	require.Equal(t, path.Join(home, "/Library/Application Support/foo/bar/state"), d.StateDir)
	require.Equal(t, path.Join(home, "/Library/Caches/foo/bar"), d.CacheDir)
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

	require.Equal(t, path.Join(home, "Library/Application Support/foo/bar"), d.ConfigDir)
	require.Equal(t, path.Join(home, "Library/Application Support/foo/bar/state"), d.StateDir)
	require.Equal(t, path.Join(home, "Library/Caches/foo/bar"), d.CacheDir)
}

func TestDarwinUserDirs(t *testing.T) {
	d, err := RetrieveUserDirs()
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, path.Join(home, "Desktop"), d.Desktop)
	require.Equal(t, path.Join(home, "Downloads"), d.Downloads)
	require.Equal(t, path.Join(home, "Documents"), d.Documents)
	require.Equal(t, path.Join(home, "Pictures"), d.Pictures)
	require.Equal(t, path.Join(home, "Movies"), d.Videos)
	require.Equal(t, path.Join(home, "Music"), d.Music)
	require.Equal(t, path.Join(home, "Templates"), d.Templates)
	require.Equal(t, path.Join(home, "Public"), d.PublicShare)

	require.Equal(t,
		[]string{
			path.Join(home, "Library/Fonts"),
			"/Library/Fonts",
			"/System/Library/Fonts",
			"/Network/Library/Fonts",
		},
		d.Fonts,
	)
}
