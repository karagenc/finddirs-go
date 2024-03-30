//go:build unix && !darwin && !ios

package finddirs

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var runningOnTermux = func() bool {
	_, err := exec.LookPath("termux-setup-storage")
	return err == nil
}()

func getValueFromXDG(key string) (string, error) {
	// We don't directly parse ~/.config/user-dirs.dirs — it is a bash script.
	// Instead, we source it, and echo out the particular variable.
	output, err := exec.Command("bash", "-c", "source ${XDG_CONFIG_HOME:-~/.config}/user-dirs.dirs && echo ${XDG_"+key+"_DIR}").CombinedOutput()
	if err == nil {
		return strings.Trim(string(output), "\n"), nil
	}

	// Try running xdg-user-dir as a fallback
	output, err = exec.Command("xdg-user-dir", key).CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.Trim(string(output), "\n"), nil
}

func readTermuxSymlink(subdir string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return os.Readlink(filepath.Join(home, "storage", subdir))
}

func desktopDir() (dir string, err error) {
	if runningOnTermux {
		return "", nil
	}

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
	if runningOnTermux {
		return readTermuxSymlink("downloads")
	}

	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir, err = getValueFromXDG("DOWNLOAD")
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
	if runningOnTermux {
		shared, err := readTermuxSymlink("shared")
		if err != nil {
			return "", err
		}
		return path.Join(shared, "Documents"), nil
	}

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
	if runningOnTermux {
		return readTermuxSymlink("pictures")
	}

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
	if runningOnTermux {
		return readTermuxSymlink("movies")
	}

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
	if runningOnTermux {
		return readTermuxSymlink("music")
	}

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
	if runningOnTermux {
		return nil, nil
	}

	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	homeLocalShareFonts := path.Join(home, ".local/share/fonts")
	xdgDataHome := os.Getenv("XDG_DATA_HOME")

	// Avoid duplicate paths
	if xdgDataHome != "" && filepath.Clean(xdgDataHome) != homeLocalShareFonts {
		dirs = append(dirs, path.Join(xdgDataHome, "fonts"))
	}
	dirs = append(dirs,
		homeLocalShareFonts,
		path.Join(home, ".fonts"),
		"/usr/share/fonts",
		"/usr/local/share/fonts",
	)
	return
}

func templatesDir() (dir string, err error) {
	if runningOnTermux {
		return "", nil
	}

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
	if runningOnTermux {
		return "", nil
	}

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

func (c *AppConfig) configDirSystem() (string, error) {
	if !runningOnTermux {
		return "/etc", nil
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "../usr/etc"), nil
}

func (c *AppConfig) configDirLocal() (string, error) {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return path.Join(home, ".config"), nil
	}
	return filepath.Clean(dir), nil
}

func (c *AppConfig) stateDirSystem() (string, error) {
	if !runningOnTermux {
		return "/var/lib", nil
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "../usr/var/lib"), nil
}

func (c *AppConfig) stateDirLocal() (string, error) {
	dir := os.Getenv("XDG_STATE_HOME")
	if dir == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return path.Join(home, ".local/state"), nil
	}
	return filepath.Clean(dir), nil
}

func (c *AppConfig) cacheDirSystem() (string, error) {
	if !runningOnTermux {
		return "/var/cache", nil
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "../usr/var/cache"), nil
}

func (c *AppConfig) cacheDirLocal() (string, error) {
	dir := os.Getenv("XDG_CACHE_HOME")
	if dir == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", err
		}
		return path.Join(home, ".cache"), nil
	}
	return filepath.Clean(dir), nil
}
