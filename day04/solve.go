package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// TODO(khatibomar): maybe use concurrency
type Board [5][5]int

func (b *Board) Place(val int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] == val {
				b[i][j] = -1
			}
		}
	}
}

func (b Board) Score(curr int) int {
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] != -1 {
				score += b[i][j]
			}
		}
	}
	return curr * score
}

type Game struct {
	input  []int
	boards []Board
}

// TODO(khatibomar): the board worth checking is the board
// that have a column or row need 1 more value to be filled
func (g Game) CheckWinBoards() []Board {
	var winningBoards []Board
	for _, board := range g.boards {
		for i := 0; i < 5; i++ {
			countCol := 0
			countRow := 0
			for j := 0; j < 5; j++ {
				if board[i][j] == -1 {
					countRow++
				}
				if board[j][i] == -1 {
					countCol++
				}
				if countCol == 5 || countRow == 5 {
					winningBoards = append(winningBoards, board)
				}
			}
		}
	}
	return winningBoards
}

func (g *Game) DeleteBoard(board Board) {
	for i, b := range g.boards {
		if b == board {
			g.boards = append(g.boards[:i], g.boards[i+1:]...)
			break
		}
	}
}

func solve(filename string) (int, error) {
	game, err := initGame(filename)
	if err != nil {
		fmt.Println(err)
		return -1, nil
	}
	for _, in := range game.input {
		for i := 0; i < len(game.boards); i++ {
			game.boards[i].Place(in)
		}
		winningBoards := game.CheckWinBoards()
		if len(winningBoards) > 0 {
			return winningBoards[0].Score(in), nil
		}
	}
	return -1, nil
}

func solve2(filename string) (int, error) {
	game, err := initGame(filename)
	if err != nil {
		fmt.Println(err)
		return -1, nil
	}
	var lastWinningBaord Board
	var lastIn int
	for _, in := range game.input {
		for i := 0; i < len(game.boards); i++ {
			game.boards[i].Place(in)
		}
		winningBoards := game.CheckWinBoards()
		if len(winningBoards) > 0 {
			for _, board := range winningBoards {
				lastIn = in
				lastWinningBaord = board
				game.DeleteBoard(board)
			}
		}
	}
	return lastWinningBaord.Score(lastIn), nil
}

func initGame(filename string) (*Game, error) {
	var game Game
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fs := bufio.NewScanner(f)
	fs.Scan()
	for _, v := range strings.Split(fs.Text(), ",") {
		iv, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		game.input = append(game.input, iv)
	}

	for fs.Scan() {
		var board Board
		if fs.Text() == " " || fs.Text() == "\n" || fs.Text() == "" {
			continue
		}
		for i := 0; i < 5; i++ {
			for j, v := range splitString(fs.Text(), " ") {
				iv, err := strconv.Atoi(string(v))
				if err != nil {
					return nil, err
				}
				board[i][j] = iv
			}
			fs.Scan()
		}
		game.boards = append(game.boards, board)
	}
	return &game, nil
}

func splitString(s, delim string) []string {
	ss := strings.Split(s, delim)
	var ns []string
	for _, v := range ss {
		if v == "" || v == "\n" {
			continue
		}
		ns = append(ns, v)
	}
	return ns
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
