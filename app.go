package finddirs

import (
	"fmt"
	"path"
	"path/filepath"
)

type AppConfig struct {
	// Application directory.
	// If non-empty, it will be appended (path.Join'ed) at the end of returned paths.
	Subdir string
	// Application directory for Unix. Overrides `Subdir`.
	// If non-empty, it will be used instead of `Subdir` on Unix based systems and
	// appended (path.Join'ed) at the end of returned paths.
	SubdirUnix string
	// Application directory for macOS and iOS. Overrides `Subdir`.
	// If non-empty, it will be used instead
	// of `Subdir` on macOS & iOS and appended (path.Join'ed) at the end of returned paths.
	SubdirDarwinIOS string
	// Application directory for Windows. Overrides `Subdir`.
	// If non-empty, it will be used instead of `Subdir` on Windows and appended (path.Join'ed)
	// at the end of returned paths.
	SubdirWindows string
	// Application directory for Plan 9. Overrides `Subdir`.
	// If non-empty, it will be used instead of `Subdir` on Plan 9 and appended (path.Join'ed)
	// at the end of returned paths.
	SubdirPlan9 string

	// Defines whether config and state directories should be synchronizable across devices.
	//
	// If true, instead of %USERPROFILE%\AppData\Local, %USERPROFILE%\AppData\Roaming is used (only with `ConfigDir`).
	// This doesn't have an effect on other systems and only applies to config directory.
	//
	// If you don't know what you're doing, just leave this as false.
	// This doesn't have an effect if `System` is set to true.
	UseRoaming bool

	// To prevent potential conflicts arising from the same directory being
	// returned by the `ConfigDir`, `StateDir`, and `CacheDir` methods, this variable can be used.
	//
	// If `SubdirState` is non-empty, and state path is the same path as one of the other paths,
	// `SubdirState` is appended at the end of state path to prevent overlap of files.
	//
	// `SubdirState` doesn't have an effect if `Subdir` is empty.
	SubdirState string

	// To prevent potential conflicts arising from the same directory being
	// returned by the `ConfigDir`, `StateDir`, and `CacheDir` methods, this variable can be used.
	//
	// If `SubdirCache` is non-empty, and cache path is the same path as one of the other paths,
	// `SubdirCache` is appended at the end of cache path to prevent overlap of files.
	//
	// `SubdirCache` doesn't have an effect if `Subdir` is empty.
	SubdirCache string
}

type AppDirs struct {
	// For files for user to configure.
	ConfigDir string
	// For files that are needed for application to save its state and
	// continue where it left off. Files inside this directory typically
	// don't need user intervention.
	//
	// Example file: SQLite database.
	StateDir string
	// For files that are not needed to persist (just like with `TempDir`)
	// but should live longer than files in `TempDir`.
	// Apple documentation had explained this properly: "Generally speaking,
	// the application does not require cache data to operate properly,
	// but it can use cache data to improve performance."
	//
	// Since cache directory can be deleted, the app must be able to
	// re-create or download the files as needed.
	//
	// Example contents: compilation cache, downloaded packages that
	// are going to be installed, downloaded videos that are going to be
	// converted to audio format and deleted afterwards.
	CacheDir string
}

func RetrieveAppDirs(systemWide bool, config *AppConfig) (appDirs *AppDirs, err error) {
	if config == nil {
		config = new(AppConfig)
	}
	appDirs = new(AppDirs)
	appDirs.ConfigDir, err = config.configDir(systemWide)
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	appDirs.StateDir, err = config.stateDir(systemWide)
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	appDirs.CacheDir, err = config.cacheDir(systemWide)
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	return
}

func (c *AppConfig) subdir() string {
	subdir := c.subdirPlatformSpecific()
	if subdir != "" {
		return subdir
	}
	return c.Subdir
}

func (c *AppConfig) configDir(systemWide bool) (configDir string, err error) {
	if systemWide {
		configDir, err = c.configDirSystem()
	} else {
		configDir, err = c.configDirLocal()
	}
	if err != nil {
		return
	}

	subdir := c.subdir()
	if len(subdir) > 0 {
		configDir = path.Join(configDir, subdir)
	}
	return filepath.ToSlash(configDir), nil
}

func (c *AppConfig) stateDir(systemWide bool) (stateDir string, err error) {
	if systemWide {
		stateDir, err = c.stateDirSystem()
	} else {
		stateDir, err = c.stateDirLocal()
	}
	if err != nil {
		return
	}

	subdir := c.subdir()
	if len(subdir) > 0 {
		stateDir = path.Join(stateDir, subdir)

		// Append `SubdirState` if necessary
		if c.SubdirState != "" {
			var (
				configDir string
				cacheDir  string
			)
			if systemWide {
				configDir, err = c.configDirSystem()
				if err != nil {
					return "", err
				}
				cacheDir, err = c.cacheDirSystem()
				if err != nil {
					return "", err
				}
			} else {
				configDir, err = c.configDirLocal()
				if err != nil {
					return "", err
				}
				cacheDir, err = c.cacheDirLocal()
				if err != nil {
					return "", err
				}
			}
			if stateDir == path.Join(configDir, subdir) || stateDir == path.Join(cacheDir, subdir) {
				stateDir = path.Join(stateDir, c.SubdirState)
			}
		}
	}
	return filepath.ToSlash(stateDir), nil
}

func (c *AppConfig) cacheDir(systemWide bool) (cacheDir string, err error) {
	if systemWide {
		cacheDir, err = c.cacheDirSystem()
	} else {
		cacheDir, err = c.cacheDirLocal()
	}
	if err != nil {
		return
	}

	subdir := c.subdir()
	if len(subdir) > 0 {
		cacheDir = path.Join(cacheDir, subdir)

		// Append `SubdirCache` if necessary
		if c.SubdirCache != "" {
			var (
				configDir string
				stateDir  string
			)
			if systemWide {
				configDir, err = c.configDirSystem()
				if err != nil {
					return "", err
				}
				stateDir, err = c.stateDirSystem()
				if err != nil {
					return "", err
				}
			} else {
				configDir, err = c.configDirLocal()
				if err != nil {
					return "", err
				}
				stateDir, err = c.stateDirLocal()
				if err != nil {
					return "", err
				}
			}
			if cacheDir == path.Join(stateDir, subdir) || cacheDir == path.Join(configDir, subdir) {
				cacheDir = path.Join(cacheDir, c.SubdirCache)
			}
		}
	}
	return filepath.ToSlash(cacheDir), nil
}
