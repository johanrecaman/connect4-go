package game

import "math"

func Minimax(board *Board, depth int, isMaximizing bool) int {
	return MinimaxAlphaBeta(board, depth, math.MinInt, math.MaxInt, isMaximizing)
}

func MinimaxAlphaBeta(board *Board, depth int, alpha, beta int, isMaximizing bool) int {
	if board.CheckWin(1) {
		return 1000000
	}
	if board.CheckWin(2) {
		return -1000000
	}
	if board.IsFull() {
		return 0
	}

	if depth == 0 {
		return EvaluateSimple(board)
	}

	if isMaximizing {
		return maximizeScore(board, depth, alpha, beta)
	}
	return minimizeScore(board, depth, alpha, beta)
}

func maximizeScore(board *Board, depth, alpha, beta int) int {
	bestScore := math.MinInt

	for col := range 7 {
		if board.MakeMove(col, 1) {
			score := MinimaxAlphaBeta(board, depth-1, alpha, beta, false)
			board.UndoMove(col)

			if score > bestScore {
				bestScore = score
			}
			if bestScore > alpha {
				alpha = bestScore
			}

			if beta <= alpha {
				break
			}
		}
	}
	return bestScore
}

func minimizeScore(board *Board, depth, alpha, beta int) int {
	bestScore := math.MaxInt

	for col := range 7 {
		if board.MakeMove(col, 2) {
			score := MinimaxAlphaBeta(board, depth-1, alpha, beta, true)
			board.UndoMove(col)

			if score < bestScore {
				bestScore = score
			}
			if bestScore < beta {
				beta = bestScore
			}

			if beta <= alpha {
				break
			}
		}
	}
	return bestScore
}

