package leetcode

// I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
func romanToInt(s string) int {

	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["IV"] = 4
	m["X"] = 10
	m["IX"] = 9
	m["L"] = 50
	m["XL"] = 40
	m["C"] = 100
	m["XC"] = 90
	m["D"] = 500
	m["CD"] = 400
	m["M"] = 1000
	m["CM"] = 900

	sum := 0

	// I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
	//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
	//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
	// 考虑好边界情况
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			sum += m[s[i:]]
		} else {
			if value, ok := m[s[i:i+2]]; ok {
				sum += value
				i += 1
			} else {
				sum += m[s[i:i+1]]
			}
		}

	}

	return sum
}
