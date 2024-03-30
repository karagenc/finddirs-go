//go:build plan9

package finddirs

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

func desktopDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func downloadsDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func documentsDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func picturesDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func videosDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func musicDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func fontsDirs() ([]string, error) {
	return nil, ErrOSNotSupportedUserDirs
}

func templatesDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func publicShareDir() (string, error) {
	return "", ErrOSNotSupportedUserDirs
}

func (c *AppConfig) subdirPlatformSpecific() string { return c.SubdirPlan9 }

func (c *AppConfig) configDirSystem() (string, error) { return "/lib", nil }

func (c *AppConfig) configDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "lib"), nil
}

func (c *AppConfig) stateDirSystem() (string, error) { return "/lib", nil }

func (c *AppConfig) stateDirLocal() (string, error) {
	os.UserCacheDir()
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "lib"), nil
}

func (c *AppConfig) cacheDirSystem() (string, error) { return "/lib/cache", nil }

func (c *AppConfig) cacheDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "lib/cache"), nil
}
