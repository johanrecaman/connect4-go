package game

func Minimax(board *Board, depth int, isMaximizing bool) int {
	var bestScore int

	if depth == 0 {
		return board.Evaluate()
	}
	if isMaximizing{
		bestScore = -1000
		for col := range 7 {
			if board.MakeMove(col){
				score := Minimax(board, depth-1, false)
				board.UndoMove(col)
				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}
	bestScore = 1000
	for col := range 7{
		if board.MakeMove(col){
			score := Minimax(board, depth-1, true)
			board.UndoMove(col)
			if score < bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}
