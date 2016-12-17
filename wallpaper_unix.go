// +build dragonfly freebsd linux netbsd openbsd solaris

package wallpaper

import (
	"os/exec"
)

// SetWallpaper sets the wallpaper on any system supporting 'feh'
func SetWallpaper(path string) error {
	exec.Command("feh", "--bg-fill", path)
}
