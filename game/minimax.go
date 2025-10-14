package game

import (
	"math"
	"sort"
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

func MinimaxNoPruning(board *Board, depth int, isMaximizing bool, evaluate func(*Board) int) int {
	if board.CheckWin(2) { // IA (amarelo) venceu
		return 1000000
	}
	if board.CheckWin(1) { // Humano (vermelho) venceu
		return -1000000
	}
	if board.IsFull() {
		return 0
	}

	if depth == 0 {
		return evaluate(board)
	}

	if isMaximizing {
		bestScore := math.MinInt
		for col := range 7 {
			if board.MakeMove(col, 2) {
				score := MinimaxNoPruning(board, depth-1, false, evaluate)
				board.UndoMove(col)
				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	} else {
		bestScore := math.MaxInt
		for col := range 7 {
			if board.MakeMove(col, 1) {
				score := MinimaxNoPruning(board, depth-1, true, evaluate)
				board.UndoMove(col)
				if score < bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}
}

func MinimaxAlphaBeta(board *Board, depth int, alpha, beta int, isMaximizing bool, evaluate func(*Board) int, deadline *time.Time) int {
	if deadline != nil && time.Now().After(*deadline) {
		return evaluate(board) // Retorna avaliação atual se tempo esgotou
	}

	if board.CheckWin(2) { // IA (amarelo) venceu
		return 1000000
	}
	if board.CheckWin(1) { // Humano (vermelho) venceu
		return -1000000
	}
	if board.IsFull() {
		return 0
	}

	if depth == 0 {
		return evaluate(board)
	}

	if isMaximizing {
		return maximizeScore(board, depth, alpha, beta, evaluate, deadline)
	}
	return minimizeScore(board, depth, alpha, beta, evaluate, deadline)
}

func maximizeScore(board *Board, depth, alpha, beta int, evaluate func(*Board) int, deadline *time.Time) int {
	bestScore := math.MinInt

	for col := range 7 {
		if board.MakeMove(col, 2) { // IA é player 2 (amarelo)
			score := MinimaxAlphaBeta(board, depth-1, alpha, beta, false, evaluate, deadline)
			board.UndoMove(col)

			if score > bestScore {
				bestScore = score
			}
			if bestScore > alpha {
				alpha = bestScore
			}

			if beta <= alpha {
				break // Poda alfa-beta
			}
		}
	}
	return bestScore
}

func minimizeScore(board *Board, depth, alpha, beta int, evaluate func(*Board) int, deadline *time.Time) int {
	bestScore := math.MaxInt

	for col := range 7 {
		if board.MakeMove(col, 1) { // Humano é player 1 (vermelho)
			score := MinimaxAlphaBeta(board, depth-1, alpha, beta, true, evaluate, deadline)
			board.UndoMove(col)

			if score < bestScore {
				bestScore = score
			}
			if bestScore < beta {
				beta = bestScore
			}

			if beta <= alpha {
				break // Poda alfa-beta
			}
		}
	}
	return bestScore
}

func getOrderedColumns() []int {
	// Ordem: 3 (centro), 2, 4, 1, 5, 0, 6
	return []int{3, 2, 4, 1, 5, 0, 6}
}

func getOrderedMoves(board *Board, evaluate func(*Board) int) []int {
	type moveScore struct {
		col   int
		score int
	}

	moves := []moveScore{}

	for col := range 7 {
		if board.MakeMove(col, 2) {
			score := evaluate(board)
			board.UndoMove(col)
			moves = append(moves, moveScore{col, score})
		}
	}

	sort.Slice(moves, func(i, j int) bool {
		return moves[i].score > moves[j].score
	})

	orderedCols := make([]int, len(moves))
	for i, m := range moves {
		orderedCols[i] = m.col
	}

	return orderedCols
}

func GetBestMove(board *Board, level AILevel) MoveResult {
	start := time.Now()

	var depth int
	var evaluate func(*Board) int
	var deadline *time.Time

	switch level {
	case Beginner:
		depth = 2 // Profundidade 2 para iniciante
		evaluate = EvaluateSimple

		bestCol := -1
		bestScore := math.MinInt

		for col := range 7 {
			if board.MakeMove(col, 2) { // IA é player 2
				score := MinimaxNoPruning(board, depth-1, false, evaluate)
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

	case Intermediate:
		depth = 5
		evaluate = EvaluateIntermediate

	case Professional:
		depth = 8 // Profundidade maior para profissional
		evaluate = EvaluateAdvanced

		timeLimit := start.Add(3 * time.Second)
		deadline = &timeLimit
	}

	bestCol := -1
	bestScore := math.MinInt

	var columns []int
	if level == Professional {
		columns = getOrderedMoves(board, evaluate)
	} else {
		columns = getOrderedColumns() // Ordenação simples por centralidade
	}

	for _, col := range columns {
		if board.MakeMove(col, 2) { // IA é player 2 (amarelo)
			score := MinimaxAlphaBeta(board, depth-1, math.MinInt, math.MaxInt, false, evaluate, deadline)
			board.UndoMove(col)

			if score > bestScore {
				bestScore = score
				bestCol = col
			}

			if deadline != nil && time.Now().After(*deadline) {
				break
			}
		}
	}

	return MoveResult{
		Column:   bestCol,
		Score:    bestScore,
		Duration: time.Since(start),
	}
}

