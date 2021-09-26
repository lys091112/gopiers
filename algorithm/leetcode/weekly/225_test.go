package weekly

import (
	"testing"
)

func Test_maximumTime(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test01",
			args: args{time: "2?:4?"},
			want: "23:49",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumTime(tt.args.time); got != tt.want {
				t.Errorf("maximumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCharacters(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{a: "ace", b: "abe"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCharacters(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kthLargestValue(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name:"test01",
			args:args{matrix: [][]int{{10,9,5},{2,0,4},{1,0,9},{3,4,8}}, k: 8},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargestValue(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthLargestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
