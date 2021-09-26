package weekly

import (
	"reflect"
	"testing"
)

func Test_countBalls(t *testing.T) {
	type args struct {
		lowLimit  int
		highLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{lowLimit: 5, highLimit: 15},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalls(tt.args.lowLimit, tt.args.highLimit); got != tt.want {
				t.Errorf("countBalls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_restoreArray(t *testing.T) {
	type args struct {
		adjacentPairs [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{adjacentPairs: [][]int{{4, -2}, {1, 4}, {-3, 1}}},
			want: []int{-2, 4, 1, -3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreArray(tt.args.adjacentPairs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canEat(t *testing.T) {
	type args struct {
		candiesCount []int
		queries      [][]int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "test01",
			args: args{candiesCount: []int{24,41,16,36,6,38,8,43,43,45,17,13,4,44,43,6,10,38,5,23,3,30,5,14,20,31,24,19,32,7,3,20,15,46,8,39,18,21,41,48,39,26,16,46,6,13,10,18,46,25,28,34,31,26,13,8,32,32,49,26,7,46,19,2,24,19,26,22,39,24,48,10,17,10,38,41,48,1,29,30,1,1,27,36,29,37,11,31,5,24,9,33,9,36,3,33,46,49,36},queries:[][]int{{39,466,2}}},
			want: []bool{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEat(tt.args.candiesCount, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}
