package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	dungeon = DrawDungeon(5, 5)
	player  Player
	enemies []Enemy
	state   string

	STATE_INITIAL       = "initial"
	STATE_COMBAT_PLAYER = "combat_player"
	STATE_COMBAT_ENEMY  = "combat_enemy"
)

func main() {
	router := NewRouter()
	dungeon.generator()

	log.Println("Server is up and running on port 8080")
	fmt.Println(dungeon)

	log.Fatal(http.ListenAndServe(":8080", router))
}
