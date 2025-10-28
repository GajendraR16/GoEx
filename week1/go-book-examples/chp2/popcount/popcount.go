package popcount

var pc [256]byte // Pre-computed table

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) // Count bits for 0-255
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
This function checks if the right most bit is 1
increments the counter and right shift the number
by 1 in binary representation which is equivalent to
x / 2^1
*/
func PopCountloop(x uint64) int {
	count := 0
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}

/*
Loop runs 64 times, irrespective of the value
and number 1's bit present in the binary
representation of the value.
*/

func PopCount64(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			count++
		}
		x = x >> 1
	}
	return count
}

/*
Brian Kernighan's Algorithm
This Algorithm removes the rightmost 1-bit and increments the
1's count, and x is set to the resulting value which is computed
by removing the right most 1-bit
*/
func PopCountEff(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
