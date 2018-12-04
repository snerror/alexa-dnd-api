package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

var (
	PlayerClassRogue   = "rogue"
	PlayerClassFighter = "fighter"
	PlayerClassCleric  = "cleric"
)

type Player struct {
	Name       string
	Class      string
	CurrentHp  int
	MaxHp      int
	ArmorClass int
	Level      int
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
	if t != PlayerClassCleric && t != PlayerClassFighter && t != PlayerClassRogue {
		return fmt.Errorf("unknown class provided")
	}

	if t == PlayerClassRogue {
		*p = TemplateClassRogue
	}

	if t == PlayerClassCleric {
		*p = TemplateClassCleric
	}

	if t == PlayerClassFighter {
		*p = TemplateClassFighter
	}

	p.Experience = 0
	p.Level = 1
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

		defeatedEnemyText := fmt.Sprintf("Ability %s defeated %s with %d damage.", a.Name, e.Name, a.Damage)
		levelUpText := ""

		if p.Experience > 5 {
			p.MaxHp += 2
			p.CurrentHp = p.MaxHp
			p.ArmorClass += 1
			p.Experience = 0
			p.Level++

			levelUpText = " Congratulations you are now level " + strconv.Itoa(p.Level) + " , your stats increased and your health refiled. Your HP is now " + strconv.Itoa(p.MaxHp) + " and your Armour Class is " + strconv.Itoa(p.ArmorClass) + "."
		}

		return defeatedEnemyText + levelUpText
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
