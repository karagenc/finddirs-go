//go:build unix

package finddirs

import (
	"path"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestUnixAppDirsSystem(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		UseRoaming:  false,
		StateSubdir: "state",
		CacheSubdir: "cache",
	}
	d, err := RetrieveAppDirs(true, config)
	require.NoError(t, err)

	require.Equal(t, "/etc/foo/bar", d.ConfigDir)
	require.Equal(t, "/var/lib/foo/bar", d.StateDir)
	require.Equal(t, "/var/cache/foo/bar", d.CacheDir)
}

func TestUnixAppDirsLocal(t *testing.T) {
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

	require.Equal(t, path.Join(home, ".config/foo/bar"), d.ConfigDir)
	require.Equal(t, path.Join(home, ".local/state/foo/bar"), d.StateDir)
	require.Equal(t, path.Join(home, ".cache/foo/bar"), d.CacheDir)
}

func TestUnixUserDirs(t *testing.T) {
	d, err := RetrieveUserDirs()
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, path.Join(home, "Desktop"), d.Desktop)
	require.Equal(t, path.Join(home, "Downloads"), d.Downloads)
	require.Equal(t, path.Join(home, "Documents"), d.Documents)
	require.Equal(t, path.Join(home, "Pictures"), d.Pictures)
	require.Equal(t, path.Join(home, "Videos"), d.Videos)
	require.Equal(t, path.Join(home, "Music"), d.Music)
	require.Equal(t, path.Join(home, "Templates"), d.Templates)
	require.Equal(t, path.Join(home, "Public"), d.PublicShare)

	require.Equal(t,
		[]string{
			path.Join(home, ".local/share/fonts"),
			path.Join(home, ".fonts"),
			"/usr/share/fonts",
			"/usr/local/share/fonts",
		},
		d.Fonts,
	)
}
