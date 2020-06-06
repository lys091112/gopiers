package fileread

import (
	"testing"
)

func TestReadUseBufio(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test01",
			args: args{path: "fileread.go"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadUseBufio(tt.args.path)
		})
	}
}

func TestBaseUse(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BaseUse()
		})
	}
}
