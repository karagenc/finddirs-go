# finddirs

This is a Go package for retrieving common directories found across all operating systems. This package is tested on Linux, macOS, Windows, and Android. However, as I have not tested on other operating systems, I cannot guarantee it will run as expected on those, but it should work.  

## Locations

### Application Directories

| Directory                      | Unix [1][2]      | Windows [3]                                | macOS                           | Plan 9        |
| ------------------------------ | ---------------- | ------------------------------------------ | ------------------------------- | ------------- |
| Config directory (system-wide) | `/etc`           | `C:\ProgramData`                           | `/Library/Preferences`          | `/lib`        |
| Config directory (local)       | `~/.config`      | `C:\<user>\AppData\<Local or Roaming>` [4] | `~/Library/Preferences`         | `~/lib`       |
| State directory (system-wide)  | `/var/lib`       | `C:\ProgramData`                           | `/Library/Application Support`  | `/lib`        |
| State directory (local)        | `~/.local/state` | `C:\<user>\AppData\Local`                  | `~/Library/Application Support` | `~/lib`       |
| Cache directory (system-wide)  | `/var/cache`     | `C:\ProgramData`                           | `/Library/Caches`               | `/lib/cache`  |
| Cache directory (local)        | `~/.cache`       | `C:\<user>\AppData\Local`                  | `~/Library/Caches`              | `~/lib/cache` |

1. On Unix based systems, XDG environment variables `$XDG_CONFIG_HOME`, `$XDG_STATE_HOME`, and `$XDG_CACHE_HOME` are first tried for paths `~/.config`, `~/.local/state`, and `~/.cache` respectively. If the particular XDG environment variable is set, it is used instead.
2. If Termux is detected on Android, system-wide directories will be under `~/../usr` (of course, as an absolute path).
3. On Windows, [KNOWNFOLDERID constants](https://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid) are used.
4. Usage of `AppData\Local` or `AppData\Roaming` depends on whether `UseRoaming` is set to true in `Config` struct.

### User Directories

| Directory   | Unix [1], macOS, and Windows (Also see [2] and [3])                                                                   |
| ----------- | --------------------------------------------------------------------------------------------------------------------- |
| Desktop     | `~/Desktop`                                                                                                           |
| Downloads   | `~/Downloads`                                                                                                         |
| Documents   | `~/Documents`                                                                                                         |
| Pictures    | `~/Pictures`                                                                                                          |
| Videos      | `~/Videos` on Linux and Windows, `~/Movies` on macOS                                                                  |
| Music       | `~/Music`                                                                                                             |
| Templates   | `~/Templates`                                                                                                         |
| Fonts       | `$XDG_DATA_HOME/fonts`, `~/.local/share/fonts`, `~/.fonts`, `/usr/share/fonts`, and `/usr/local/share/fonts` on Linux |
|             | `~/Library/Fonts`, `/Library/Fonts`, `/System/Library/Fonts`, and `/Network/Library/Fonts` on macOS                   |
|             | `C:/WINDOWS/Fonts` and `C:/Users/<USER>/AppData/Local/Microsoft/Windows/Fonts` on Windows                             |
| PublicShare | `~/Public` on Linux and macOS, `C:\Users\Public` on Windows                                                           |

1. On Unix based systems, entries in `user-dirs.dirs` are read. If `user-dirs.dirs` cannot be found, or it's malformed, `RetrieveUserDirs` returns with error. If an entry is `$HOME/` (that means, it is empty), it is set to an empty string (`""`), and no error is returned. On Unix, check for empty directories.
2. Plan 9 is not supported. `RetrieveUserDirs` on Plan 9 system will return an error.
3. If Termux is detected on Android, the Desktop, Templates, Fonts, and PublicShare directories will be empty, as they don't exist on the this platform.

## Usage

Very straightforward: `go get -u github.com/tomruk/finddirs-go`

```go
package main

import (
	"fmt"
	"github.com/tomruk/finddirs-go"
)

func main() {
	userAppDirs, _ := finddirs.RetrieveAppDirs(false, nil)
	fmt.Println(userAppDirs)
	systemAppDirs, _ := finddirs.RetrieveAppDirs(true, nil)
	fmt.Println(systemAppDirs)

	userDirs, _ := finddirs.RetrieveUserDirs()
	fmt.Println(userDirs)
}
```

## Remarks/Notes

- Since you're dealing with directories:
  - Let me recommend [github.com/mitchellh/go-homedir](https://github.com/mitchellh/go-homedir). It is used by this library and is much more reliable than `os.UserHomeDir()`.
  - For those who want to dive deep, here are platform-specific documentations:
    - [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/latest/ar01s03.html)
    - [Apple File System Programming Guide](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/FileSystemOverview/FileSystemOverview.html)
    - [Microsoft Documentation on KNOWNFOLDERID constants](https://learn.microsoft.com/en-us/windows/win32/shell/knownfolderid)

- `filepath.ToSlash` was used for every directory returned from functions in this package. This is to prevent [shellwords](https://github.com/mattn/go-shellwords) from interpreting backslash as escape character. To test this behavior, remove [this line](https://github.com/tomruk/kopyaship/blob/460b68628d589c27f7e740f1368c79a8f57a2642/backup/backup_test.go#L164), and see what happens. Use `filepath.FromSlash` to convert slashes to OS-specific path seperators.
