//go:build darwin

package finddirs

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func (c *AppConfig) configDirSystem() (string, error) { return "/Library/Application Support", nil }

func (c *AppConfig) configDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library/Application Support"), nil
}

func (c *AppConfig) stateDirSystem() (string, error) { return "/Library/Application Support", nil }

func (c *AppConfig) stateDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library/Application Support"), nil
}

func (c *AppConfig) cacheDirSystem() (string, error) { return "/Library/Caches", nil }

func (c *AppConfig) cacheDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library/Caches"), nil
}
