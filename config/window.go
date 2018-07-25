package config

import (
	"os"
	"os/exec"
	"runtime"
)

type window struct {
	gopath string
}

var windowProfile *window

func init() {

	if runtime.GOOS == "windows" {
		windowProfile = &window{os.Getenv(GOPATH_CONSTANT)}
		PATH = windowProfile
	}

}

func (l *window) GetPath() (path string, err error) {

	path = l.gopath
	return
}

func (l *window) ChangePath(path string) (err error) {
	windowProfile.gopath = path
	cmd := exec.Command("setx", GOPATH_CONSTANT, path)
	cmd.Run()
	return
}
