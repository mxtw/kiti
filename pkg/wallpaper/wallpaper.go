package wallpaper

import (
	"errors"
	"os"
	"os/exec"

	"github.com/mxtw/kiti/pkg/web"
)

type sessionType int8

const (
	UNKNOWN sessionType = iota
	WAYLAND
	X11
)

func getSessionType() sessionType {
	switch os.Getenv("XDG_SESSION_TYPE") {
	case "wayland":
		return WAYLAND
	case "x11":
		return X11
	default:
		return UNKNOWN
	}
}

func setWallpaper(file string) error {

	switch getSessionType() {
	case WAYLAND:
		exec.Command("killall", "swaybg").Run()
		exec.Command("swaybg", "-m", "scale", "-i", file).Start()
		return nil
	case X11:
		exec.Command("feh", "--bg-scale", file).Start()
		return nil
	default:
		return errors.New("unsupported session type")
	}
}

func SetFromUrl(url string) error {
	file, err := web.DownloadImage(url)
	if err != nil {
		return err
	}

	setWallpaper(file)

	return nil
}
