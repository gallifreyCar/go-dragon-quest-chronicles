package main

import (
	"fmt"
	"github.com/gallifreyCar/go-dragon-quest-chronicles/role"
)

func main() {
	rabbitMan := role.NewHero(role.Attribute{HP: 100, MP: 100, ATK: 110, DEF: 100, SPD: 20, LVL: 1, EXP: 0, Name: "rabbitMan"})
	birdMan := role.NewHero(role.Attribute{HP: 200, MP: 100, ATK: 110, DEF: 100, SPD: 5, LVL: 1, EXP: 0, Name: "birdMan"})
	evil := role.NewDragon(role.Attribute{HP: 1250, MP: 100, ATK: 100, DEF: 100, SPD: 2, LVL: 1, EXP: 0, Name: "evil"})

	go rabbitMan.Fight(evil)
	go birdMan.Fight(evil)

	wg := evil.GetWg()
	wg.Wait()

	fmt.Printf("rabbitMan:%d\n", rabbitMan.AttackCount)
	fmt.Printf("birdMan:%d\n", birdMan.AttackCount)

}
