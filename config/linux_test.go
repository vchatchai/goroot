package config

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

func TestGetPath(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			path, err := PATH.GetPath()

			t.Error(path)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func Test_linux_ChangePath(t *testing.T) {
	type fields struct {
		goroot  string
		profile string
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Test Change Path", fields{goroot: `$HOME/go_workspace`, profile: `/home/ee56054/.profile`}, args{`/home/ee56054/go101`}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &linux{
				goroot:  tt.fields.goroot,
				profile: tt.fields.profile,
			}
			if err := l.ChangePath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("linux.ChangePath() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Fatal("done.")
		})
	}
}

func TestT(t *testing.T) {
	// cmd := exec.Command("ls -l ") ///home/ee56054/go_workspace/export.sh")
	// cmd.Run()
	// cmd.
	// fmt.Println(cmd.Env)
	cmd := exec.Command("/home/ee56054/go_workspace/export.sh")

	// cmd := exec.Command("source ~/.profile")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	for _, s := range syscall.Environ() {
		fmt.Println(s)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
