package main

import (
	"fmt"
	"strconv"
	"strings"
)

type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place
func Sort(values []int) *tree {
	var root *tree
	for _, val := range values {
		root = add(root, val)
	}
	return root
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	var values []int

	values = appendValues(values, t)
	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = strconv.Itoa(v)
	}

	// Format as a comma-separated sequence
	return strings.Join(strValues, ", ")
}

func main() {
	values := []int{3, 1, 4, 2}
	root := Sort(values)
	fmt.Println(root)
}
