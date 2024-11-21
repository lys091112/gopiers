package algorithm

func Gcb(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q

	return Gcb(q, r)
}
