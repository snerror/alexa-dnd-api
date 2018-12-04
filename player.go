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

type Player struct {
	Name       string
	Class      string
	CurrentHp  int
	MaxHp      int
	ArmorClass int
	Experience int
	Abilities  []Ability
	Position   Position
}

type Ability struct {
	Name      string
	Attack    int
	Damage    int
	CD        int
	CurrentCD int
}

type Position struct {
	X int
	Y int
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
			Attack:    10,
			Damage:    2,
			CD:        3,
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

	p.Experience = 0
	p.Position = Position{0, 0}

	return nil
}

func (p *Player) AttackEnemy(a *Ability, e *Enemy) string {
	if a.CD != 0 && a.CurrentCD != 0 {
		return fmt.Sprintf("Ability %s currently on cooldown. You can use it in %d turns.", a.Name, a.CurrentCD)
	}
	log.Printf("PLAYER used %s on %s.\n", a.Name, e.Name)

	diceRoll := DiceRoll()

	if (diceRoll + a.Attack) < e.ArmorClass {
		log.Printf("PLAYER missed %s.\n", e.Name)
		return fmt.Sprintf("You missed %s with %s.", e.Name, a.Name)
	}

	e.CurrentHp = e.CurrentHp - a.Damage

	log.Printf("PLAYER hits %s for %d.\n", e.Name, a.Damage)

	if a.CD > 0 && a.CurrentCD == 0 {
		a.CurrentCD = a.CD
		log.Printf("PLAYER ability %s cooldown set to %d.\n", a.Name, a.CurrentCD)
	}

	if e.CurrentHp <= 0 {
		log.Printf("PLAYER kills %s.\n", e.Name)
		p.Experience += e.Experience
		return fmt.Sprintf("Ability %s defeated %s with %d damage.", a.Name, e.Name, a.Damage)
	}

	return fmt.Sprintf("Ability %s hit %s with %d damage.", a.Name, e.Name, a.Damage)
}

func (p *Player) Move(d *Dungeon, strDirection string) string {
	direction := 0

	if strDirection == "up" {
		direction = up
	}

	if strDirection == "down" {
		direction = down
	}

	if strDirection == "left" {
		direction = left
	}

	if strDirection == "right" {
		direction = right
	}

	if direction == 0 {
		return "Unknown direction given. Try again."
	}

	if d.CheckMoveDirection(p.Position.X, p.Position.Y, direction) == false {
		return "Unable to move there. Try again."
	}

	if direction == up {
		player.Position.X--
	}

	if direction == down {
		player.Position.X++
	}

	if direction == right {
		player.Position.Y++
	}

	if direction == left {
		player.Position.Y--
	}

	return "You moved deeper into the dark hallway."
}

func DiceRoll() int {
	i := rand.Intn(20)
	log.Printf("Dice rolled %d.\n", i)

	return i
}
