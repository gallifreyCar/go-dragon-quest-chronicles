package main

import (
	"fmt"
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
	pinkRabbit := hero.Default("PinkRab")
	waterElement := hero.Default("WaterEle")
	pinkRabbit.ATK = 5
	blueBird.ATK = 15
	waterElement.ATK = 10
	eval := dragon.Default("Eval")

	a := make(chan int, 20)
	b := make(chan int, 10)
	c := make(chan int, 15)
	d := make(chan int, 5)

	heroList := []*hero.Hero{fireGuide, blueBird, pinkRabbit, waterElement}
	//dragonList := []*dragon.Dragon{eval}
	signalList := []chan int{a, b, c, d}
	allDead := false
	for {

		for _, h := range heroList {
			if h.IsDead == false {
				allDead = false
				break
			}
			allDead = true
		}
		if allDead == true {
			fmt.Println("英雄全部阵亡，任务失败")
			break
		}
		if eval.IsDead == true {
			fmt.Println("--------------------------")
			fmt.Println("恶龙阵亡，任务成功")
			for _, h := range heroList {
				if h.IsDead == true {
					fmt.Printf("英雄-%-10s阵亡\n", h.Name)
				} else {
					fmt.Printf("英雄-%-10s剩下生命值%-2d，累计造成伤害值%-3d，累计攻击次数%-3d\n", h.Name, h.HP, h.Damage, h.AttackTimes)
				}

			}
			fmt.Println("--------------------------")
			break
		}

		for i, h := range heroList {
			if h.IsDead == false {
				go h.Attack(signalList[i], eval)
			}

		}

		time.Sleep(1000 * time.Millisecond)
	}

}
