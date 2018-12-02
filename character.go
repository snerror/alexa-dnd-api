package main

import "fmt"

var (
	PLAYER_CLASS_ROGUE   = "rogue"
	PLAYER_CLASS_FIGHTER = "fighter"
	PLAYER_CLASS_CLERIC  = "cleric"
)

type Player struct {
	Name       string
	Class      string
	CurrentHp  int
	MaxHp      int
	ArmorClass int
	Abilities  []Ability
}

type Ability struct {
	Name      string
	Attack    int
	Damage    int
	CD        int
	CurrentCD int
}

func (p *Player) CreateFromTemplate(t string) error {
	if t != PLAYER_CLASS_CLERIC && t != PLAYER_CLASS_FIGHTER && t != PLAYER_CLASS_ROGUE {
		return fmt.Errorf("unknown class provided")
	}

	if t == PLAYER_CLASS_ROGUE {
		p.Name = "Kat"
		p.Class = PLAYER_CLASS_ROGUE
		p.CurrentHp = 8
		p.MaxHp = 8
		p.ArmorClass = 14

		a1 := Ability{
			Name:      "Backstab",
			Attack:    7,
			Damage:    1,
			CD:        0,
			CurrentCD: 0,
		}

		a2 := Ability{
			Name:      "Strike",
			Attack:    7,
			Damage:    1,
			CD:        0,
			CurrentCD: 0,
		}

		p.Abilities = append(p.Abilities, a1)
		p.Abilities = append(p.Abilities, a2)
	}

	return nil
}
