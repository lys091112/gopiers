package weekly

import (
	"testing"
)

func Test_checkOnesSegment(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{s: "1001"},
			want: false,
		},
		{
			name: "test02",
			args: args{s: "1100"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkOnesSegment(tt.args.s); got != tt.want {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minElements(t *testing.T) {
	type args struct {
		nums  []int
		limit int
		goal  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{1, -10, 9, 1}, limit: 100, goal: 0},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minElements(tt.args.nums, tt.args.limit, tt.args.goal); got != tt.want {
				t.Errorf("minElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countRestrictedPaths(t *testing.T) {
	type args struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test01",
			args:args{n:5,edges: [][]int{{1,2,3},{1,3,3},{2,3,1},{1,4,2},{5,2,2},{3,5,1},{5,4,10}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countRestrictedPaths(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("countRestrictedPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
