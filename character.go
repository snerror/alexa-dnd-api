package main

import (
	"fmt"
	"log"
	"math/rand"
)

var (
	PLAYER_CLASS_ROGUE   = "rogue"
	PLAYER_CLASS_FIGHTER = "fighter"
	PLAYER_CLASS_CLERIC  = "cleric"
)

var (
	COMBAT_ABILITY_ON_COOLDOWN = "ability_on_cooldown"
	COMBAT_ATTACK_MISSED       = "attack_missed"
	COMBAT_ENEMY_DAMAGED       = "enemy_damaged"
	COMBAT_ENEMY_KILLED        = "enemy_killed"
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

func (p *Player) AttackEnemy(a Ability, e Enemy) string {
	if a.CD != 0 && a.CurrentCD != 0 {
		return fmt.Sprintf("Ability %s currently on cooldown. You can use it in %d turns.", a.Name, a.CD-a.CurrentCD)
	}
	log.Printf("Ability %s used on %s.\n", a.Name, e.Name)

	diceRoll := DiceRoll()

	if (diceRoll + a.Attack) < e.ArmorClass {
		return fmt.Sprintf("Ability %s missed %s.", a.Name, e.Name)
	}

	e.CurrentHp = e.CurrentHp - a.Damage

	log.Printf("You hit %s for %d.\n", e.Name, a.Damage)

	if e.CurrentHp <= 0 {
		return fmt.Sprintf("Ability %s defeated %s with %d damage.", a.Name, e.Name, a.Damage)
	}

	return fmt.Sprintf("Ability %s hit %s with %d damage.", a.Name, e.Name, a.Damage)
}

func DiceRoll() int {
	i := rand.Intn(20)
	log.Printf("Dice rolled %i.\n", i)

	return i
}
