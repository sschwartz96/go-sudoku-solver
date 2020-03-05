package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const LOG = false

var start time.Time

func main() {
	// Load puzzle
	puzzle := loadPuzzle()

	// Print beginning puzzle
	puzzle.printPuzzle()

	// Start timer
	start = time.Now()

	// Solve the puzzle
	puzzle.solve()
}

// isPossible checks the vertical row, horizontal row, and square section if value can be placed
func (puzzle *SudokuPuzzle) isPossible(x, y, value int) bool {
	// check the rows
	for i := 0; i < 9; i++ {
		if puzzle.Grid[x][i].Value == value || puzzle.Grid[i][y].Value == value {
			return false
		}
	}

	// get the beginning coords of square section
	baseI := x / 3 * 3
	baseJ := y / 3 * 3
	// check the section
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if puzzle.Grid[i+baseI][j+baseJ].Value == value {
				return false
			}
		}
	}

	return true
}

// solve uses recursive back tracking to solve the puzzle
func (puzzle *SudokuPuzzle) solve() bool {
	// the first two for loops will scan through the entire puzzle column by column
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			// if a "move" can be made
			if puzzle.Grid[x][y].Value == 0 {
				// go through all possibilities
				for v := 1; v < 10; v++ {
					// if we can enter the value then recursively solve the rest of the puzzle
					if puzzle.isPossible(x, y, v) {
						puzzle.Grid[x][y].Value = v
						// check if we did solve the puzzle
						if puzzle.solve() {
							puzzle.showSolution()
						}
						// this allows us to continue whether solved or not
						puzzle.Grid[x][y].Value = 0
					}
				}
				return false
			}
		}
	}

	// only way to get at this point was if the puzzle was solved (all cells have values != 0)
	// program will continue to exhaust all other possibilities
	return true
}

func (puzzle *SudokuPuzzle) showSolution() {
	elapsed := time.Since(start)
	fmt.Printf("finished solving in %v: \n", elapsed)
	puzzle.printPuzzle()
	fmt.Println("\n Try more?")
	bufio.NewReader(os.Stdin).ReadLine()
}

func loadPuzzle() *SudokuPuzzle {
	// Get path of puzzle
	fmt.Println("Please enter path of puzzle: ")
	var path string
	_, err := fmt.Scanln(&path)
	if err != nil {
		if strings.Contains(err.Error(), "newline") {
			path = "EXAMPLE.puzzle"
		} else {
			log.Fatalf("Could not load in puzzle, error: %v", err)
		}
	}

	// Open file of puzzle and load into program
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Invalid path, error: %v", err)
	}
	defer file.Close()

	// Scan line by line
	scanner := bufio.NewScanner(file)
	lineCount := 0
	puzzle := &SudokuPuzzle{}
	for scanner.Scan() {
		line := scanner.Text()
		// Setup our puzzle
		for col, cell := range line {
			puzzle.Grid[col][lineCount] = SudokuCell{Value: int(cell) - 48}
		}
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning puzzle file, error: %v", err)
	}

	return puzzle
}
