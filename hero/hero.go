package hero

import "fmt"

type Hero struct {
	//HP 生命值 MP 魔法值 ATK 攻击力 DEF 防御力 SPD 速度 LVL 等级 EXP 经验值
	HP, MP, ATK, DEF, SPD, LVL, EXP int
}

func Default() *Hero {
	return &Hero{
		HP: 100, MP: 100, ATK: 10, DEF: 10, SPD: 10, LVL: 1, EXP: 0,
	}
}
func NewHero(hp, mp, atk, def, spd, lvl, exp int) *Hero {
	return &Hero{
		HP: hp, MP: mp, ATK: atk, DEF: def, SPD: spd, LVL: lvl, EXP: exp,
	}
}

func (hero *Hero) Attack(atkSignal chan int) {
	x := <-atkSignal
	hero.EXP += x
	if x != 0 {
		fmt.Printf("Hero Attack, ATK:%d,EXP:+%d \n", hero.ATK, x)
	}
}
