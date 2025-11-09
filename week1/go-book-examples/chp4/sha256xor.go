package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(BitDiff(c1, c2))
}

func BitDiff(hash1, hash2 [32]byte) int {
	count := 0
	for i := range len(hash1) {
		count += popCount(int(hash1[i] ^ hash2[i]))
	}
	return count
}

func popCount(x int) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
