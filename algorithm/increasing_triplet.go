package algorithm

func IncreasingTriplet(elements []int) bool {
	if len(elements) < 3 {
		return false
	}

	var first = 1 << 31
	var second = 1 << 31

	for _, v := range elements {
		if v < first {
			first = v
			continue
		}

		if v > first && v < second {
			second = v
			continue
		}

		if v > second {
			return true
		}
	}

	return false
}
