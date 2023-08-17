package role

import (
	"fmt"
	"math/rand"
	"sync"
)

// Dragon 龙的属性
type Dragon struct {
	attribute Attribute
	mu        sync.Mutex
	wg        sync.WaitGroup
}

func (d *Dragon) BeAttack(role Role) {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Println("------------------------------------------------------------")
	if d.attribute.HP <= 0 {
		d.wg.Done()
		fmt.Printf("%s已死亡\n", d.GetAttribute().Name)
		fmt.Println("------------------------------------------------------------")
		return
	}

	randDom := rand.Intn(3)
	atk := role.GetAttribute().ATK * randDom
	fmt.Printf("%-8s攻击 %-2s,", role.GetAttribute().Name, d.GetAttribute().Name)
	if atk > d.attribute.DEF {
		d.attribute.HP -= atk - d.attribute.DEF
		fmt.Printf("攻击有效,造成%-4d伤害,%s剩余血量:%d\n", atk-d.attribute.DEF, d.attribute.Name, d.attribute.HP)
	} else {
		fmt.Printf("攻击无效\n")
	}
	fmt.Println("------------------------------------------------------------")
}

func (d *Dragon) GetAttribute() *Attribute {
	return &d.attribute
}

func NewDragon(attribute Attribute) *Dragon {
	d := Dragon{
		attribute: attribute,
		wg:        sync.WaitGroup{},
	}
	d.wg.Add(1)
	return &d

}

func (d *Dragon) Attack(role Role) {
	//TODO implement me
	panic("implement me")
}

func (d *Dragon) Rest() {
	//TODO implement me
	panic("implement me")
}

func (d *Dragon) GetWg() *sync.WaitGroup {
	return &d.wg
}
