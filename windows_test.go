//go:build windows

package finddirs

import (
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/require"
)

func TestWindowsAppDirsSystem(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		UseRoaming:  false,
		StateSubdir: "state",
		CacheSubdir: "cache",
	}
	d, err := RetrieveAppDirs(true, config)
	require.NoError(t, err)

	require.Equal(t, "C:/ProgramData/foo/bar", d.ConfigDir)
	require.Equal(t, "C:/ProgramData/foo/bar/state", d.StateDir)
	require.Equal(t, "C:/ProgramData/foo/bar/cache", d.CacheDir)
}

func TestWindowsAppDirsLocal(t *testing.T) {
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
	home = filepath.ToSlash(home)

	require.Equal(t, home+"/AppData/Local/foo/bar", d.ConfigDir)
	require.Equal(t, home+"/AppData/Local/foo/bar/state", d.StateDir)
	require.Equal(t, home+"/AppData/Local/foo/bar/cache", d.CacheDir)
}

func TestWindowsAppDirsLocalRoaming(t *testing.T) {
	config := &AppConfig{
		Subdir:      "foo/bar",
		UseRoaming:  true,
		StateSubdir: "state",
		CacheSubdir: "cache",
	}
	d, err := RetrieveAppDirs(false, config)
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)
	home = filepath.ToSlash(home)

	require.Equal(t, home+"/AppData/Roaming/foo/bar", d.ConfigDir)
	require.Equal(t, home+"/AppData/Local/foo/bar/state", d.StateDir)
	require.Equal(t, home+"/AppData/Local/foo/bar/cache", d.CacheDir)
}

func TestWindowsUserDirs(t *testing.T) {
	d, err := RetrieveUserDirs()
	require.NoError(t, err)
	home, err := homedir.Dir()
	require.NoError(t, err)
	home = filepath.ToSlash(home)

	require.Equal(t, home+"/Desktop", d.Desktop)
	require.Equal(t, home+"/Downloads", d.Downloads)
	require.Equal(t, home+"/Documents", d.Documents)
	require.Equal(t, home+"/Pictures", d.Pictures)
	require.Equal(t, home+"/Videos", d.Videos)
	require.Equal(t, home+"/Music", d.Music)
	require.Equal(t, home+"/AppData/Roaming/Microsoft/Windows/Templates", d.Templates)
	require.Equal(t, "C:/Users/Public", d.PublicShare)

	fontDirs := [][]string{
		{
			"C:/Windows/Fonts",
			"C:/WINDOWS/Fonts",
		},
		{home + "/AppData/Local/Microsoft/Windows/Fonts"},
	}

	for i, fontDir := range d.Fonts {
		found := false
		tests := fontDirs[i]
		for _, test := range tests {
			if test == fontDir {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("expected one of: %s found: %s", tests, fontDir)
		}
	}
}
