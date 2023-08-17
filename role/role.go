package role

// Attribute 基本属性
type Attribute struct {
	//HP 生命值 MP 魔法值 ATK 攻击力 DEF 防御力 SPD 速度 LVL 等级 EXP 经验值
	HP, MP, ATK, DEF, SPD, LVL, EXP int
	Name                            string
}

type Role interface {
	// Attack 攻击
	Attack(Role)

	// GetAttribute 获取属性
	GetAttribute() *Attribute
	// BeAttack 被攻击
	BeAttack(role Role)
}
