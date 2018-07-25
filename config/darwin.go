package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
)

type darwin struct {
	gopath  string
	profile string
}

var darwinProfile *darwin

func init() {
	if runtime.GOOS == "darwin" {
		darwinProfile = &darwin{}
		var profiles = []string{PROFILE, BASH_PROFILE, BASHRC}
		homepath, _ := homedir.Dir()

		for _, pf := range profiles {

			profilePath := path.Join(homepath, pf)
			_, err := os.Stat(profilePath)
			if err == nil {
				if darwinProfile.gopath == "" {
					darwinProfile.profile = profilePath
				}

				gopath := getDarwinPath(profilePath)

				if gopath != "" {
					darwinProfile.profile = profilePath
					darwinProfile.gopath = gopath
				}
			}

		}
		PATH = darwinProfile
	}

}

func getDarwinPath(profile string) (gopath string) {
	f, _ := os.Open(profile)
	defer f.Close()
	data, _ := ioutil.ReadAll(bufio.NewReader(f))

	result := re.FindAllStringSubmatch(string(data), 1)
	for _, value := range result {
		for _, s := range value {
			gopath = s
		}
		return
	}

	return
}

func setDarwinPath(profile, gopath string) (result string, err error) {
	f, err := os.Open(profile)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	f.Close()
	result = re.ReplaceAllString(string(data), fmt.Sprintf(`export GOPATH=%#v`, gopath))

	ioutil.WriteFile(profile, []byte(result), 0644)

	return
}

func (l *darwin) GetPath() (path string, err error) {

	path = l.gopath
	return
}

func (l *darwin) ChangePath(path string) (err error) {
	darwinProfile.gopath = path

	_, err = setDarwinPath(darwinProfile.profile, darwinProfile.gopath)

	if err != nil {
		return
	}

	reloadProfile()

	return
}
