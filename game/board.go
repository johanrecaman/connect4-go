package game

type Board struct {
	Grid [6][7]int // 6 rows, 7 columns
}

func NewBoard() Board {
	return Board{}
}

func (b Board) PrintBoard() {
	for _, row := range b.Grid {
		for _, cell := range row {
			print(cell, " ")
		}
		println()
	}
}

func (b *Board) MakeMove(column int, player int) bool {
	if column < 0 || column >= 7 {
		return false
	}
	for i := 5; i >= 0; i-- {
		if b.Grid[i][column] == 0 {
			b.Grid[i][column] = player
			return true
		}
	}
	return false
}

func (b *Board) UndoMove(column int) {
	if column < 0 || column >= 7 {
		return
	}
	for i := range b.Grid {
		if b.Grid[i][column] != 0 {
			b.Grid[i][column] = 0
			return
		}
	}
}


func (b Board) Evaluate() int {
	if b.checkWin(1) {
		return 1000
	} else if b.checkWin(2) {
		return -1000
	}
	score := 0
	score += b.countPotentialSequences(1, 2) * 10
	score += b.countPotentialSequences(1, 3) * 50
	score -= b.countPotentialSequences(2, 2) * 10
	score -= b.countPotentialSequences(2, 3) * 50

	return score
}

func (b *Board) checkWin(player int) bool {
	return b.countPotentialSequences(player, 4) > 0
}

func (b *Board) countPotentialSequences(player, n int) int {
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
				for k := range make([]int, n) { // loop moderno
					x, y := row+k*dx, col+k*dy
					if x < 0 || x >= 6 || y < 0 || y >= 7 {
						break
					}
					seq = append(seq, b.Grid[x][y])
				}
				if len(seq) == n && b.isSequenceOpen(seq, player) {
					count++
				}
			}
		}
	}
	return count
}

func (b *Board) isSequenceOpen(seq []int, player int) bool {
	for _, v := range seq {
		if v != player { 
			return false
		}
	}
	return true
}



