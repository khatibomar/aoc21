package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func solve2(filename string) (int, error) {
	oxygen, err := calculate(filename, func(i1, i2 int) byte {
		if i1 <= i2 {
			return '0'
		}
		return '1'
	})
	if err != nil {
		return -1, err
	}
	co2, err := calculate(filename, func(i1, i2 int) byte {
		if i1 <= i2 {
			return '1'
		}
		return '0'
	})
	if err != nil {
		return -1, err
	}
	return oxygen * co2, nil
}

func calculate(filename string, minor func(int, int) byte) (int, error) {
	lines, err := getLines(filename)
	if err != nil {
		return -1, err
	}
	var currWidth int
	for len(lines) > 1 {
		zeros := 0
		ones := 0
		for _, line := range lines {
			if line[currWidth] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		i := 0
		for len(lines) > 0 && i < len(lines) {
			if lines[i][currWidth] == minor(zeros, ones) {
				lines = append(lines[:i], lines[i+1:]...)
				i = 0
				continue
			}
			i++
		}

		currWidth++
	}
	val, err := strconv.ParseInt(lines[0], 2, 64)
	return int(val), err
}

func getLines(filename string) ([]string, error) {
	var lines []string
	f, err := os.Open(filename)
	if err != nil {
		return lines, err
	}
	defer f.Close()
	fs := bufio.NewScanner(f)

	for fs.Scan() {
		lines = append(lines, fs.Text())
	}

	return lines, nil
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

	fmt.Println("---- [Part 02] ----")
	ans, err = solve2("input.txt")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(ans)
}
