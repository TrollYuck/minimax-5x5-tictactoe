package main

import (
	"math"
)

// node
type Node struct {
	state [5][5]int
	eval  int
}

// minimax, X is maxing
func minimax(board *[5][5]int, depth, alpha, beta int, is_maxminizing bool) int {
	var max_eval, min_eval int
	var game_over bool = false

	if depth == 0 || game_over {
		return evaluateNode(*board)
	}

	if is_maxminizing {
		max_eval = -math.MaxInt8
		moves := possibleMoves(*board)
		for _, move := range moves {
			row, col := move[0], move[1]

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

			board[row][col] = 2 // Apply move for player O
			eval := minimax(board, depth-1, alpha, beta, true)
			board[row][col] = 0 // Undo move

			min_eval = max(min_eval, eval)
			beta = min(beta, eval)
			if beta <= alpha {
				break
			}
		}
		return min_eval
	}
}

// returns list of posibble moves for a player in certain state
func possibleMoves(board [5][5]int) [][2]int {
	var moves [][2]int
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if board[r][c] == 0 { // 0 represents an empty cell
				moves = append(moves, [2]int{r, c})
			}
		}
	}
	return moves
}

func evaluateNode(board [5][5]int) int {
	//TODO : implement node eval according to the position
	return 0
}

func isGameOver(board [5][5]int) (bool, int) {
	//TODO : implement game end check
	return false, 0
}
