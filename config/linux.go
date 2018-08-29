package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
)

func reloadProfile() {
}

var re = regexp.MustCompile(`export\s+GOROOT=(.+)`)

const PROFILE = ".profile"
const BASH_PROFILE = ".bash_profile"
const BASHRC = ".bashrc"

const SOURCE_COMMAND = "source"

type linux struct {
	goroot  string
	profile string
}

var linuxProfile *linux

func init() {

	if runtime.GOOS == "linux" {
		linuxProfile = &linux{}
		profile()
		PATH = linuxProfile
	}

}

func profile() {

	var profiles = []string{PROFILE, BASH_PROFILE, BASHRC}
	homepath, _ := homedir.Dir()

	for _, pf := range profiles {

		profilePath := path.Join(homepath, pf)
		_, err := os.Stat(profilePath)
		if err == nil {
			if linuxProfile.goroot == "" {
				linuxProfile.profile = profilePath
			}

			goroot := getGoRoot(profilePath)

			if goroot != "" {
				linuxProfile.profile = profilePath
				linuxProfile.goroot = goroot
			}
		}

	}

}

func getGoRoot(profile string) (goroot string) {
	f, _ := os.Open(profile)
	defer f.Close()
	data, _ := ioutil.ReadAll(bufio.NewReader(f))

	result := re.FindAllStringSubmatch(string(data), 1)
	for _, value := range result {
		for _, s := range value {
			goroot = s
		}
		return
	}

	return
}

func setGoRoot(profile, goroot string) (result string, err error) {
	f, err := os.Open(profile)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	f.Close()
	result = re.ReplaceAllString(string(data), fmt.Sprintf(`export GOROOT=%#v`, goroot))

	ioutil.WriteFile(profile, []byte(result), 0644)

	return
}

func (l *linux) GetPath() (path string, err error) {

	path = l.goroot
	return
}

func (l *linux) ChangePath(path string) (err error) {
	linuxProfile.goroot = path

	_, err = setGoRoot(linuxProfile.profile, linuxProfile.goroot)

	if err != nil {
		return
	}

	reloadProfile()

	return
}
