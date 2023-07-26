package hero

import (
	"fmt"
	"github.com/gallifreyCar/go-dragon-quest-chronicles/dragon"
	"math/rand"
	"time"
)

type Hero struct {
	//HP 生命值 MP 魔法值 ATK 攻击力 DEF 防御力 SPD 速度 LVL 等级 EXP 经验值
	HP, MP, ATK, DEF, SPD, LVL, EXP int
	Name                            string
	isSleep                         bool
}

func Default(name string) *Hero {
	return &Hero{
		HP: 100, MP: 100, ATK: 20, DEF: 10, SPD: 10, LVL: 1, EXP: 0,
		Name: name,
	}
}
func NewHero(hp, mp, atk, def, spd, lvl, exp int) *Hero {
	return &Hero{
		HP: hp, MP: mp, ATK: atk, DEF: def, SPD: spd, LVL: lvl, EXP: exp,
	}
}

func (hero *Hero) Attack(atkSignal chan int, dragon *dragon.Dragon) {
	if dragon.IsDead == true {
		return
	}
	if hero.isSleep == true {
		return
	} else if len(atkSignal) <= 0 {
		go hero.Sleep(atkSignal)
		return
	}
	x := <-atkSignal
	dragon.WhoAttack.Lock()
	if x == 0 {
		return
	}
	trulyAtk := hero.ATK + hero.ATK*(rand.Intn(3)-rand.Intn(3)) //攻击力随机波动
	if trulyAtk < 0 {
		trulyAtk = hero.ATK
	}

	if dragon.DEF > trulyAtk {
		hero.HP -= dragon.DEF - trulyAtk
		fmt.Printf("英雄-%s攻击了恶龙-%s,但是攻击力太低，攻击被反噬，英雄-%sHP减少%d，剩余HP:%d\n", hero.Name, dragon.Name, hero.Name, dragon.DEF-trulyAtk, hero.HP)
	}
	if trulyAtk > dragon.DEF {
		dragon.HP -= trulyAtk - dragon.DEF
		if dragon.HP <= 0 {
			dragon.IsDead = true
		}
		fmt.Printf("英雄-%-10s攻击了恶龙-%s，成功攻击，恶龙-%-5sHP减少%-2d，剩余HP:%d\n", hero.Name, dragon.Name, dragon.Name, trulyAtk-dragon.DEF, dragon.HP)
	}
	if trulyAtk == dragon.DEF {
		fmt.Printf("英雄-%s攻击了恶龙-%s,实力相当，无事发生\n", hero.Name, dragon.Name)
	}
	time.Sleep(100 * time.Millisecond)
	dragon.WhoAttack.Unlock()

}

func (hero *Hero) Sleep(atkSignal chan int) {
	hero.isSleep = true
	fmt.Printf("英雄-%-10s正在休息\n", hero.Name)
	for i := 0; i < 2; i++ {
		atkSignal <- 1
	}
	time.Sleep(2 * time.Second)
	hero.isSleep = false
}
