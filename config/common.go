package config

type GOPATH interface {
	ChangePath(path string) error
	GetPath() (string, error)
}

var PATH GOPATH
