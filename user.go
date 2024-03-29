package finddirs

type UserDirs struct {
	Desktop     string
	Downloads   string
	Documents   string
	Pictures    string
	Videos      string
	Music       string
	Fonts       []string
	Templates   string
	PublicShare string
}

// On Linux, XDG directories may be unset. If a directory is unset,
// its value within `UserDirs` struct will be empty.
func RetrieveUserDirs() (userDirs *UserDirs, err error) {
	userDirs = new(UserDirs)

	userDirs.Desktop, err = desktopDir()
	if err != nil {
		return
	}
	userDirs.Downloads, err = desktopDir()
	if err != nil {
		return
	}
	userDirs.Documents, err = documentsDir()
	if err != nil {
		return
	}
	userDirs.Pictures, err = picturesDir()
	if err != nil {
		return
	}
	userDirs.Videos, err = videosDir()
	if err != nil {
		return
	}
	userDirs.Music, err = musicDir()
	if err != nil {
		return
	}
	userDirs.Fonts, err = fontsDirs()
	if err != nil {
		return
	}
	userDirs.Templates, err = templatesDir()
	if err != nil {
		return
	}
	userDirs.PublicShare, err = publicShareDir()
	if err != nil {
		return
	}
	return
}
