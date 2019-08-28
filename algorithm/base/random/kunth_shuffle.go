package random

import "math/rand"

// kunth-shuffle 洗牌算法 保证每个位置被填充的概率都是1/length
// 每个位置都等于以前未选中和这次被选中的概率乘积
func KunthShuffle(values []int) (res []int) {
	if values == nil || len(values) <= 0 {
		return
	}

	for i := len(values) - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		if r == i {
			continue
		}

		values[i], values[r] = values[r], values[i]
	}

	return values
}
