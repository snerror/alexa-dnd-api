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

//
//func (d *Dungeon) Create(rows, cols int) *Dungeon {
//
//	return &Dungeon{c, h, v, cell, hor, ver}
//}

//func (d *Dungeon) generate(row, col int) {
//	fmt.Printf("up %v\n", up)
//	fmt.Printf("down %v\n", down)
//	fmt.Printf("right %v\n", right)
//	fmt.Printf("left %v\n", left)
//
//	d.cells[row][col] = down + right
//	fmt.Printf("current cell value %v\n", d.cells[row][col])
//}

const (
	up    = 1 << iota //1
	down              //2
	right             //4
	left              //8
)

func GenerateDungeon(row, col int) {
	rand.Seed(time.Now().UnixNano())

	if dungeon.cells[row][col] != 0 {
		fmt.Printf("CELL %d %d already generated with value %d. SKIPPING\n", row, col, dungeon.cells[row][col])
		return
	}

	dungeon.cells[row][col] = rand.Intn(16)
	fmt.Printf("CELL %d %d GENERATED value %v\n", row, col, dungeon.cells[row][col])

	if dungeon.cells[row][col] == 0 {
		fmt.Printf("CELL %d %d is NOT VALID with value %v. Recreating cell. \n", row, col, dungeon.cells[row][col])
		dungeon.cells[row][col] = 0
		GenerateDungeon(row, col)
	}

	// Border edge case
	if dungeon.cells[row][col]&up != 0 && row == 0 {
		fmt.Printf("CELL %d %d is NOT VALID with value %v. Recreating cell. \n", row, col, dungeon.cells[row][col])
		dungeon.cells[row][col] = 0
		GenerateDungeon(row, col)
	}

	if dungeon.cells[row][col]&left != 0 && col == 0 {
		fmt.Printf("CELL %d %d is NOT VALID with value %v. Recreating cell. \n", row, col, dungeon.cells[row][col])
		dungeon.cells[row][col] = 0
		GenerateDungeon(row, col)
	}

	if dungeon.cells[row][col]&down != 0 && row == dungeon.rows-1 {
		fmt.Printf("CELL %d %d is NOT VALID with value %v. Recreating cell. \n", row, col, dungeon.cells[row][col])
		dungeon.cells[row][col] = 0
		GenerateDungeon(row, col)
	}

	if dungeon.cells[row][col]&right != 0 && col == dungeon.cols-1 {
		fmt.Printf("CELL %d %d is NOT VALID with value %v. Recreating cell. \n", row, col, dungeon.cells[row][col])
		dungeon.cells[row][col] = 0
		GenerateDungeon(row, col)
	}

	// Generate next cell.
	if dungeon.cells[row][col]&up != 0 && row != 0 {
		fmt.Printf("UP to next CELL %d %d \n", row-1, col)
		GenerateDungeon(row-1, col)
	}

	fmt.Printf("DEBUG - ROW: %d DUNGEON ROWS: %d \n", row, dungeon.rows-1)
	if dungeon.cells[row][col]&down != 0 && row != dungeon.rows-1 {
		fmt.Printf("DOWN to next CELL %d %d \n", row+1, col)
		GenerateDungeon(row+1, col)
	}

	if dungeon.cells[row][col]&left != 0 && col != 0 {
		fmt.Printf("LEFT to next CELL %d %d \n", row, col-1)
		GenerateDungeon(row, col-1)
	}

	if dungeon.cells[row][col]&right != 0 && col != dungeon.cols-1 {
		fmt.Printf("RIGHT to next CELL %d %d \n", row, col+1)
		GenerateDungeon(row, col+1)
	}

	fmt.Printf("current cell value %v\n", dungeon.cells[row][col])
}

func (d *Dungeon) DrawDungeon() {

	hWall := "+---"
	hOpen := "+   "
	leftWall := "|   "
	rightWall := "   |"
	leftRightWall := "|   |"
	noWall := "    "
	//vOpen := []byte("    ")
	//rightCorner := []byte("+\n")
	//rightWall := []byte("|\n")

	var drawRow string

	for i := 0; i < d.cols; i++ {
		drawRow = drawRow + hWall
	}

	fmt.Printf("%v+\n", drawRow)

	drawRow = ""
	for i := 0; i < d.rows; i++ {
		for j := 0; j < d.cols; j++ {
			if d.cells[i][j]&left == 0 && d.cells[i][j]&right != 0 {
				drawRow = drawRow + leftWall
			}

			if d.cells[i][j]&left != 0 && d.cells[i][j]&right == 0 {
				if j != 0 && d.cells[i][j-1]&right == 0 {
					drawRow = drawRow[:len(drawRow)-1]
					drawRow = drawRow + " "
				}

				drawRow = drawRow + rightWall
			}

			if d.cells[i][j]&left == 0 && d.cells[i][j]&right == 0 {
				if j != 0 && d.cells[i][j-1]&right == 0 {
					drawRow = drawRow[:len(drawRow)-1]
					drawRow = drawRow + " "
				}
				drawRow = drawRow + leftRightWall
			}

			if d.cells[i][j]&left != 0 && d.cells[i][j]&right != 0 {
				drawRow = drawRow + noWall
			}
		}
		fmt.Printf("%v\n", drawRow)
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
