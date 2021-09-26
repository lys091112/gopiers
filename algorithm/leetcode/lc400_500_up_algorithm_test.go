package leetcode

import "testing"

func Test_findMaxForm(t *testing.T) {
	type args struct {
		strs []string
		m    int
		n    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test01",
			args: args{[]string{"10", "0001", "111001", "1", "0"},5,3},
			want: 4,
		},
		{
			name: "test02",
			args: args{[]string{"11111","100","1101","1101","11000"},5,7},
			want: 3,
		},
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForm(tt.args.strs, tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("findMaxForm() = %v, want %v", got, tt.want)
			}
		})
	}
}
