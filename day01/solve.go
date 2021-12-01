package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solve(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	var preVal int
	var newVal int

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		newVal, err = strconv.Atoi(line)
		if err != nil {
			return -1, err
		}

		if newVal > preVal {
			count++
		}
		preVal = newVal
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return count - 1, nil
}

func main() {
	count, err := solve("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("solution: %d\n", count)
}
