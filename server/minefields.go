package entrapped

const (
	minSize int = 7
	maxSize int = 10
)

const (
	bonusNumLife int = 3
)

const (
	open      int = 9
	empty     int = 0
	mine      int = -2
	bonusLife int = 3
)

type mineField struct {
	field []int
	size  int
}

func createEmptyMineField(size int) *mineField {
	if size < minSize {
		size = minSize
	}
	if size > maxSize {
		size = maxSize
	}

	field := make([]int, (size * size))

	return &mineField{field, size}
}

func (m *mineField) addRandomBombs(numBombs int) *mineField {
	size := len(m.field)

	for i := 0; i < numBombs; i++ {
		x := randomInt(size - 1)

		val := m.field[x]
		if val == empty {
			m.field[x] = mine
		} else if val != bonusLife {
			i--
		}
	}

	return m
}

func (m *mineField) addRandomLifes(numBonusLifes int) *mineField {
	size := len(m.field)

	for i := 0; i < numBonusLifes; i++ {
		x := randomInt(size - 1)

		val := m.field[x]
		if val == empty {
			m.field[x] = bonusLife
		} else if val != mine {
			i--
		}
	}

	return m
}

func (m *mineField) checkIndex(idx int) int {
	ele := m.field[idx]
	m.field[idx] = open
	return ele
}
