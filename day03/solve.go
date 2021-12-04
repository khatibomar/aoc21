package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solve(filename string) (int, error) {
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

	var gamma, eps, ones, zeros int

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
		gamma = gamma<<1 + major(zeros, ones)
		eps = eps<<1 + (1 - major(zeros, ones))

		ones = 0
		zeros = 0
	}
	return gamma * eps, nil
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
