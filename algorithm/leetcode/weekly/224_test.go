package weekly

import (
	"testing"
)

func Test_countGoodRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{rectangles: [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodRectangles(tt.args.rectangles); got != tt.want {
				t.Errorf("countGoodRectangles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tupleSameProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{nums: []int{20, 10, 6, 24, 15, 5, 4, 30}},
			want: 48,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tupleSameProduct(tt.args.nums); got != tt.want {
				t.Errorf("tupleSameProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestSubmatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{matrix:[][]int{{0,0,1},{1,1,1},{1,0,1}}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestSubmatrix(tt.args.matrix); got != tt.want {
				t.Errorf("largestSubmatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
