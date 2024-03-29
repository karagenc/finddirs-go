package finddirs

import (
	"fmt"
	"path"
	"path/filepath"
)

type AppConfig struct {
	// Application directory. If non-empty, appended (path.Join'ed) at the end of returned paths.
	Subdir string

	// Defines whether config and state directories should be synchronizable across devices.
	//
	// If true, %USERPROFILE%\AppData\Roaming is used instead of
	// %USERPROFILE%\AppData\Local on Windows.
	// This doesn't have an effect on other systems.
	//
	// If you don't know what you're doing, just leave this as false.
	// This doesn't have an effect if `System` is set to true.
	UseRoaming bool

	// To prevent potential conflicts arising from the same directory being
	// returned by the `ConfigDir`, `StateDir`, and `CacheDir` methods, this variable can be used.
	//
	// If `StateSubdir` is non-empty, and state path is the same path as one of the other paths,
	// `StateSubdir` is appended at the end of state path to prevent overlap of files.
	//
	// `StateSubdir` doesn't have an effect if `Subdir` is empty.
	StateSubdir string

	// To prevent potential conflicts arising from the same directory being
	// returned by the `ConfigDir`, `StateDir`, and `CacheDir` methods, this variable can be used.
	//
	// If `CacheSubdir` is non-empty, and cache path is the same path as one of the other paths,
	// `CacheSubdir` is appended at the end of cache path to prevent overlap of files.
	//
	// `CacheSubdir` doesn't have an effect if `Subdir` is empty.
	CacheSubdir string
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

func (c *AppConfig) configDir(systemWide bool) (configDir string, err error) {
	if systemWide {
		configDir, err = c.configDirSystem()
	} else {
		configDir, err = c.configDirLocal()
	}
	if err != nil {
		return
	}

	if len(c.Subdir) > 0 {
		configDir = path.Join(configDir, c.Subdir)
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

	if len(c.Subdir) > 0 {
		stateDir = path.Join(stateDir, c.Subdir)

		// Append `StateSubdir` if necessary
		if c.StateSubdir != "" {
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
			if stateDir == path.Join(configDir, c.Subdir) || stateDir == path.Join(cacheDir, c.Subdir) {
				stateDir = path.Join(stateDir, c.StateSubdir)
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

	if len(c.Subdir) > 0 {
		cacheDir = path.Join(cacheDir, c.Subdir)

		// Append `CacheSubdir` if necessary
		if c.CacheSubdir != "" {
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
			if cacheDir == path.Join(stateDir, c.Subdir) || cacheDir == path.Join(configDir, c.Subdir) {
				cacheDir = path.Join(cacheDir, c.CacheSubdir)
			}
		}
	}
	return filepath.ToSlash(cacheDir), nil
}
