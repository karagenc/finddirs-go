//go:build darwin

package finddirs

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func desktopDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Desktop"), nil
}

func downloadsDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Downloads"), nil
}

func documentsDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Documents"), nil
}

func picturesDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Pictures"), nil
}

func videosDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Movies"), nil
}

func musicDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Music"), nil
}

func fontsDirs() ([]string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	return []string{
		filepath.Join(home, "Library/Fonts"),
		"Library/Fonts",
		"System/Library/Fonts",
		"Network/Library/Fonts",
	}, nil
}

func templatesDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Templates"), nil
}

func publicShareDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Public"), nil
}

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
