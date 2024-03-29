package finddirs

import (
	"fmt"
	"path/filepath"
)

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
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Desktop = filepath.ToSlash(userDirs.Desktop)

	userDirs.Downloads, err = downloadsDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Downloads = filepath.ToSlash(userDirs.Downloads)

	userDirs.Documents, err = documentsDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Documents = filepath.ToSlash(userDirs.Documents)

	userDirs.Pictures, err = picturesDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Pictures = filepath.ToSlash(userDirs.Pictures)

	userDirs.Videos, err = videosDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Videos = filepath.ToSlash(userDirs.Videos)

	userDirs.Music, err = musicDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Music = filepath.ToSlash(userDirs.Music)

	userDirs.Fonts, err = fontsDirs()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	for i, font := range userDirs.Fonts {
		userDirs.Fonts[i] = filepath.ToSlash(font)
	}

	userDirs.Templates, err = templatesDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.Templates = filepath.ToSlash(userDirs.Templates)

	userDirs.PublicShare, err = publicShareDir()
	if err != nil {
		err = fmt.Errorf("finddirs: %w", err)
		return
	}
	userDirs.PublicShare = filepath.ToSlash(userDirs.PublicShare)

	return
}
