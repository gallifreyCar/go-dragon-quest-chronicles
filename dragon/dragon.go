package dragon

import "sync"

// Dragon 龙的属性
type Dragon struct {
	HP, MP, ATK, DEF, SPD, LVL, EXP int
	Name                            string
	//龙一次只能被一个英雄攻击
	WhoAttack sync.Mutex
}

// Default 默认龙
func Default(name string) *Dragon {
	return &Dragon{
		HP: 100, MP: 100, ATK: 10, DEF: 10, SPD: 10, LVL: 1, EXP: 0,
		Name: name,
	}
}

// NewDragon 新龙
func NewDragon(hp, mp, atk, def, spd, lvl, exp int) *Dragon {
	return &Dragon{
		HP: hp, MP: mp, ATK: atk, DEF: def, SPD: spd, LVL: lvl, EXP: exp,
	}
}
