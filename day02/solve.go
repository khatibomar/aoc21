package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	FORWARD = "forward"
	UP      = "up"
	DOWN    = "down"
)

var (
	UNKNOW_CMD = errors.New("Unkown Command")
)

func solve(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	var cmd string
	var val, depth, hori int
	for {
		_, err := fmt.Fscanf(f, "%s%d", &cmd, &val)
		if err != nil {
			break
		}
		switch cmd {
		case FORWARD:
			hori += val
		case UP:
			depth -= val
		case DOWN:
			depth += val
		default:
			return -1, UNKNOW_CMD
		}
	}
	return depth * hori, nil
}

func solve2(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	var cmd string
	var val, depth, hori, aim int
	for {
		_, err := fmt.Fscanf(f, "%s%d", &cmd, &val)
		if err != nil {
			break
		}
		switch cmd {
		case FORWARD:
			hori += val
			depth += aim * val
		case UP:
			aim -= val
		case DOWN:
			aim += val
		default:
			return -1, UNKNOW_CMD
		}
	}
	return depth * hori, nil
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
