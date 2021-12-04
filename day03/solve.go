package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solve(filename string) (int64, error) {
	var s []int
	var b []byte
	var err error
	var width, height int

	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	first := true

	for {
		b, err = bufReader.ReadBytes('\n')
		if err != nil {
			break
		}
		if first {
			width = len(b) - 1
			first = false
		}

		for _, v := range b {
			if int(v) != 10 {
				s = append(s, int(v)-48)
			} else {
				height++
			}
		}
	}

	var gamma, eps string
	var ones, zeros int

	for i := 0; i < width; i++ {
		for j := 0; ; j += width {
			if i+j >= len(s) {
				break
			}
			if s[i+j] == 1 {
				ones++
			} else {
				zeros++
			}
		}
		if major(zeros, ones) == 1 {
			gamma += "1"
			eps += "0"
		} else {
			gamma += "0"
			eps += "1"
		}
		ones = 0
		zeros = 0
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(eps, 2, 64)
	return g * e, nil
}

func major(zeros, ones int) int {
	if zeros > ones {
		return 0
	}
	return 1
}

func main() {
	fmt.Println("---- [Part 01] ----")
	ans, err := solve("input.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ans)
}
