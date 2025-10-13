package game

import (
	"math"
	"time"
)

type AILevel int

const (
	Beginner     AILevel = 0
	Intermediate AILevel = 1
	Professional AILevel = 2
)

type MoveResult struct {
	Column   int
	Score    int
	Duration time.Duration
}

func Minimax(board *Board, depth int, isMaximizing bool, evaluate func(*Board) int) int {
	return MinimaxAlphaBeta(board, depth, math.MinInt, math.MaxInt, isMaximizing, evaluate)
}

func MinimaxAlphaBeta(board *Board, depth int, alpha, beta int, isMaximizing bool, evaluate func(*Board) int) int {
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
		return evaluate(board)
	}

	if isMaximizing {
		return maximizeScore(board, depth, alpha, beta, evaluate)
	}
	return minimizeScore(board, depth, alpha, beta, evaluate)
}

func maximizeScore(board *Board, depth, alpha, beta int, evaluate func(*Board) int) int {
	bestScore := math.MinInt

	for col := range 7 {
		if board.MakeMove(col, 1) {
			score := MinimaxAlphaBeta(board, depth-1, alpha, beta, false, evaluate)
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

func minimizeScore(board *Board, depth, alpha, beta int, evaluate func(*Board) int) int {
	bestScore := math.MaxInt

	for col := range 7 {
		if board.MakeMove(col, 2) {
			score := MinimaxAlphaBeta(board, depth-1, alpha, beta, true, evaluate)
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

func GetBestMove(board *Board, level AILevel) MoveResult {
	start := time.Now()

	var depth int
	var evaluate func(*Board) int

	switch level {
	case Beginner:
		depth = 3
		evaluate = EvaluateSimple
	case Intermediate:
		depth = 5
		evaluate = EvaluateIntermediate
	case Professional:
		depth = 7
		evaluate = EvaluateAdvanced
	}

	bestCol := -1
	bestScore := math.MinInt

	for col := range 7 {
		if board.MakeMove(col, 1) {
			score := MinimaxAlphaBeta(board, depth-1, math.MinInt, math.MaxInt, false, evaluate)
			board.UndoMove(col)

			if score > bestScore {
				bestScore = score
				bestCol = col
			}
		}
	}
	return MoveResult{
		Column:   bestCol,
		Score:    bestScore,
		Duration: time.Since(start),
	}
}
