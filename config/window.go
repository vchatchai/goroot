package config

import (
	"os"
	"os/exec"
	"runtime"
)

type window struct {
	goroot string
}

var windowProfile *window

func init() {

	if runtime.GOOS == "windows" {
		windowProfile = &window{os.Getenv(GOROOT_CONSTANT)}
		PATH = windowProfile
	}

}

func (l *window) GetPath() (path string, err error) {

	path = l.goroot
	return
}

func (l *window) ChangePath(path string) (err error) {
	windowProfile.goroot = path
	cmd := exec.Command("setx", GOROOT_CONSTANT, path)
	cmd.Run()
	return
}
