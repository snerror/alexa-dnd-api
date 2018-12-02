package main

import "fmt"

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
		e.CurrentHp = 1
		e.MaxHp = 1
		e.ArmorClass = 16
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
