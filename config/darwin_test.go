package config

import (
	"testing"
)

func Test_darwin_ChangePath(t *testing.T) {
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
		{"Test Change Path", fields{goroot: `/users/chatchaivichai/goroot`, profile: `/users/chatchaivichai/.profile`}, args{`/users/chatchaivichai/goroot`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &darwin{
				goroot:  tt.fields.goroot,
				profile: tt.fields.profile,
			}
			if err := l.ChangePath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("darwin.ChangePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_darwin_GetPath(t *testing.T) {
	type fields struct {
		goroot  string
		profile string
	}
	tests := []struct {
		name     string
		fields   fields
		wantPath string
		wantErr  bool
	}{
		{"Test Change Path", fields{goroot: `/users/chatchaivichai/goroot`, profile: `/users/chatchaivichai/.profile`}, `/users/chatchaivichai/goroot`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &darwin{
				goroot:  tt.fields.goroot,
				profile: tt.fields.profile,
			}
			gotPath, err := l.GetPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("darwin.GetPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("darwin.GetPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}
