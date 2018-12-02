package main

import (
	"fmt"
	"log"
	"net/http"
)

var dungeon = DrawDungeon(5, 5)
var player Player
var enemies []Enemy

func main() {
	router := NewRouter()
	dungeon.generator()

	var tempEnemy Enemy
	tempEnemy.CreatePreset(ENEMY_SKELETON)

	enemies = append(enemies, tempEnemy)

	log.Println("Server is up and running on port 8080")
	fmt.Println(dungeon)

	log.Fatal(http.ListenAndServe(":8080", router))
}
