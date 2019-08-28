package palindromic

import "testing"

func TestPalindromicTree_Add(t *testing.T) {
	NewPalindromicTree(8)
	PalindromicAdd('a')
	PalindromicAdd('b')
	PalindromicAdd('b')
	PalindromicAdd('a')
	PalindromicAdd('a')
	PalindromicAdd('b')
	PalindromicAdd('b')
	PalindromicAdd('a')
}
