package algorithm

func gcb(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcb(q, r)
}
