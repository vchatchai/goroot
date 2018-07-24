package config

import (
	"testing"
)

func Test_darwin_GetPath(t *testing.T) {
	type fields struct {
		gopath string
	}
	tests := []struct {
		name     string
		fields   fields
		wantPath string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &darwin{
				gopath: tt.fields.gopath,
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

func Test_darwin_ChangePath(t *testing.T) {
	type fields struct {
		gopath string
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &darwin{
				gopath: tt.fields.gopath,
			}
			if err := l.ChangePath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("darwin.ChangePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
