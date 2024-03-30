//go:build unix && !darwin && !ios

package finddirs

import (
	"os/exec"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestUnixAppDirsSystem(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		SubdirState: "state",
		SubdirCache: "cache",
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
		SubdirState: "state",
		SubdirCache: "cache",
	}
	d, err := RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/.config/foo/bar", d.ConfigDir)
	require.Equal(t, home+"/.local/state/foo/bar", d.StateDir)
	require.Equal(t, home+"/.cache/foo/bar", d.CacheDir)
}

func TestUnixAppDirsSubdirUnix(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		SubdirUnix:  "zoo/zar",
		SubdirState: "state",
		SubdirCache: "cache",
	}
	d, err := RetrieveAppDirs(true, config)
	require.NoError(t, err)

	require.Equal(t, "/etc/zoo/zar", d.ConfigDir)
	require.Equal(t, "/var/lib/zoo/zar", d.StateDir)
	require.Equal(t, "/var/cache/zoo/zar", d.CacheDir)

	d, err = RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/.config/zoo/zar", d.ConfigDir)
	require.Equal(t, home+"/.local/state/zoo/zar", d.StateDir)
	require.Equal(t, home+"/.cache/zoo/zar", d.CacheDir)
}

func TestUnixUserDirs(t *testing.T) {
	_, err := exec.Command("xdg-user-dirs-update").CombinedOutput()
	require.NoError(t, err)

	d, err := RetrieveUserDirs()
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)

	require.Equal(t, home+"/Desktop", d.Desktop)
	require.Equal(t, home+"/Downloads", d.Downloads)
	require.Equal(t, home+"/Documents", d.Documents)
	require.Equal(t, home+"/Pictures", d.Pictures)
	require.Equal(t, home+"/Videos", d.Videos)
	require.Equal(t, home+"/Music", d.Music)
	require.Equal(t, home+"/Templates", d.Templates)
	require.Equal(t, home+"/Public", d.PublicShare)

	require.Equal(t,
		[]string{
			home + "/.local/share/fonts",
			home + "/.fonts",
			"/usr/share/fonts",
			"/usr/local/share/fonts",
		},
		d.Fonts,
	)
}
