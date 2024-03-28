//go:build windows

package finddirs

import "golang.org/x/sys/windows"

func programData() (string, error) {
	return windows.KnownFolderPath(windows.FOLDERID_ProgramData, 0)
}

func appData(roaming bool, localLow bool) (string, error) {
	if roaming {
		return windows.KnownFolderPath(windows.FOLDERID_RoamingAppData, 0)
	} else if localLow {
		return windows.KnownFolderPath(windows.FOLDERID_LocalAppDataLow, 0)
	}
	return windows.KnownFolderPath(windows.FOLDERID_LocalAppData, 0)
}

func (c *AppConfig) configDirSystem() (string, error) { return programData() }

func (c *AppConfig) configDirLocal() (string, error) { return appData(c.UseRoaming, false) }

func (c *AppConfig) stateDirSystem() (string, error) { return c.configDirSystem() }

func (c *AppConfig) stateDirLocal() (string, error) { return c.configDirLocal() }

func (c *AppConfig) cacheDirSystem() (string, error) { return c.configDirSystem() }

func (c *AppConfig) cacheDirLocal() (string, error) { return c.configDirLocal() }
