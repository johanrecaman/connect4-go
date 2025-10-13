package game

import "math"

func Minimax(board *Board, depth int, isMaximizing bool) int {
	var bestScore int

	if depth == 0 {
		return board.Evaluate()
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
