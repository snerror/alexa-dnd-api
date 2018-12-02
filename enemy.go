package main

import (
	"fmt"
	"log"
)

var (
	ENEMY_SKELETON = "skeleton"
	ENEMY_SPIDER   = "spider"
)

var ENEMY_ID = 0

type Enemy struct {
	ID         int
	Name       string
	CurrentHp  int
	MaxHp      int
	ArmorClass int
	Abilities  []Ability
}

func (e *Enemy) CreatePreset(p string) error {
	if p != ENEMY_SKELETON && p != ENEMY_SPIDER {
		return fmt.Errorf("unknown enemy preset")
	}

	if p == ENEMY_SKELETON {
		e.Name = ENEMY_SKELETON
		e.CurrentHp = 10
		e.MaxHp = 1
		e.ArmorClass = 16

		a1 := Ability{
			Name:      "Slice",
			Attack:    9,
			Damage:    2,
			CD:        2,
			CurrentCD: 0,
		}

		a2 := Ability{
			Name:      "Strike",
			Attack:    7,
			Damage:    1,
			CD:        0,
			CurrentCD: 0,
		}

		e.Abilities = append(e.Abilities, a1)
		e.Abilities = append(e.Abilities, a2)
	}

	if p == ENEMY_SPIDER {
		e.Name = ENEMY_SPIDER
		e.CurrentHp = 1
		e.MaxHp = 1
		e.ArmorClass = 15
	}

	e.ID = ENEMY_ID
	ENEMY_ID++

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

	return fmt.Sprintf("You have been hit with %s for %d.", ability.Name, ability.Damage)
}
