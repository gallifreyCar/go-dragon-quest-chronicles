package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

type quest struct {
	dx, dy int
}

var gameMap = make([][]string, 3)

func clearScreen() { fmt.Print("\033[H\033[2J") }

func (q *quest) move(r rune, k keyboard.Key) {
	switch r {
	case 'a':
		q.dx--
		gameMap[1][0] = "<"

	case 'd':
		q.dx++
		gameMap[1][2] = ">"

	case 'w':
		q.dy--
		gameMap[0][1] = "^"

	case 's':
		q.dy++
		gameMap[2][1] = "v"

	}

}

func initGameMap(gameMap [][]string) {
	for i := range gameMap {
		gameMap[i] = make([]string, 3)
	}

	for i := range gameMap {
		for j := range gameMap[i] {
			gameMap[i][j] = "*"
		}
	}
}

func printGameMap(gameMap [][]string) {
	for i := range gameMap {
		for j := range gameMap[i] {
			fmt.Print(gameMap[i][j])
		}
		fmt.Println()
	}
}

func main() {

	q := quest{0, 0}
	initGameMap(gameMap)

	// 启动键盘监听s
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		println("x: ", q.dx, " y: ", q.dy)
		printGameMap(gameMap)
		initGameMap(gameMap)

		r, key, err := keyboard.GetKey()
		if r == 'q' {
			break
		}
		if err != nil {
			panic(err)
		}

		q.move(r, key)

		time.Sleep(100 * time.Millisecond)
		clearScreen()
	}
}
