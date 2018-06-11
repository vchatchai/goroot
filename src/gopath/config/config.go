package config

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const FILE_CONFIG_NAME = "gopath"

var FILE_CONFIG_PATH string

const KEY_GOPATH = "GOPATH"

type Path struct {
	Key   string
	Value string
}

func (p Path) String() string {
	return p.Value + " : " + p.Key
}
func init() {
	path, _ := homedir.Dir()

	// filepath := path.Join(path.Dir(filename), "../config/settings.toml")
	FILE_CONFIG_PATH = path + "/.gopath.json"
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

	fmt.Println("RemovePath", paths)
	delete(paths, path)
	fmt.Println("RemovePath", paths)

	viper.Set(KEY_GOPATH, paths)
	err = viper.WriteConfig()

	return
}

func writeConfig() (err error) {
	// viper.SetDefault("config", "")
	fmt.Println("tesat")
	// viper.AddConfigPath(FILE_CONFIG_PATH)
	fmt.Println(FILE_CONFIG_PATH)
	viper.SetConfigFile(FILE_CONFIG_PATH)
	// viper.SetConfigType("json")
	// viper.SetConfigName(".gopath")
	// paths := []string{
	// 	"/usr",
	// 	"/var",
	// }

	paths := map[string]string{
		"/var": "",
		"/usr": "",
	}

	viper.Set(KEY_GOPATH, paths)

	err = viper.WriteConfig()

	if err != nil {
		// log.Print(fmt.Errorf("Fatal error config file: %s \n", err))

		panic(err)
	}

	return
}
