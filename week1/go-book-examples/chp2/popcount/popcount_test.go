package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCount64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64(uint64(i))
	}
}

func BenchmarkPopCountloop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountloop(uint64(i))
	}
}

func BenchmarkPopCountEff(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountEff(uint64(i))
	}
}
