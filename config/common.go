package config

const GOPATH_CONSTANT = `GOPATH`

type GOPATH interface {
	ChangePath(path string) error
	GetPath() (string, error)
}

var PATH GOPATH
