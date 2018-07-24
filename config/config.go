package config

import (
	"fmt"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const FILE_CONFIG_NAME = "gopath"
const QUIT = "quit"
const CONFIG_FILE = ".gopath.json"

var FILE_CONFIG_PATH string

const KEY_GOPATH = "GOPATH"

type Path struct {
	Key   string
	Value string
}

func (p Path) String() string {
	if QUIT == p.Key {
		return QUIT
	}
	return p.Value + "::" + p.Key
}
func init() {
	p, _ := homedir.Dir()
	FILE_CONFIG_PATH = path.Join(p, CONFIG_FILE)
	viper.SetConfigFile(FILE_CONFIG_PATH)
	viper.ReadInConfig()
}

func GetConfig(key string) (str map[string]string, err error) {
	str = viper.GetStringMapString(key)
	return
}

func GetPath() (path []Path, err error) {
	mapPath, err := GetConfig(KEY_GOPATH)
	if err != nil {
		return
	}
	for key, value := range mapPath {
		path = append(path, Path{Key: key, Value: value})
	}
	return
}

func AddNewPath(key, value string) (err error) {

	paths := viper.GetStringMap(KEY_GOPATH)

	paths[key] = value

	viper.Set(KEY_GOPATH, paths)
	err = viper.WriteConfig()
	return
}

func RemovePath(path string) (err error) {
	paths := viper.GetStringMap(KEY_GOPATH)

	delete(paths, path)

	viper.Set(KEY_GOPATH, paths)
	err = viper.WriteConfig()

	return
}

func writeConfig() (err error) {
	fmt.Println(FILE_CONFIG_PATH)
	viper.SetConfigFile(FILE_CONFIG_PATH)

	paths := map[string]string{
		"/var": "",
		"/usr": "",
	}

	viper.Set(KEY_GOPATH, paths)

	err = viper.WriteConfig()

	if err != nil {
		panic(err)
	}

	return
}
