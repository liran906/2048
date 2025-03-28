package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var leftFlag bool
var upFlag bool

func main() {
	fmt.Println("|=== Enjoy a game of ===|")
	game := createNewGame()
	printGame(game, 0, 0)
	score := 0

	for moves := 0; !loseGame(game); moves++ {
		var nextstep string = "n"

		// for invalid steps
		for !strings.Contains("asdwASDW", nextstep) || len(nextstep) != 1 {
			fmt.Printf("Please determine next move: (W/S/A/D)")
			fmt.Scan(&nextstep)
		}
		score += nextGame(game, nextstep)

		if !leftFlag && strings.Contains("Aa", nextstep) {
			leftFlag = true
		} else if leftFlag && strings.Contains("Dd", nextstep) {
			leftFlag = false
		}
		if !upFlag && strings.Contains("Ww", nextstep) {
			upFlag = true
		} else if upFlag && strings.Contains("Ss", nextstep) {
			upFlag = false
		}
		printGame(game, score, moves)
	}
	fmt.Println("\n=======  Game Over  =======")
	fmt.Printf("Congrats! Your Score: % 4d!\n", score)
	fmt.Println("===========================")
}

func createNewGame() (game [][]int) {
	game = make([][]int, 4)
	for i := range game {
		game[i] = make([]int, 4)
	}
	addRandomNums(game, 4)
	return
}

func printGame(game [][]int, score int, moves int) {
	if score == 0 {
		fmt.Println("|=======[  2048  ]========|")
	}
	fmt.Println("|=========================|")
	fmt.Printf("||Score:% 5d|Moves: % 4d||\n", score, moves)
	fmt.Println("|=========================|")

	for r := range game {
		if !upFlag {
			fmt.Println("||     |     |     |     ||")
		}
		fmt.Printf("|")
		for c := range game[r] {
			if !leftFlag {
				if game[r][c] > 0 {
					fmt.Printf("|%5d", game[r][c])
				} else {
					fmt.Printf("|     ")
				}
			} else {
				if game[r][c] > 0 {
					fmt.Printf("|%-5d", game[r][c])
				} else {
					fmt.Printf("|     ")
				}
			}
		}
		fmt.Printf("||\n")
		if upFlag {
			fmt.Println("||     |     |     |     ||")
		}
		fmt.Printf("|=========================|\n")
	}
}

// adding random 2 or 4 in the 0
func addRandomNums(game [][]int, nums int) {
	for range nums {
		emptyCount := 0
		for r := range game {
			for c := range game[r] {
				if game[r][c] == 0 {
					emptyCount++
				}
			}
		}
		if emptyCount > 0 {
			randNum := rand.Intn(emptyCount) + 1
		OuterLoop:
			for r := range game {
				for c := range game[r] {
					if game[r][c] == 0 {
						randNum--
					}
					if randNum == 0 {
						game[r][c] = 2 * (rand.Intn(2) + 1)
						break OuterLoop
					}
				}
			}
		}
	}
}

func loseGame(game [][]int) bool {
	n := len(game)
	for r := range game {
		for c := range game[0] {
			if game[r][c] == 0 {
				return false
			}
			if r < n-1 && game[r][c] == game[r+1][c] {
				return false
			}
			if r > 1 && game[r][c] == game[r-1][c] {
				return false
			}
			if c < n-1 && game[r][c] == game[r][c+1] {
				return false
			}
			if c > 1 && game[r][c] == game[r][c-1] {
				return false
			}
		}
	}
	return true
}

func nextGame(game [][]int, move string) (count int) {
	if move == "A" || move == "a" {
		for row := range game {
			count += moveAndMergeLeft(game[row])
		}
	} else if move == "W" || move == "w" {
		rotateCCW(game)
		for row := range game {
			count += moveAndMergeLeft(game[row])
		}
		rotateCW(game)
	} else if move == "S" || move == "s" {
		rotateCW(game)
		for row := range game {
			count += moveAndMergeLeft(game[row])
		}
		rotateCCW(game)
	} else if move == "D" || move == "d" {
		rotateCW(game)
		rotateCW(game)
		for row := range game {
			count += moveAndMergeLeft(game[row])
		}
		rotateCCW(game)
		rotateCCW(game)
	}
	addRandomNums(game, 1)
	return
}

// return the score by this step
func moveAndMergeLeft(row []int) (count int) {
	// merge same numbers
	for i := range row {
		if row[i] != 0 {
			for j := i + 1; j < len(row); j++ {
				if row[j] == row[i] {
					count += row[i]
					row[i] = row[i] * 2
					row[j] = 0
					break
				} else if row[j] != 0 {
					break
				}
			}
		}
	}

	// move all zeros to right side
	for l, r := 0, 0; r < len(row); r++ {
		if row[r] != 0 {
			row[l], row[r] = row[r], row[l]
			l++
		}
	}
	return
}

func rotateCCW(game [][]int) {
	n := len(game)
	// transpose matrix first
	for r := 0; r < n; r++ {
		for c := r; c < n; c++ {
			game[r][c], game[c][r] = game[c][r], game[r][c]
		}
	}
	// then flip horizonal
	for r := 0; r < n/2; r++ {
		game[r], game[n-r-1] = game[n-r-1], game[r]
	}
}

func rotateCW(game [][]int) {
	n := len(game)
	// transpose matrix first
	for r := 0; r < n; r++ {
		for c := r; c < n; c++ {
			game[r][c], game[c][r] = game[c][r], game[r][c]
		}
	}
	// then flip vertical
	for r := 0; r < n; r++ {
		for c := 0; c < n/2; c++ {
			game[r][c], game[r][n-c-1] = game[r][n-c-1], game[r][c]
		}
	}
}
