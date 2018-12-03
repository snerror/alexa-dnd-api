package main

import (
	"fmt"
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
	//dungeon.Create(5, 5)
	//dungeon.generate(0, 0)

	dungeon = Dungeon{
		3,
		3,
		[][]int{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}

	//dungeon = Dungeon{
	//	3,
	//	3,
	//	[][]int{
	//		{2, 14, 8},
	//		{7, 7, 11},
	//		{5, 13, 9},
	//	},
	//}

	GenerateDungeon(0, 0, 0)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v ", dungeon.cells[i][j])
		}
		fmt.Printf("\n")
	}

	dungeon.DrawDungeon()

	log.Println("Server is up and running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
