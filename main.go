package main

import (
	"log"
	"net/http"
)

var (
	dungeon             = Dungeon{}
	player              Player
	enemies             []Enemy
	state               string
	STATE_INITIAL       = "initial"
	STATE_COMBAT_PLAYER = "combat_player"
	STATE_COMBAT_ENEMY  = "combat_enemy"
)

func main() {
	router := NewRouter()

	dungeon = Dungeon{
		5,
		5,
		[][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0},
		},
	}

	dungeon.generate()
	dungeon.DrawDungeon()

	log.Println("Server is up and running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
