package main

import (
	"fmt"
	"log"
	"os"
)

const (
	width  = 1000
	height = 1000
)

func solve(filename string) (int, error) {
	var grid [width][height]int
	var x1, x2, y1, y2, overlapCount int

	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	for {
		_, err := fmt.Fscanf(f, "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
		if err != nil {
			break
		}
		if x1 == x2 {
			for i := min(y1, y2); i <= max(y1, y2); i++ {
				grid[i][x1]++
			}
		} else if y1 == y2 {
			for i := min(x1, x2); i <= max(x1, x2); i++ {
				grid[y1][i]++
			}
		}
	}
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if grid[i][j] > 1 {
				overlapCount++
			}
		}
	}
	return overlapCount, nil
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println("---- [Part 01] ----")
	ans, err := solve("input.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ans)

	// fmt.Println("---- [Part 02] ----")
	// ans, err = solve2("input.txt")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(ans)
}
