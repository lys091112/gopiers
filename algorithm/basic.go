package algorithm

import sync

func gcb(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q

	// http.DefaultClient.Do()
	return gcb(q, r)
}
