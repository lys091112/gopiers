package leetcode

// 22 括号生成
// f(n) = ['(' + f(n-1) + ')', f(n-1) + '()' , '()' + f(n-1)]
// 排除掉都是单个括号的情况类似()()()()这种
// TODO 错误的 重复情况太多 没有考虑这种 (())(())

var lc_21_30_result []string

func generateParenthesis(n int) []string {

	lc_21_30_result = make([]string, 0)
	doGenerateParenthesis("", 0, 0, n)

	return lc_21_30_result
}

func doGenerateParenthesis(str string, left int, right int, n int) {
	if left == right && left == n {
		lc_21_30_result = append(lc_21_30_result, str)
		return
	}

	if left <= n {
		doGenerateParenthesis(str+"(", left+1, right, n)
	}

	if right < left {
		doGenerateParenthesis(str+")", left, right+1, n)
	}
}
