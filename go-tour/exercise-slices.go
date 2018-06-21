package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	img := make([][]uint8, dy)

	for i, _ := range img {
		img[i] = make([]uint8, dx)
	}

	for x, v := range img {
		for y, _ := range v {
			img[x][y] = uint8((x + y) / 2)
		}
	}

	return img
}

func main() {
	pic.Show(Pic)
}
