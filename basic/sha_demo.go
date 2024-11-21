package basic

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(str string) [32]byte {
	return sha256.Sum256([]byte(str))
}

func TestSha() {
	sha256.New()
	fmt.Printf("%x", Sha256("my name is crescent"))
}
