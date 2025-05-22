package main

import (
	"fmt"
	"math"
)

// node
type Node struct {
	state [5][5]int
	eval  int
}

func getBestMove(board *[5][5]int, depth, player int) int {
	var is_maxminizing bool //1 -> X maxing, 2 -> O mining
	var best_move [2]int
	var best_eval, curr_eval, ret int
	if player == 1 {
		is_maxminizing = true
		best_eval = -math.MaxInt8
	} else {
		is_maxminizing = false
		best_eval = math.MaxInt8
	}
	for _, move := range possibleMoves(*board) {
		row, col := move[0], move[1]
		board[row][col] = player
		curr_eval = minimax(board, depth, -math.MaxInt8, math.MaxInt8, is_maxminizing)
		board[row][col] = 0

		if is_maxminizing { // Current player is maximizing
			if curr_eval > best_eval {
				best_eval = curr_eval
				best_move = move
			}
		} else { // Current player is minimizing
			if curr_eval < best_eval {
				best_eval = curr_eval
				best_move = move
			}
		}

		ret = (best_move[0]+1)*10 + best_move[1] + 1
		fmt.Println(best_eval, " <-score | move ->", ret)
	}
	// Concatenate best_move[0] and best_move[1] as described

	return ret
}

// minimax, X is maxing
func minimax(board *[5][5]int, depth, alpha, beta int, is_maxminizing bool) int {
	var max_eval, min_eval int
	game_over, end_eval := isGameOver(*board)

	if depth == 0 || game_over {
		if game_over {
			return end_eval
		}
		return evaluateNode(*board)
	}

	if is_maxminizing {
		max_eval = -math.MaxInt8
		moves := possibleMoves(*board)
		for _, move := range moves {
			row, col := move[0], move[1]
			// fmt.Println(row, " <r c> ", col, " max")

			board[row][col] = 1 // Apply move for player X
			eval := minimax(board, depth-1, alpha, beta, false)
			board[row][col] = 0 // Undo move

			max_eval = max(max_eval, eval)
			alpha = max(alpha, eval)
			if beta <= alpha {
				break
			}
		}
		return max_eval
	} else {
		min_eval = math.MaxInt8
		moves := possibleMoves(*board)
		for _, move := range moves {
			row, col := move[0], move[1]

			// fmt.Println(row, " <r c> ", col, " min")

			board[row][col] = 2 // Apply move for player O
			eval := minimax(board, depth-1, alpha, beta, true)
			board[row][col] = 0 // Undo move

			min_eval = min(min_eval, eval)
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return min_eval
	}
}

// returns list of possible moves for a player in certain state
func possibleMoves(board [5][5]int) [][2]int {
	var moves [][2]int
	for r := range 5 {
		for c := range 5 {
			if board[r][c] == 0 { // 0 represents an empty cell
				moves = append(moves, [2]int{r, c})
			}
		}
	}
	return moves
}

func evaluateNode(board [5][5]int) int {
	score := 0
	patternFound := false

	// Helper function to check line patterns
	checkLine := func(line [5]int) {
		// Block O from winning
		for i := range 3 {
			if line[i] == 2 && line[i+1] == 2 && line[i+2] == 2 {
				score += 10 // Block O
				patternFound = true
			}
			if line[i] == 1 && line[i+1] == 1 && line[i+2] == 1 {
				score -= 10 // Block X
				patternFound = true
			}
		}
		// Two in a row
		for i := range 4 {
			if line[i] == 1 && line[i+1] == 1 {
				score += 5
				patternFound = true
			}
			if line[i] == 2 && line[i+1] == 2 {
				score -= 5
				patternFound = true
			}
		}
		// X 0 X pattern
		for i := range 3 {
			if line[i] == 1 && line[i+1] == 0 && line[i+2] == 1 {
				score += 3
				patternFound = true
			}
			if line[i] == 2 && line[i+1] == 0 && line[i+2] == 2 {
				score -= 3
				patternFound = true
			}
		}
		// X 0 0 X pattern
		for i := range 2 {
			if line[i] == 1 && line[i+1] == 0 && line[i+2] == 0 && line[i+3] == 1 {
				score += 2
				patternFound = true
			}
			if line[i] == 2 && line[i+1] == 0 && line[i+2] == 0 && line[i+3] == 2 {
				score -= 2
				patternFound = true
			}
		}
		// X X 0 X pattern
		for i := range 2 {
			if line[i] == 1 && line[i+1] == 1 && line[i+2] == 0 && line[i+3] == 1 {
				score += 20
				patternFound = true
			}
			if line[i] == 2 && line[i+1] == 2 && line[i+2] == 0 && line[i+3] == 2 {
				score -= 20
				patternFound = true
			}
		}
		// Strongly unfavor X 0 X 0 X pattern
		if len(line) == 5 {
			if line[0] == 1 && line[1] == 0 && line[2] == 1 && line[3] == 0 && line[4] == 1 {
				score -= 50 // Strong penalty for X 0 X 0 X
				patternFound = true
			}
			if line[0] == 2 && line[1] == 0 && line[2] == 2 && line[3] == 0 && line[4] == 2 {
				score += 50 // Strong penalty for O 0 O 0 O (from X's perspective)
				patternFound = true
			}
		}
	}

	// Check rows
	for r := range 5 {
		var row [5]int
		for c := range 5 {
			row[c] = board[r][c]
		}
		checkLine(row)
	}

	// Check columns
	for c := range 5 {
		var col [5]int
		for r := range 5 {
			col[r] = board[r][c]
		}
		checkLine(col)
	}

	// Check diagonals (top-left to bottom-right)
	for d := -1; d <= 1; d++ {
		var diag [5]int
		count := 0
		for i := range 5 {
			j := i + d
			if j >= 0 && j < 5 {
				diag[count] = board[i][j]
				count++
			}
		}
		if count >= 2 {
			checkLine(diag)
		}
	}

	// Check anti-diagonals (top-right to bottom-left)
	for d := 3; d <= 5; d++ {
		var adiag [5]int
		count := 0
		for i := range 5 {
			j := d - i
			if j >= 0 && j < 5 {
				adiag[count] = board[i][j]
				count++
			}
		}
		if count >= 2 {
			checkLine(adiag)
		}
	}

	// Base case: no pattern found
	if !patternFound {
		return 1
	}

	return score
}

func isGameOver(board [5][5]int) (bool, int) {
	//TODO : implement game end check
	game_over, p := drawCheck(board)

	if game_over {
		return true, 0
	}

	game_over, p = winCheck(board)
	if game_over {
		if p == 1 {
			return true, 100
		} else {
			return true, -100
		}
	}

	game_over, p = loseCheck(board)
	if game_over {
		if p == 1 {
			return true, -100
		} else {
			return true, 100
		}
	}
	return false, 0
}
