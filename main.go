package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/gallifreyCar/go-dragon-quest-chronicles/dragon"
	"github.com/gallifreyCar/go-dragon-quest-chronicles/hero"
	"sync"
	"time"
)

type quest struct {
	dx, dy int
}

var m sync.Mutex

func clearScreen() { fmt.Print("\033[H\033[2J") }

func main() {

	fireGuide := hero.Default("fireGuide")
	blueBird := hero.Default("BlueBird")
	pinkRabbit := hero.Default("PinkRabbit")
	pinkRabbit.ATK = 12
	blueBird.ATK = 15
	eval := dragon.Default("Eval")
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	a := make(chan int, 20)
	for i := 0; i < 7; i++ {
		a <- 1
	}
	b := make(chan int, 10)
	for i := 0; i < 5; i++ {
		b <- 1
	}
	c := make(chan int, 15)

	for {
		if eval.IsDead == true {
			fmt.Println("恶龙死亡，任务完成")
			break
		}
		go fireGuide.Attack(a, eval)
		go blueBird.Attack(b, eval)
		go pinkRabbit.Attack(c, eval)

		time.Sleep(1000 * time.Millisecond)
	}

}
