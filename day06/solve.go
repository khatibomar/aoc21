package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func solve(filename string) (int, error) {
	b, err := ioutil.ReadFile(filename)
	b = b[:len(b)-1]
	if err != nil {
		return -1, err
	}
	in := strings.Split(string(b), ",")
	var fishes []int
	for _, val := range in {
		iv, err := strconv.Atoi(val)
		if err != nil {
			return -1, err
		}
		fishes = append(fishes, iv)
	}

	for day := 1; day <= 80; day++ {
		for i := 0; i < len(fishes); i++ {
			fishes[i] = fishes[i] - 1
			if fishes[i] == -1 {
				fishes[i] = 6
				fishes = append(fishes, 9)
			}
		}
	}

	return len(fishes), nil
}

func main() {
	fmt.Println("----[Part 01]----")
	res, err := solve("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("solution: %d\n", res)

	// fmt.Println("----[Part 02]----")
	// res, err = solve2("input.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("solution: %d\n", res)
}
