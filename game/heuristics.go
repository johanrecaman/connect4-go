package game

func EvaluateSimple(b *Board) int {
	if b.CheckWin(1) {
		return 1000
	} 
	if b.CheckWin(2) {
		return -1000
	}
	score := 0
	score += countPotentialSequences(b, 1, 2) * 10
	score += countPotentialSequences(b, 1, 3) * 50
	score -= countPotentialSequences(b, 2, 2) * 10
	score -= countPotentialSequences(b, 2, 3) * 50

	return score
}

func countPotentialSequences(b *Board, player, n int) int {
	count := 0
	directions := [][2]int{
		{0, 1},  // horizontal
		{1, 0},  // vertical
		{1, 1},  // diagonal \
		{-1, 1}, // diagonal /
	}

	for row := range b.Grid {
		for col := range b.Grid[row] {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				seq := []int{}

				for k := range n {
					x, y := row+k*dx, col+k*dy
					if x < 0 || x >= 6 || y < 0 || y >= 7 {
						break
					}
					seq = append(seq, b.Grid[x][y])
				}

				if len(seq) == n && isSequenceOpen(seq, player) {
					count++
				}
			}
		}
	}
	return count
}

func isSequenceOpen(seq []int, player int) bool {
	for _, v := range seq {
		if v != player {
			return false
		}
	}
	return true
}
