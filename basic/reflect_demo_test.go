package basic

import "testing"

func Test_reflect_demo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name:"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reflect_demo()
		})
	}
}
