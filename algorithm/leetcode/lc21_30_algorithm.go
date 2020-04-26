package leetcode

// 22 括号生成
// f(n) = ['(' + f(n-1) + ')', f(n-1) + '()' , '()' + f(n-1)]
// 排除掉都是单个括号的情况类似()()()()这种
func generateParenthesis(n int) []string {

	result := make([]string, 0)
	doGenerateParenthesis(&result, "", 0, 0, n)
	return result
}

func doGenerateParenthesis(result *[]string, str string, left int, right int, n int) {
	if left == right && left == n {
		*result = append(*result, str)
		return
	}

	if left <= n {
		doGenerateParenthesis(result, str+"(", left+1, right, n)
	}

	if right < left {
		doGenerateParenthesis(result ,str+")", left, right+1, n)
	}
}
