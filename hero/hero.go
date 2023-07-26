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
	IsDead                          bool
	Damage                          int //累计伤害
	AttackTimes                     int //累计攻击次数

}

func Default(name string) *Hero {
	return &Hero{
		HP: 100, MP: 100, ATK: 20, DEF: 10, SPD: 10, LVL: 1, EXP: 0,
		Name:        name,
		Damage:      0,
		AttackTimes: 0,
		IsDead:      false,
	}
}
func NewHero(hp, mp, atk, def, spd, lvl, exp int) *Hero {
	return &Hero{
		HP: hp, MP: mp, ATK: atk, DEF: def, SPD: spd, LVL: lvl, EXP: exp,
	}
}

func (hero *Hero) Attack(atkSignal chan int, dragon *dragon.Dragon) {

	if hero.isSleep == true {
		return
	} else if len(atkSignal) <= 0 {
		go hero.Sleep(atkSignal)
		return
	}
	//龙每次只能被一个英雄攻击
	dragon.WhoAttack.Lock()
	if dragon.IsDead == true || hero.IsDead == true {
		return
	}
	x := <-atkSignal
	hero.AttackTimes++
	if x == 0 {
		return
	}
	//攻击力随机波动
	trulyAtk := hero.ATK - hero.ATK*rand.Intn(2) + hero.ATK*rand.Intn(3)
	if trulyAtk < 0 {
		trulyAtk = hero.ATK
	}

	if dragon.DEF > trulyAtk {
		hero.HP -= dragon.DEF - trulyAtk

		fmt.Printf("英雄-%-10s攻击了恶龙-%-5s,但是攻击力太低，攻击被反噬，英雄-%-10sHP减少%-2d，剩余HP:%d\n", hero.Name, dragon.Name, hero.Name, dragon.DEF-trulyAtk, hero.HP)
		if hero.HP <= 0 {
			hero.IsDead = true
			fmt.Printf("英雄-%-10s阵亡\n", hero.Name)
		}
	}
	if trulyAtk > dragon.DEF {
		dragon.HP -= trulyAtk - dragon.DEF
		hero.Damage += trulyAtk - dragon.DEF
		fmt.Printf("英雄-%-10s攻击了恶龙-%s，成功攻击,恶龙-%-5sHP减少%-2d，剩余HP:%d\n", hero.Name, dragon.Name, dragon.Name, trulyAtk-dragon.DEF, dragon.HP)
		if dragon.HP <= 0 {
			dragon.IsDead = true
		}
	}
	if trulyAtk == dragon.DEF {
		fmt.Printf("英雄-%-10s攻击了恶龙-%-5s,实力相当，无事发生\n", hero.Name, dragon.Name)
	}
	time.Sleep(100 * time.Millisecond)
	dragon.WhoAttack.Unlock()

}

func (hero *Hero) Sleep(atkSignal chan int) {
	hero.isSleep = true
	fmt.Printf("英雄-%-10s正在休息\n", hero.Name)
	for i := 0; i < cap(atkSignal); i++ {
		atkSignal <- 1
		time.Sleep(500 * time.Millisecond)
	}

	hero.isSleep = false
}
