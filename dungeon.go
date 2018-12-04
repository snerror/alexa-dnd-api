package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dungeon struct {
	rows, cols int
	cells      [][]int
}

func (d *Dungeon) generate() {
	d.generateRecursive(0, 0, 0)
	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.cols; j++ {
			fmt.Printf("%v ", dungeon.cells[i][j])
		}
		fmt.Printf("\n")
	}
}

const (
	up    = 1 << iota //1
	down              //2
	right             //4
	left              //8
)

func (d *Dungeon) generateRecursive(row, col, previousValue int) {
	rand.Seed(time.Now().UnixNano())

	var previousCell int

	if previousValue == up {
		previousCell = down
	} else if previousValue == down {
		previousCell = up
	} else if previousValue == right {
		previousCell = left
	} else if previousValue == left {
		previousCell = right
	}

	var possiblePaths []int

	if row != 0 && dungeon.cells[row-1][col] == 0 {
		possiblePaths = append(possiblePaths, up)
	}

	if row != dungeon.rows-1 && dungeon.cells[row+1][col] == 0 {
		possiblePaths = append(possiblePaths, down)
	}

	if col != 0 && dungeon.cells[row][col-1] == 0 {
		possiblePaths = append(possiblePaths, left)
	}

	if col != dungeon.cols-1 && dungeon.cells[row][col+1] == 0 {
		possiblePaths = append(possiblePaths, right)
	}

	if len(possiblePaths) == 0 {
		fmt.Printf("CELL %d %d reached DEAD END %d. SKIPPING\n", row, col, dungeon.cells[row][col])
		return
	}

	if dungeon.cells[row][col] != 0 {
		fmt.Printf("CELL %d %d ALREADY GENERATED  %d. SKIPPING\n", row, col, dungeon.cells[row][col])
		return
	}

	dungeon.cells[row][col] = possiblePaths[rand.Int()%len(possiblePaths)]
	fmt.Printf("CELL %d %d GENERATED value %v\n", row, col, dungeon.cells[row][col])

	dungeon.cells[row][col] = dungeon.cells[row][col] + previousCell
	fmt.Printf("CELL %d %d CONNECTED TO PREVIOUS CELL, VALUE NOW %v\n", row, col, dungeon.cells[row][col])

	generateEnemy := rand.Intn(5)
	if generateEnemy == 1 {
		var enemy Enemy

		possibleEnemies := []string{EnemySkeleton, EnemySpider, EnemyWraith, EnemyZombie, EnemyGargoyle, EnemyBlazingSkeleton}

		err := enemy.CreatePreset(possibleEnemies[rand.Int()%len(possibleEnemies)], row, col)
		if err != nil {
			fmt.Printf("Error while generating enemy: %s\n", err)
		} else {
			enemies = append(enemies, enemy)
			fmt.Printf("CELL %d %d GENERATED ENEMY %s\n", row, col, enemy.Name)
		}
	}

	// Generate next cell.
	if dungeon.cells[row][col]&up != 0 && row != 0 && up != previousCell {
		fmt.Printf("UP to next CELL %d %d \n", row-1, col)
		d.generateRecursive(row-1, col, up)
	}

	if dungeon.cells[row][col]&down != 0 && row != dungeon.rows-1 && down != previousCell {
		fmt.Printf("DOWN to next CELL %d %d \n", row+1, col)
		d.generateRecursive(row+1, col, down)
	}

	if dungeon.cells[row][col]&left != 0 && col != 0 && left != previousCell {
		fmt.Printf("LEFT to next CELL %d %d \n", row, col-1)
		d.generateRecursive(row, col-1, left)
	}

	if dungeon.cells[row][col]&right != 0 && col != dungeon.cols-1 && right != previousCell {
		fmt.Printf("RIGHT to next CELL %d %d \n", row, col+1)
		d.generateRecursive(row, col+1, right)
	}

	fmt.Printf("CELL %d %d is NOT VALID with value %v. Recreating cell. \n", row, col, dungeon.cells[row][col])
	d.generateRecursive(row, col, previousValue)
}

func (d *Dungeon) DrawDungeon() {
	hWall := "+---"
	hOpen := "+   "
	wall := "|"
	noWall := " "
	betweenWalls := "   "
	playerPos := " P "
	enemyPos := " E "

	var drawRow string

	for i := 0; i < d.cols; i++ {
		drawRow = drawRow + hWall
	}

	fmt.Printf("%v+\n", drawRow)

	drawRow = ""
	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.cols; j++ {
			// Find enemy on this location
			var enemy *Enemy
			for k := 0; k < len(enemies); k++ {
				if enemies[k].Position.X == i && enemies[k].Position.Y == j {
					enemy = &enemies[k]
				}
			}

			if d.cells[i][j]&left == 0 {
				drawRow = drawRow + wall
			} else {
				drawRow = drawRow + noWall
			}

			if player.Position.X == i && player.Position.Y == j {
				drawRow = drawRow + playerPos
			} else if enemy != nil {
				drawRow = drawRow + enemyPos
			} else {
				drawRow = drawRow + betweenWalls
			}
		}
		fmt.Printf("%v|\n", drawRow)
		drawRow = ""

		for j := 0; j < d.cols; j++ {
			if d.cells[i][j]&down == 0 {
				drawRow = drawRow + hWall
			} else {
				drawRow = drawRow + hOpen
			}
		}
		fmt.Printf("%v+\n", drawRow)
		drawRow = ""
	}
}

func (d *Dungeon) CheckMoveDirection(x, y, direction int) bool {
	if direction == up {
		if x == 0 {
			return false
		}

		if d.cells[x][y]&up == 0 {
			return false
		}

		return true
	}

	if direction == down {
		if x == d.rows-1 {
			return false
		}

		if d.cells[x][y]&down == 0 {
			return false
		}

		return true
	}

	if direction == left {
		if y == 0 {
			return false
		}

		if d.cells[x][y]&left == 0 {
			return false
		}

		return true
	}

	if direction == right {
		if y == d.cols-1 {
			return false
		}

		if d.cells[x][y]&right == 0 {
			return false
		}

		return true
	}

	return false
}
