package finddirs

import "fmt"

var (
	ErrOSNotSupportedUserDirs         = fmt.Errorf("RetrieveUserDirs doesn't support this operating system")
	ErrOSNotSupportedAppDirsSystemIOS = fmt.Errorf("cannot get system-wide app directories: iOS apps are inside a sandbox, therefore iOS apps cannot have system-wide app directories")
)
