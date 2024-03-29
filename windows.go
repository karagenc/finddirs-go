//go:build windows

package finddirs

import (
	"path"
	"path/filepath"

	"golang.org/x/sys/windows"
)

func knownFolderPath(id *windows.KNOWNFOLDERID) (path string, err error) {
	flags := []uint32{windows.KF_FLAG_DEFAULT, windows.KF_FLAG_DEFAULT_PATH}
	for _, flag := range flags {
		path, err = windows.KnownFolderPath(id, flag|windows.KF_FLAG_DONT_VERIFY)
		if err == nil {
			path = filepath.Clean(path)
			return
		}
	}
	return
}

func desktopDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Desktop)
}

func downloadsDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Downloads)
}

func documentsDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Documents)
}

func picturesDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Pictures)
}

func videosDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Videos)
}

func musicDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Music)
}

func fontsDirs() (dirs []string, err error) {
	dir, err := knownFolderPath(windows.FOLDERID_Fonts)
	if err != nil {
		return nil, err
	}
	localAppData, err := knownFolderPath(windows.FOLDERID_LocalAppData)
	if err != nil {
		return nil, err
	}
	return []string{
		dir,
		path.Join(localAppData, "Microsoft/Windows/Fonts"),
	}, nil
}

func templatesDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Templates)
}

func publicShareDir() (string, error) {
	return knownFolderPath(windows.FOLDERID_Public)
}

func (c *AppConfig) configDirSystem() (string, error) { return programData() }

func (c *AppConfig) configDirLocal() (string, error) { return appData(c.UseRoaming) }

func (c *AppConfig) stateDirSystem() (string, error) { return programData() }

func (c *AppConfig) stateDirLocal() (string, error) { return appData(false) }

func (c *AppConfig) cacheDirSystem() (string, error) { return programData() }

func (c *AppConfig) cacheDirLocal() (string, error) { return appData(false) }

func programData() (string, error) {
	return knownFolderPath(windows.FOLDERID_ProgramData)
}

func appData(roaming bool) (string, error) {
	if roaming {
		return knownFolderPath(windows.FOLDERID_RoamingAppData)
	}
	return knownFolderPath(windows.FOLDERID_LocalAppData)
}
