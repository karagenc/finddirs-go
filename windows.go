//go:build windows

package finddirs

import "golang.org/x/sys/windows"

func knownFolderPath(id *windows.KNOWNFOLDERID) (path string, err error) {
	flags := []uint32{windows.KF_FLAG_DEFAULT, windows.KF_FLAG_DEFAULT_PATH}
	for _, flag := range flags {
		path, err = windows.KnownFolderPath(id, flag|windows.KF_FLAG_DONT_VERIFY)
		if err == nil {
			return
		}
	}
	return
}

func (c *AppConfig) configDirSystem() (string, error) {
	return knownFolderPath(windows.FOLDERID_ProgramData)
}

func (c *AppConfig) configDirLocal() (string, error) {
	if c.UseRoaming {
		return knownFolderPath(windows.FOLDERID_RoamingAppData)
	}
	return knownFolderPath(windows.FOLDERID_LocalAppData)
}

func (c *AppConfig) stateDirSystem() (string, error) { return c.configDirSystem() }

func (c *AppConfig) stateDirLocal() (string, error) { return c.configDirLocal() }

func (c *AppConfig) cacheDirSystem() (string, error) { return c.configDirSystem() }

func (c *AppConfig) cacheDirLocal() (string, error) { return c.configDirLocal() }
