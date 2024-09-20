package basic

import (
	"testing"
)

func TestMake_Use(t *testing.T) {
	make_use()
}

func TestSlice_demo(t *testing.T) {
	slice_demo()
}

func TestSliceAppendDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "d1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceAppendDemo()
		})
	}
}
