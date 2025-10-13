package game

import "math"

func Minimax(board *Board, depth int, isMaximizing bool) int {

	if board.CheckWin(1) {
		return 1000
	}
	if board.CheckWin(2) {
		return -1000
	}
	if board.IsFull() {
		return 0
	}
	if depth == 0 {
		return EvaluateSimple(board)
	}

	var bestScore int

	if depth == 0 {
		return EvaluateSimple(board)
	}
	if isMaximizing{
		bestScore = math.MinInt
		for col := range 7 {
			if board.MakeMove(col, 1){
				score := Minimax(board, depth-1, false)
				board.UndoMove(col)
				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}
	bestScore = math.MaxInt
	for col := range 7{
		if board.MakeMove(col, 2){
			score := Minimax(board, depth-1, true)
			board.UndoMove(col)
			if score < bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}
