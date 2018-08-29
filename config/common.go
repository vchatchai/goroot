package config

const GOROOT_CONSTANT = `GOROOT`

type GOROOT interface {
	ChangePath(path string) error
	GetPath() (string, error)
}

var PATH GOROOT
