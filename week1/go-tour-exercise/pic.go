package main

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := range dy {
		picture[y] = make([]uint8, dx)
		for x := range dx {
			picture[y][x] = uint8(x * y)
		}
	}
	return picture
}

func main() {
	dx := 2
	dy := 5
	Pic(dx, dy)
}
