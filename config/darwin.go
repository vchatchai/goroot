package config

import (
	"runtime"
)

type darwin struct {
	gopath string
}

var darwinProfile *darwin

func init() {

	if runtime.GOOS == "darwin" {
		darwinProfile = &darwin{}
		PATH = darwinProfile
	}

}

func (l *darwin) GetPath() (path string, err error) {

	path = l.gopath
	return
}

func (l *darwin) ChangePath(path string) (err error) {
	darwinProfile.gopath = path

	return
}
