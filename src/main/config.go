package main

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const FILE_CONFIG_NAME = "gopath"

// const FILE_CONFIG_PATH = "/home/chatchai/.gopath.json"

var FILE_CONFIG_PATH string

const KEY_GOPATH = "GOPATH"

func init() {
	path, _ := homedir.Dir()
	FILE_CONFIG_PATH = path + "/.gopath.json"
	viper.SetConfigFile(FILE_CONFIG_PATH)
	viper.ReadInConfig()
	fmt.Println(FILE_CONFIG_PATH)
}

func GetConfig() (str map[string]string, err error) {
	str = viper.GetStringMapString(KEY_GOPATH)

	return
}

func AddNewPath(path string) (err error) {

	paths := viper.GetStringMap(KEY_GOPATH)

	paths[path] = ""

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
