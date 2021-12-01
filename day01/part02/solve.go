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

	var vals []int

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

		vals = append(vals, newVal)
	}

	newVal = 0

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	for {
		if len(vals) >= 3 {
			for i := 0; i < 3; i++ {
				newVal += vals[i]
			}
		} else {
			for i := 0; i < len(vals); i++ {
				newVal += vals[i]
			}
		}
		if newVal > preVal {
			count++
		}
		preVal = newVal
		newVal = 0
		if len(vals) == 0 {
			break
		}
		vals = vals[1:]
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
