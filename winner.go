package main


func (b Board) hasAnyHorizontalMatch(player Player) (bool, int) {
	for index, row := range b {
		if hasHorizontalMatch(row, player) {
			return true, index + 1
		}
	}
	return false, -1
}

func (b Board) hasAnyVerticalMatch(player Player) (bool, int) {
	i := 0
	for i < len(b[0]) {
		if hasVerticalMatch(b, i, 1) {
			return true, 1
		}
		if hasVerticalMatch(b, i, 2) {
			return true, 2
		}
		i++
	}
	return false, -1
}

func (b Board) hasAnyDiagonalMatch(player Player) (bool, int) {
	if hasDiagonalMatch(b, 1) {
		return true, 1
	}
	if hasDiagonalMatch(b, 2) {
		return true, 2
	}

	return false, -1
}

func hasHorizontalMatch(line [3]Player, char Player) bool {
	return line[0] == char && line[1] == char && line[2] == char
}
func hasVerticalMatch(b Board, vline int, char Player) bool {
	return b[0][vline] == char && b[1][vline] == char && b[2][vline] == char
}
func hasDiagonalMatch(b Board, char Player) bool {
	return (b[0][0] == char && b[1][1] == char && b[2][2] == char) || (b[0][2] == char && b[1][1] == char && b[2][0] == char)
}

