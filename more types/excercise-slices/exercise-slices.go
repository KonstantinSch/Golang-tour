package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mpic := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		mpic[i] = make([]uint8, dx)
		for j := 0; j < dy; j++ {
			mpic[i][j] = uint8((i + j) / 2)
		}
	}
	return mpic
}

func main() {
	pic.Show(Pic)
}
