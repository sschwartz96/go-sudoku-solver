package main

import (
	"fmt"
)

type SudokuPuzzle struct {
	Grid [9][9]SudokuCell
}

func (puzzle *SudokuPuzzle) printPuzzle() {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			fmt.Print(puzzle.Grid[x][y].Value)
			if x != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

type SudokuCell struct {
	Value int
}
