package main

import "fmt"

type IntSet struct {
	words []uint64
}

// Add adds x to the set
func (s *IntSet) Add(x int) {
	word := x / 64      // Which uint64?
	bit := uint(x % 64) // Which bit?

	// Expand slice if needed
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}

	s.words[word] |= 1 << bit // Set the bit
}

// Has reports whether x is in the set
func (s *IntSet) Has(x int) bool {
	word := x / 64
	bit := uint(x % 64)

	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		for word != 0 {
			word = word & (word - 1)
			count++
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint8(x%64)

	if word >= len(s.words) {
		return
	}

	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{
		words: make([]uint64, len(s.words)),
	}
	copy(newSet.words, s.words)
	return newSet
}

func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

// String returns a string representation like "{1 3 5}"
func (s *IntSet) String() string {
	result := "{"
	for i, word := range s.words {
		for bit := 0; bit < 64; bit++ {
			if word&(1<<uint(bit)) != 0 {
				if len(result) > 1 {
					result += " "
				}
				result += fmt.Sprintf("%d", 64*i+bit)
			}
		}
	}
	return result + "}"
}

func main() {
	var s IntSet

	// Test AddAll
	s.AddAll(1, 5, 10, 70, 130)
	fmt.Println("After AddAll:", s.String())
	fmt.Println("Len:", s.Len())

	// Test Has
	fmt.Println("Has 5?", s.Has(5))
	fmt.Println("Has 100?", s.Has(100))

	// Test Remove
	s.Remove(5)
	fmt.Println("After Remove(5):", s.String())
	fmt.Println("Has 5?", s.Has(5))

	// Test Copy
	copy := s.Copy()
	copy.Add(999)
	fmt.Println("Original:", s.String())
	fmt.Println("Copy:", copy.String())

	// Test Clear
	s.Clear()
	fmt.Println("After Clear:", s.String())
	fmt.Println("Len:", s.Len())
}