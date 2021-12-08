package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func solve(filename string, days int) (int64, error) {
	b, err := ioutil.ReadFile(filename)
	b = b[:len(b)-1]
	if err != nil {
		return -1, err
	}
	in := strings.Split(string(b), ",")
	var fishesState [10]int

	for _, val := range in {
		iv, err := strconv.Atoi(val)
		if err != nil {
			return -1, err
		}
		fishesState[iv]++
	}
	for day := 1; day <= days; day++ {
		if fishesState[0] > 0 {
			fishesState[7] += fishesState[0]
			fishesState[9] += fishesState[0]
			fishesState[0] = 0
		}

		for i := 1; i < len(fishesState); i++ {
			if fishesState[i] > 0 {
				fishesState[i-1] += fishesState[i]
				fishesState[i] = 0
			}
		}
	}

	return sumArr(fishesState), nil
}

func sumArr(arr [10]int) int64 {
	var sum int64
	for _, v := range arr {
		sum += int64(v)
	}
	return sum
}

func main() {
	fmt.Println("----[Part 01]----")
	res, err := solve("input.txt", 80)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("solution: %d\n", res)

	fmt.Println("----[Part 02]----")
	res, err = solve("input.txt", 256)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("solution: %d\n", res)
}
