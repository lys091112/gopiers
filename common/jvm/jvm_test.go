package jvm

import "testing"

func Test_parse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "AddTest",
			args: args{path: "/Users/langle/demo/xianyue/TestData.class"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse(tt.args.path)
		})
	}
}
