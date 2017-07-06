package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	// create dy slice

	dySlice := make([][]uint8, dy)

	// fill dy slice with dx slice
	for y := range dySlice {
		temp := make([]uint8, dx)

		for x := range temp {
			temp[x] = uint8((x + y) / 2)
		}

		dySlice[y] = temp
	}

	return dySlice

}

func main() {
	pic.Show(Pic)
}
