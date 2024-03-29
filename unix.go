//go:build unix && !darwin

package finddirs

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

func getValueFromXDG(key string) (value string, err error) {
	// Try running xdg-user-dir first
	output, err := exec.Command("xdg-user-dir", key).CombinedOutput()
	if err == nil {
		return strings.Trim(string(output), "\n"), nil
	}
	// We don't directly parse ~/.config/user-dirs.dirs — it is a bash script.
	// Instead, we source it, and echo out the particular variable.
	output, err = exec.Command("bash", "-c", "source ${XDG_CONFIG_HOME:-~/.config}/user-dirs.dirs && echo ${XDG_"+key+"_DIR}").CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(output), "\n"), nil
}

func desktopDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("DESKTOP")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func downloadsDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("DOWNLOADS")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func documentsDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("DOCUMENTS")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func picturesDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("PICTURES")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func videosDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("VIDEOS")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func musicDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("MUSIC")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func fontsDirs() (dirs []string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	homeLocalShareFonts := filepath.Join(home, ".local/share/fonts")
	xdgDataHome := os.Getenv("XDG_DATA_HOME")

	// Avoid duplicate paths
	if xdgDataHome != "" && filepath.Clean(xdgDataHome) != homeLocalShareFonts {
		dirs = append(dirs, filepath.Join(xdgDataHome, "fonts"))
	}
	dirs = append(dirs,
		homeLocalShareFonts,
		filepath.Join(home, ".fonts"),
		"/usr/share/fonts",
		"/usr/local/share/fonts",
	)
	return
}

func templatesDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("TEMPLATES")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func publicShareDir() (dir string, err error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("PUBLICSHARE")
	if err != nil {
		return "", err
	}
	dir = filepath.Clean(dir)
	home = filepath.Clean(home)
	if home == dir {
		return "", nil
	}
	return
}

func (c *AppConfig) configDirSystem() (string, error) { return "/etc", nil }

func (c *AppConfig) configDirLocal() (string, error) {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".config"), nil
	}
	return filepath.Clean(dir), nil
}

func (c *AppConfig) stateDirSystem() (string, error) { return "/var/lib", nil }

func (c *AppConfig) stateDirLocal() (string, error) {
	dir := os.Getenv("XDG_STATE_HOME")
	if dir == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".local/state"), nil
	}
	return filepath.Clean(dir), nil
}

func (c *AppConfig) cacheDirSystem() (string, error) { return "/var/cache", nil }

func (c *AppConfig) cacheDirLocal() (string, error) {
	dir := os.Getenv("XDG_CACHE_HOME")
	if dir == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".cache"), nil
	}
	return filepath.Clean(dir), nil
}
