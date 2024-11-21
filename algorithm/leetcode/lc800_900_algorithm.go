package leetcode

// N: 856 括号的分数
// st 栈顶记录可计算的括号分数的值
func scoreOfParentheses(s string) int {
	if s == "" {
		return 0
	}

	st := []int{0}

	for _, v := range s {
		if v == '(' {
			st = append(st, 0)
		} else {
			top := st[len(st)-1]
			st = st[0 : len(st)-1]
			st[len(st)-1] += If(top > 0, 2*top, 1).(int)

		}
	}
	return st[0]
}
