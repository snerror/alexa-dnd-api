package main

import (
	"log"
	"net/http"
)

var (
	dungeon = Dungeon{}
	player  Player
	enemies []Enemy
)

func main() {
	router := NewRouter()

	dungeon = CreateDungeon(10, 30)
	dungeon.generate()
	dungeon.DrawDungeon()

	log.Println("Server is up and running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
