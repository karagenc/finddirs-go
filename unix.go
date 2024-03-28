//go:build unix && !darwin

package finddirs

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

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
	return dir, nil
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
	return dir, nil
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
	return dir, nil
}
