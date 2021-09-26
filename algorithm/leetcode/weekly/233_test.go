package weekly

import (
	"testing"
)

func Test_maxAscendingSum(t *testing.T) {
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
			args: args{nums: []int{10, 20, 30, 5, 10, 50}},
			want: 65,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAscendingSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAscendingSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getNumberOfBacklogOrders(t *testing.T) {
	type args struct {
		orders [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {
		// 	name: "test01",
		// 	args: args{[][]int{{7,1000000000,1},{15,3,0},{5,999999995,0},{5,1,1}}},
		// 	want: 999999984,
		// },
		// {
		// 	name: "test02",
		// 	args: args{[][]int{{19,28,0},{9,4,1},{25,15,1}}},
		// 	want: 39,
		// },
		{
			name: "test03",
			args: args{[][]int{{27,30,0},{10,10,1},{28,17,1},{19,28,0},{16,8,1},{14,22,0},{12,18,1},{3,15,0},{25,6,1}}},
			want: 82,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNumberOfBacklogOrders(tt.args.orders); got != tt.want {
				t.Errorf("getNumberOfBacklogOrders() = %v, want %v", got, tt.want)
			}
		})
	}

}

// N:5696. 统计异或值在范围内的数对有多少
func countPairs(nums []int, low int, high int) int {

}