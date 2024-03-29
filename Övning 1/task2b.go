package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dy)

	for x := range pic {

		pic[x] = make([]uint8, dx)
		for y := 0; y < len(pic[x]); y++ {
			pic[x][y] = uint8(x * y)

		}
	}

	return pic

}

func main() {
	pic.Show(Pic)
}
