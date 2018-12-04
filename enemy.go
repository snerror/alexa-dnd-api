package main

import (
	"fmt"
	"log"
)

var (
	EnemySkeleton        = "skeleton"
	EnemySpider          = "spider"
	EnemyWraith          = "wraith"
	EnemyZombie          = "zombie"
	EnemyGargoyle        = "gargoyle"
	EnemyBlazingSkeleton = "blazing skeleton"
)

var EnemyId = 0

type Enemy struct {
	ID         int
	Name       string
	CurrentHp  int
	MaxHp      int
	ArmorClass int
	Experience int
	Abilities  []Ability
	Position   Position
}

func (e *Enemy) CreatePreset(p string, x int, y int) error {
	if p == EnemySkeleton {
		*e = TemplateEnemySkeleton
	}

	if p == EnemySpider {
		*e = TemplateEnemySpider
	}

	if p == EnemyWraith {
		*e = TemplateEnemyWraith
	}

	if p == EnemyZombie {
		*e = TemplateEnemyZombie
	}

	if p == EnemyGargoyle {
		*e = TemplateEnemyGargoyle
	}

	if p == EnemyBlazingSkeleton {
		*e = TemplateEnemyBlazingSkeleton
	}

	if &p == nil {
		return fmt.Errorf("unknown enemy preset %s \n", p)
	}

	e.Position = Position{x, y}
	e.ID = EnemyId

	EnemyId++

	return nil
}

func (e *Enemy) AttackPlayer() string {
	ability := &Ability{
		"",
		0,
		0,
		0,
		0,
	}

	for i := 0; i < len(e.Abilities); i++ {
		if (e.Abilities[i].Damage > ability.Damage) && e.Abilities[i].CurrentCD == 0 {
			ability = &e.Abilities[i]
		}

		if e.Abilities[i].CurrentCD > 0 {
			e.Abilities[i].CurrentCD--
			log.Printf("Cooldown for %s is now %d\n", e.Abilities[i].Name, e.Abilities[i].CurrentCD)
		}
	}

	if ability.Name == "" {
		return fmt.Sprint("All abilities on cooldown")
	}

	diceRoll := DiceRoll()

	if (diceRoll + ability.Attack) < player.ArmorClass {
		return fmt.Sprintf("Ability %s missed you.", ability.Name)
	}

	if ability.CD > 0 && ability.CurrentCD == 0 {
		ability.CurrentCD = ability.CD
		log.Printf("Cooldown for %s set to %d\n", ability.Name, ability.CurrentCD)
	}

	player.CurrentHp = player.CurrentHp - ability.Damage
	log.Printf("PLAYER hit with %s for %d.", ability.Name, ability.Damage)

	if player.CurrentHp <= 0 {
		return fmt.Sprintf("You died.")
	}

	return fmt.Sprintf("You have been hit with %s for %d. Your HP is now %s.", ability.Name, ability.Damage, player.CurrentHp)
}
