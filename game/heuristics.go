package game

func EvaluateSimple(b *Board) int {
	if b.CheckWin(2) { // IA venceu
		return 1000
	}
	if b.CheckWin(1) { // Humano venceu
		return -1000
	}

	score := 0
	score += countPotentialSequences(b, 2, 2) * 10  // IA
	score -= countPotentialSequences(b, 1, 2) * 10  // Humano

	return score
}

func EvaluateIntermediate(b *Board) int {
	if b.CheckWin(2) { // IA venceu
		return 1000
	}
	if b.CheckWin(1) { // Humano venceu
		return -1000
	}

	score := 0
	score += countPotentialSequences(b, 2, 2) * 10   // IA - 2 peças
	score += countPotentialSequences(b, 2, 3) * 50   // IA - 3 peças
	score -= countPotentialSequences(b, 1, 2) * 10   // Humano - 2 peças
	score -= countPotentialSequences(b, 1, 3) * 50   // Humano - 3 peças

	return score
}

func EvaluateAdvanced(b *Board) int {
	if b.CheckWin(2) { // IA venceu
		return 1000
	}
	if b.CheckWin(1) { // Humano venceu
		return -1000
	}

	score := 0

	score += countPotentialSequences(b, 2, 2) * 10   // IA - 2 peças
	score += countPotentialSequences(b, 2, 3) * 100  // IA - 3 peças (mais crítico)
	score -= countPotentialSequences(b, 1, 2) * 10   // Humano - 2 peças
	score -= countPotentialSequences(b, 1, 3) * 100  // Humano - 3 peças (bloqueio crítico)

	for row := 0; row < 6; row++ {
		if b.Grid[row][3] == 2 { // IA na coluna central
			score += 4
		} else if b.Grid[row][3] == 1 { // Humano na coluna central
			score -= 4
		}

		if b.Grid[row][2] == 2 || b.Grid[row][4] == 2 { // IA nas colunas próximas
			score += 2
		} else if b.Grid[row][2] == 1 || b.Grid[row][4] == 1 { // Humano nas colunas próximas
			score -= 2
		}
	}

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

// isSequenceOpen verifica se a sequência pertence ao jogador
func isSequenceOpen(seq []int, player int) bool {
	for _, v := range seq {
		if v != player {
			return false
		}
	}
	return true
}
