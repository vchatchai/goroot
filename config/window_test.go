package config

import (
	"testing"
)

func Test_window_GetPath(t *testing.T) {
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
			l := &window{
				gopath: tt.fields.gopath,
			}
			gotPath, err := l.GetPath()
			if (err != nil) != tt.wantErr {
				t.Errorf("window.GetPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotPath != tt.wantPath {
				t.Errorf("window.GetPath() = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}

func Test_window_ChangePath(t *testing.T) {
	type fields struct {
		gopath string
	}
	type args struct {
		path string
	}

	f := fields{"C:\\Users\\gl\\go"}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestWindowChangePath", f, args{f.gopath}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &window{
				gopath: tt.fields.gopath,
			}
			if err := l.ChangePath(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("window.ChangePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
