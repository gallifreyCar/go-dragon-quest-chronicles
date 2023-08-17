package role

import (
	"fmt"
	"sync"
	"time"
)

type Hero struct {
	attribute   Attribute
	AttackCount int

	mu         sync.Mutex
	isRest     bool
	attackOver chan bool
}

func (h *Hero) BeAttack(role Role) {
	//TODO implement me
	panic("implement me")
}

func (h *Hero) GetAttribute() *Attribute {
	return &h.attribute
}

func NewHero(attribute Attribute) *Hero {
	return &Hero{attribute: attribute, isRest: false}

}

func (h *Hero) Attack(r Role) {
	if r.GetAttribute().HP <= 0 {
		return
	}
	h.Rest()
	h.AttackCount++
	r.BeAttack(h)

	time.Sleep(time.Millisecond * 100)

}

func (h *Hero) Fight(r Role) {
	go func() {
		for {
			if r.GetAttribute().HP <= 0 {
				return
			}
			h.Attack(r)
		}
	}()

}

var Speed = 200

func (h *Hero) Rest() {
	if h.isRest {
		return
	}
	h.isRest = true
	fmt.Printf("%s休息中\n", h.attribute.Name)
	t := 0
	for t <= Speed {
		t += h.attribute.SPD
		time.Sleep(time.Millisecond * 10)
	}
	h.isRest = false
	fmt.Printf("%s休息结束\n", h.attribute.Name)
}

func (h *Hero) Over() {
}

var _ Role = (*Hero)(nil)
