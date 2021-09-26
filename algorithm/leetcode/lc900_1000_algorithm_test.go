package leetcode

import (
	"testing"
)

func TestCanReorderDoubled(t *testing.T) {
	res := canReorderDoubled([]int{-3, -6})
	if res {
		t.Log("success")
	} else {
		t.Errorf("excepted=%t, actual=%t\n", false, res)
	}
}

func TestValidateStackSequences(t *testing.T) {
	pushed := []int{8, 2, 1, 4, 7, 9, 0, 3, 5, 6}
	popped := []int{1, 2, 7, 3, 6, 4, 0, 9, 5, 8}
	res := validateStackSequences(pushed, popped)

	if !res {
		t.Log("success")
	} else {
		t.Error("failed")
	}
}

func TestLargestComponentSize(t *testing.T) {
	s := largestComponentSize([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	// s := largestComponentSize([]int{65, 35, 43, 76, 15, 11, 81, 22, 55, 92, 31})
	resultOfInt(6, s, t)
}

func TestRecentCounter_Ping(t *testing.T) {
	type args struct {
		t []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test01",
			args: args{build933()},
			want: buildWant933(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc := Constructor933()
			for i, v := range tt.args.t {
				if got := rc.Ping(v); got != tt.want[i] {
					t.Errorf("RecentCounter.Ping(%v) = %v, want %v", v, got, tt.want[i])
				}
			}
		})
	}
}
func build933() []int {
	v := make([]int, 3003)
	for i := 0; i < 3003; i++ {
		v[i] = i + 1
	}
	return v
}

func buildWant933() []int {
	v := make([]int, 3003)
	for i := 0; i < 3001; i++ {
		v[i] = i + 1
	}
	v[3001] = 3001
	v[3002] = 3001
	return v
}

func Test_isLongPressedName(t *testing.T) {
	type args struct {
		name  string
		typed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"test01",
			args: args{name:"abd",typed:"aabbddd"},
			want: true,
		},
		{
			name:"test02",
			args: args{name:"saeed",typed:"ssaaedd"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
