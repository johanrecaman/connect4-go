package game

// EvaluateSimple - Heurística simples para nível iniciante
func EvaluateSimple(b *Board) int {
	if b.CheckWin(1) {
		return 1000
	}
	if b.CheckWin(2) {
		return -1000
	}

	score := 0
	// Heurística simples: só sequências de 2
	score += countPotentialSequences(b, 1, 2) * 10
	score -= countPotentialSequences(b, 2, 2) * 10

	return score
}

// EvaluateIntermediate - Heurística intermediária
func EvaluateIntermediate(b *Board) int {
	if b.CheckWin(1) {
		return 1000
	}
	if b.CheckWin(2) {
		return -1000
	}

	score := 0
	// Sequências de 2 e 3 peças (ponderação)
	score += countPotentialSequences(b, 1, 2) * 10
	score += countPotentialSequences(b, 1, 3) * 50
	score -= countPotentialSequences(b, 2, 2) * 10
	score -= countPotentialSequences(b, 2, 3) * 50

	return score
}

// EvaluateAdvanced - Heurística avançada para nível profissional
func EvaluateAdvanced(b *Board) int {
	if b.CheckWin(1) {
		return 1000
	}
	if b.CheckWin(2) {
		return -1000
	}

	score := 0

	// Sequências de 2 e 3 peças
	score += countPotentialSequences(b, 1, 2) * 10
	score += countPotentialSequences(b, 1, 3) * 50
	score -= countPotentialSequences(b, 2, 2) * 10
	score -= countPotentialSequences(b, 2, 3) * 50

	// Bonus por centralidade (colunas 2, 3, 4)
	for row := 0; row < 6; row++ {
		if b.Grid[row][3] == 1 {
			score += 4 // Coluna central vale mais
		} else if b.Grid[row][3] == 2 {
			score -= 4
		}

		if b.Grid[row][2] == 1 || b.Grid[row][4] == 1 {
			score += 2 // Colunas próximas ao centro
		} else if b.Grid[row][2] == 2 || b.Grid[row][4] == 2 {
			score -= 2
		}
	}

	return score
}

// countPotentialSequences conta quantas sequências de tamanho n o jogador possui
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

