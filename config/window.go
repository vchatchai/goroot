package config

import (
	"runtime"
)

type window struct {
	gopath string
}

var windowProfile *window

func init() {

	if runtime.GOOS == "windows" {
		windowProfile = &window{}
		PATH = windowProfile
	}

}

func (l *window) GetPath() (path string, err error) {

	path = l.gopath
	return
}

func (l *window) ChangePath(path string) (err error) {
	windowProfile.gopath = path

	return
}
