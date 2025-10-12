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

func (b *Board) MakeMove(column int) bool {
	if column < 0 || column >= 7 {
		return false
	}
	for i := 5; i >= 0; i-- {
		if b.Grid[i][column] == 0 {
			b.Grid[i][column] = 1
			return true
		}
	}
	return false
}

func (b *Board) UndoMove(column int){
	if column < 0 || column >= 7 {
		return
	}
	for i := range 6 {
		if b.Grid[i][column] != 0 {
			b.Grid[i][column] = 0
			return
		}
	}
}

func (b Board) Evaluate() int {
	return 0
}
