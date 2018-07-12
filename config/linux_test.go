package config

import (
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
		gopath  string
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
		{"Test Change Path", fields{gopath: `$HOME/go_workspace`, profile: `/home/ee56054/.profile`}, args{`/home/ee56054/go101`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &linux{
				gopath:  tt.fields.gopath,
				profile: tt.fields.profile,
			}
			if err := l.ChangePath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("linux.ChangePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
