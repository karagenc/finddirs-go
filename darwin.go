//go:build darwin || ios

package finddirs

import (
	"path"
	"runtime"

	"github.com/mitchellh/go-homedir"
)

var isIOS = runtime.GOOS == "ios"

func desktopDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Desktop"), nil
}

func downloadsDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Downloads"), nil
}

func documentsDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Documents"), nil
}

func picturesDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Pictures"), nil
}

func videosDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Movies"), nil
}

func musicDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Music"), nil
}

func fontsDirs() ([]string, error) {
	if isIOS {
		return nil, ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	return []string{
		path.Join(home, "Library/Fonts"),
		"/Library/Fonts",
		"/System/Library/Fonts",
		"/Network/Library/Fonts",
	}, nil
}

func templatesDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Templates"), nil
}

func publicShareDir() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedUserDirs
	}
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Public"), nil
}

func (c *AppConfig) configDirSystem() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedAppDirsSystemIOS
	}
	return "/Library/Application Support", nil
}

func (c *AppConfig) configDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Library/Application Support"), nil
}

func (c *AppConfig) stateDirSystem() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedAppDirsSystemIOS
	}
	return "/Library/Application Support", nil
}

func (c *AppConfig) stateDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Library/Application Support"), nil
}

func (c *AppConfig) cacheDirSystem() (string, error) {
	if isIOS {
		return "", ErrOSNotSupportedAppDirsSystemIOS
	}
	return "/Library/Caches", nil
}

func (c *AppConfig) cacheDirLocal() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, "Library/Caches"), nil
}
